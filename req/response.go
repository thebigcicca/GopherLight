package req

import (
	"encoding/json"
	"net/http"

	"github.com/BrunoCiccarino/GopherLight/logger"
)

type Response struct {
	http.ResponseWriter
	statusCodeWritten bool // This is to check if the status code has been written
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{ResponseWriter: w}
}

func (res *Response) Send(data string) {
	res.writeStatusIfNotWritten(http.StatusOK)
	res.Write([]byte(data))
}

func (res *Response) Status(statusCode int) *Response {
	res.writeStatusIfNotWritten(statusCode)
	return res
}

func (res *Response) JSON(data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.LogError("Error encoding JSON response: " + err.Error())
		res.Status(http.StatusInternalServerError).JSONError("Error encoding JSON")
		return
	}
	res.writeStatusIfNotWritten(http.StatusOK)
	res.Write(jsonData)
}

// New method to send standardized JSON error responses
func (res *Response) JSONError(message string) {
	res.Header().Set("Content-Type", "application/json")
	errorData := map[string]string{"error": message}
	jsonData, _ := json.Marshal(errorData)
	res.Write(jsonData)
}

// Helper method to write status code if not already written
func (res *Response) writeStatusIfNotWritten(statusCode int) {
	if !res.statusCodeWritten {
		res.WriteHeader(statusCode)
		res.statusCodeWritten = true
	}
}
