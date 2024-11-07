package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/BrunoCiccarino/GopherLight/logger"
	"github.com/BrunoCiccarino/GopherLight/plugins"
	"github.com/BrunoCiccarino/GopherLight/req"
)

// Middleware defines a function signature for middleware.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// App represents the core web application structure, managing routes, middlewares, and plugins.
type App struct {
	root        *Node
	middlewares []Middleware
	plugins     []plugins.Plugin
}

// NewApp creates a new App instance with an initialized root node.
// Returns:
//
//	*App: A new App instance.
func NewApp() *App {
	return &App{
		root: NewNode("/"),
	}
}

// Use adds a middleware function to the App's middleware stack.
// Args:
//
//	mw (Middleware): The middleware function to add.
func (a *App) Use(mw Middleware) {
	a.middlewares = append(a.middlewares, mw)
}

// AddPlugin adds a plugin to the App's plugin stack.
// Args:
//
//	p (plugins.Plugin): The plugin to add.
func (a *App) AddPlugin(p plugins.Plugin) {
	a.plugins = append(a.plugins, p)
}

// RegisterPlugins registers all plugins added to the App.
func (a *App) RegisterPlugins() {
	for _, plugin := range a.plugins {
		plugin.Register(a.Route)
	}
}

// Route registers a route for a specific HTTP method and path.
// Args:
//
//	method (string): The HTTP method (e.g., "GET").
//	path (string): The route path.
//	handler (req.Handler): The handler function for the route.
func (a *App) Route(method, path string, handler req.Handler) {
	segments := strings.Split(strings.Trim(path, "/"), "/")

	var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		request := req.NewRequest(r)
		response := req.NewResponse(w)
		handler(request, response)
	}

	for i := len(a.middlewares) - 1; i >= 0; i-- {
		h = a.middlewares[i](h)
	}

	a.root.AddRoute(append([]string{method}, segments...), h)
}

// Get registers a handler for the GET HTTP method.
// Args:
//
//	path (string): The route path.
//	handler (req.Handler): The handler function.
func (a *App) Get(path string, handler req.Handler) {
	a.Route(http.MethodGet, path, handler)
}

// Post registers a handler for the POST HTTP method.
// Args:
//
//	path (string): The route path.
//	handler (req.Handler): The handler function.
func (a *App) Post(path string, handler req.Handler) {
	a.Route(http.MethodPost, path, handler)
}

// Put registers a handler for the PUT HTTP method.
// Args:
//
//	path (string): The route path.
//	handler (req.Handler): The handler function.
func (a *App) Put(path string, handler req.Handler) {
	a.Route(http.MethodPut, path, handler)
}

// Delete registers a handler for the DELETE HTTP method.
// Args:
//
//	path (string): The route path.
//	handler (req.Handler): The handler function.
func (a *App) Delete(path string, handler req.Handler) {
	a.Route(http.MethodDelete, path, handler)
}

// Patch registers a handler for the PATCH HTTP method.
// Args:
//
//	path (string): The route path.
//	handler (req.Handler): The handler function.
func (a *App) Patch(path string, handler req.Handler) {
	a.Route(http.MethodPatch, path, handler)
}

// Options registers a handler for the OPTIONS HTTP method.
// Args:
//
//	path (string): The route path.
//	handler (req.Handler): The handler function.
func (a *App) Options(path string, handler req.Handler) {
	a.Route(http.MethodOptions, path, handler)
}

// Head registers a handler for the HEAD HTTP method.
// Args:
//
//	path (string): The route path.
//	handler (req.Handler): The handler function.
func (a *App) Head(path string, handler req.Handler) {
	a.Route(http.MethodHead, path, handler)
}

// Connect registers a handler for the CONNECT HTTP method.
// Args:
//
//	path (string): The route path.
//	handler (req.Handler): The handler function.
func (a *App) Connect(path string, handler req.Handler) {
	a.Route(http.MethodConnect, path, handler)
}

// Trace registers a handler for the TRACE HTTP method.
// Args:
//
//	path (string): The route path.
//	handler (req.Handler): The handler function.
func (a *App) Trace(path string, handler req.Handler) {
	a.Route(http.MethodTrace, path, handler)
}

// ServeHTTP handles incoming HTTP requests and dispatches them to the appropriate route.
// Args:
//
//	w (http.ResponseWriter): The response writer.
//	r (*http.Request): The incoming request.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	fullPath := append([]string{r.Method}, pathSegments...)

	handler, routeExists := a.root.FindRoute(fullPath)

	if routeExists {
		handler(w, r)
		return
	}

	for method := range httpMethods {
		alternatePath := append([]string{method}, pathSegments...)
		_, exists := a.root.FindRoute(alternatePath)
		if exists {
			w.Header().Set("Allow", strings.Join(allowedMethods(pathSegments, a.root), ", "))
			http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
	}

	http.NotFound(w, r)
}

var httpMethods = map[string]struct{}{
	http.MethodGet:     {},
	http.MethodPost:    {},
	http.MethodPut:     {},
	http.MethodDelete:  {},
	http.MethodPatch:   {},
	http.MethodOptions: {},
	http.MethodHead:    {},
	http.MethodConnect: {},
	http.MethodTrace:   {},
}

// allowedMethods returns the HTTP methods allowed for a specific route path.
// Args:
//
//	pathSegments ([]string): The segments of the route path.
//	root (*Node): The root node of the route tree.
//
// Returns:
//
//	[]string: A list of allowed HTTP methods.
func allowedMethods(pathSegments []string, root *Node) []string {
	allowed := []string{}
	for method := range httpMethods {
		alternatePath := append([]string{method}, pathSegments...)
		_, exists := root.FindRoute(alternatePath)
		if exists {
			allowed = append(allowed, method)
		}
	}
	return allowed
}

// Listen starts the HTTP server on the specified address and handles graceful shutdown.
// Args:
//
//	addr (string): The address to listen on (e.g., ":8080").
//
// Returns:
//
//	error: An error if the server fails to start or shutdown.
func (a *App) Listen(addr string) error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(stop)

	srv := &http.Server{
		Addr:    addr,
		Handler: a,
	}
	serverError := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverError <- err
		}
	}()

	select {
	case <-stop:
		logger.LogInfo("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("server shutdown failed: %w", err)
		}
		logger.LogInfo("Server gracefully stopped.")
		return nil
	case err := <-serverError:
		return fmt.Errorf("server error: %w", err)
	}
}
