package util

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type ResponseBadRequest struct {
	ID           string `json:"id"`
	StatusCode   int    `json:"status_code"`
	ErrorMessage any    `json:"error_message"`
}

func Response[T any](write http.ResponseWriter, status int, data T) {
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(status)
	json.NewEncoder(write).Encode(data)
}

func ResponseBadRequestError[T any](write http.ResponseWriter, status int, data T) {
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(status)
	result := ResponseBadRequest{
		ID:           uuid.UUID.String(uuid.New()),
		StatusCode:   status,
		ErrorMessage: data,
	}
	json.NewEncoder(write).Encode(result)
}
