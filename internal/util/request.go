package util

import (
	"encoding/json"
	"net/http"
)

func RecoverBody[T any](req *http.Request) (T, error) {
	var result T
	err := json.NewDecoder(req.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
