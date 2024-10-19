package req

import (
	"io"
	"net/http"
)

type Request struct {
	Req  *http.Request
	Body string
}

func NewRequest(req *http.Request) *Request {

	bodyBytes, _ := io.ReadAll(req.Body)
	bodyString := string(bodyBytes)

	return &Request{
		Req:  req,
		Body: bodyString,
	}
}

func (r *Request) QueryParam(key string) string {
	params := r.Req.URL.Query()
	return params.Get(key)
}

func (r *Request) Header(key string) string {
	return r.Req.Header.Get(key)
}

func (r *Request) BodyAsString() string {
	return r.Body
}
