package plugins

import (
	"github.com/BrunoCiccarino/GopherLight/req"
)

type Plugin interface {
	Register(route func(method, path string, handler func(req *req.Request, res *req.Response)))
}
