package router

import (
	"express-go/req"
	"net/http"
)

type App struct {
	routes map[string]func(req *req.Request, res *req.Response)
}

func NewApp() *App {
	return &App{
		routes: make(map[string]func(req *req.Request, res *req.Response)),
	}
}

func (a *App) Route(path string, handler func(req *req.Request, res *req.Response)) {
	a.routes[path] = handler
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, exists := a.routes[r.URL.Path]; exists {
		request := req.NewRequest(r)
		response := req.NewResponse(w)
		handler(request, response)
	} else {
		http.NotFound(w, r)
	}
}

func (a *App) Listen(addr string) error {
	return http.ListenAndServe(addr, a)
}
