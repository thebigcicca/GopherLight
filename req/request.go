package req

import (
	"io"
	"net/http"

	"github.com/BrunoCiccarino/GopherLight/logger"
)

type Request struct {
	Req  *http.Request
	Body string
}

func NewRequest(req *http.Request) *Request {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		logger.LogError("Error reading request body: " + err.Error())
		return &Request{Req: req, Body: ""}
	}
	bodyString := string(bodyBytes)
	logger.LogInfo("Received request body: " + bodyString)

	return &Request{
		Req:  req,
		Body: bodyString,
	}
}

func (r *Request) QueryParam(key string) string {
	return r.Req.URL.Query().Get(key)
}

func (r *Request) Header(key string) string {
	return r.Req.Header.Get(key)
}

func (r *Request) BodyAsString() string {
	return r.Body
}
