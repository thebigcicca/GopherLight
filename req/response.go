package req

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w}
}

func (res *Response) Send(data string) {
	res.Write([]byte(data))
}

func (res *Response) Status(statusCode int) *Response {
	res.WriteHeader(statusCode)
	return res
}

func (res *Response) JSON(data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		res.Status(http.StatusInternalServerError).Send(`{"error": "Error encoding JSON"}`)
		return
	}
	res.Write(jsonData)
}
