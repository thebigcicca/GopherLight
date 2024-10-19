package router

import (
	"github.com/BrunoCiccarino/express-go/req"
)

type Route struct {
	Path    string
	Handler func(req *req.Request, res *req.Response)
}

func NewRoute(path string, handler func(req *req.Request, res *req.Response)) *Route {
	return &Route{
		Path:    path,
		Handler: handler,
	}
}
