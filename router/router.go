package router

import (
	"net/http"

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

func (a *App) Route(method string, path string, handler func(req *req.Request, res *req.Response)) {
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

func (a *App) Get(path string, handler func(req *req.Request, res *req.Response)) {
	a.Route(http.MethodGet, path, handler)
}

func (a *App) Post(path string, handler func(req *req.Request, res *req.Response)) {
	a.Route(http.MethodPost, path, handler)
}

func (a *App) Put(path string, handler func(req *req.Request, res *req.Response)) {
	a.Route(http.MethodPut, path, handler)
}

func (a *App) Delete(path string, handler func(req *req.Request, res *req.Response)) {
	a.Route(http.MethodDelete, path, handler)
}

func (a *App) Patch(path string, handler func(req *req.Request, res *req.Response)) {
	a.Route(http.MethodPatch, path, handler)
}
func (a *App) Options(path string, handler func(req *req.Request, res *req.Response)) {
	a.Route(http.MethodOptions, path, handler)
}
func (a *App) Head(path string, handler func(req *req.Request, res *req.Response)) {
	a.Route(http.MethodHead, path, handler)
}

func (a *App) Connect(path string, handler func(req *req.Request, res *req.Response)) {
	a.Route(http.MethodConnect, path, handler)
}
func (a *App) Trace(path string, handler func(req *req.Request, res *req.Response)) {
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
	return http.ListenAndServe(addr, a)
}
