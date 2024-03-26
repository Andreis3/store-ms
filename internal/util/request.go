package util

import (
	"encoding/json"
	"net/http"
)

func RecoverBody[T any](req *http.Request) (*T, error) {
	var result T
	err := json.NewDecoder(req.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
