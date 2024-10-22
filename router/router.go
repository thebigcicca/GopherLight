package router

import (
	"net/http"

	"github.com/BrunoCiccarino/GopherLight/req"
)

type App struct {
	routes map[string]map[string]func(req *req.Request, res *req.Response)
}

func NewApp() *App {
	return &App{
		routes: make(map[string]map[string]func(req *req.Request, res *req.Response)),
	}
}

func (a *App) Route(method string, path string, handler func(req *req.Request, res *req.Response)) {
	if a.routes[path] == nil {
		a.routes[path] = make(map[string]func(req *req.Request, res *req.Response))
	}
	a.routes[path][method] = handler
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handlers, exists := a.routes[r.URL.Path]; exists {
		if handler, methodExists := handlers[r.Method]; methodExists {
			request := req.NewRequest(r)
			response := req.NewResponse(w)
			handler(request, response)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.NotFound(w, r)
	}
}

func (a *App) Listen(addr string) error {
	return http.ListenAndServe(addr, a)
}
