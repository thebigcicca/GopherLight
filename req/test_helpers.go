// In req/test_helpers.go

package req

import (
	"io"
	"net/http/httptest"
)

func NewTestRequest(method, target string, body io.Reader) (*Request, *Response, *httptest.ResponseRecorder) {
	httpReq := httptest.NewRequest(method, target, body)
	httpRes := httptest.NewRecorder()
	req := NewRequest(httpReq)
	res := NewResponse(httpRes)
	return req, res, httpRes
}
