package util

import (
	"encoding/json"
	"net/http"
)

type ResponseBadRequest struct {
	RequestID    string `json:"request_id"`
	StatusCode   int    `json:"status_code"`
	ErrorMessage any    `json:"error_message"`
}

const (
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
)

func Response[T any](write http.ResponseWriter, status int, data T) {
	write.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	write.WriteHeader(status)
	json.NewEncoder(write).Encode(data)
}

func ResponseBadRequestError[T any](write http.ResponseWriter, status int, requestID string, data T) {
	write.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	write.WriteHeader(status)
	result := ResponseBadRequest{
		RequestID:    requestID,
		StatusCode:   status,
		ErrorMessage: data,
	}
	json.NewEncoder(write).Encode(result)
}
