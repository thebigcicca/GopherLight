package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BrunoCiccarino/GopherLight/logger"
	"github.com/BrunoCiccarino/GopherLight/plugins"
	"github.com/BrunoCiccarino/GopherLight/req"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type App struct {
	routes      map[string]map[string]http.HandlerFunc
	middlewares []Middleware
	plugins     []plugins.Plugin
}

func NewApp() *App {
	return &App{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}

// Use method to add middleware to the App
func (a *App) Use(mw Middleware) {
	a.middlewares = append(a.middlewares, mw)
}

func (a *App) AddPlugin(p plugins.Plugin) {
	a.plugins = append(a.plugins, p)
}

func (a *App) RegisterPlugins() {
	for _, plugin := range a.plugins {
		plugin.Register(a.Route)
	}
}

func (a *App) Route(method, path string, handler req.Handler) {
	if a.routes[path] == nil {
		a.routes[path] = make(map[string]http.HandlerFunc)
	}

	// Wrap the handler to convert it to http.HandlerFunc
	var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		request := req.NewRequest(r)
		response := req.NewResponse(w)
		handler(request, response)
	}

	// Apply middlewares to the handler
	for i := len(a.middlewares) - 1; i >= 0; i-- {
		h = a.middlewares[i](h)
	}

	a.routes[path][method] = h
}

func (a *App) Get(path string, handler req.Handler) {
	a.Route(http.MethodGet, path, handler)
}

func (a *App) Post(path string, handler req.Handler) {
	a.Route(http.MethodPost, path, handler)
}

func (a *App) Put(path string, handler req.Handler) {
	a.Route(http.MethodPut, path, handler)
}

func (a *App) Delete(path string, handler req.Handler) {
	a.Route(http.MethodDelete, path, handler)
}

func (a *App) Patch(path string, handler req.Handler) {
	a.Route(http.MethodPatch, path, handler)
}

func (a *App) Options(path string, handler req.Handler) {
	a.Route(http.MethodOptions, path, handler)
}

func (a *App) Head(path string, handler req.Handler) {
	a.Route(http.MethodHead, path, handler)
}

func (a *App) Connect(path string, handler req.Handler) {
	a.Route(http.MethodConnect, path, handler)
}

func (a *App) Trace(path string, handler req.Handler) {
	a.Route(http.MethodTrace, path, handler)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handlers, exists := a.routes[r.URL.Path]; exists {
		if handler, methodExists := handlers[r.Method]; methodExists {
			handler(w, r)
		} else {
			logger.LogWarning("Method not allowed: " + r.Method + " on path: " + r.URL.Path)
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		logger.LogError("Route not found: " + r.URL.Path)
		http.NotFound(w, r)
	}
}

func (a *App) Listen(addr string) error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(stop)

	srv := &http.Server{
		Addr:    addr,
		Handler: a,
	}
	// Start the server in a goroutine
	serverError := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverError <- err
		}
	}()

	// Wait for interrupt signal or server error
	select {
	case <-stop:
		log.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("server shutdown failed: %w", err)
		}
		log.Println("Server gracefully stopped.")
		return nil
	case err := <-serverError:
		return fmt.Errorf("server error: %w", err)
	}
}
