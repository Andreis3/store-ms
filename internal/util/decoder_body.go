package util

import (
	"encoding/json"
	"errors"
	"net/http"
)

func DecoderBodyRequest[T any](req *http.Request) (T, *ValidationError) {
	var result T
	var jsonUnmarshalTypeError *json.UnmarshalTypeError
	var jsonSyntaxError *json.SyntaxError
	err := json.NewDecoder(req.Body).Decode(&result)
	switch {
	case errors.As(err, &jsonSyntaxError):
		return result, &ValidationError{
			Code:        "DJ-400",
			Status:      http.StatusBadRequest,
			ClientError: []string{"invalid json syntax"},
			LogError:    []string{jsonSyntaxError.Error()},
		}
	case errors.As(err, &jsonUnmarshalTypeError):
		return result, &ValidationError{
			Code:        "DJ-401",
			Status:      http.StatusBadRequest,
			ClientError: []string{"invalid json field type"},
			LogError:    []string{jsonUnmarshalTypeError.Error()},
		}
	case err != nil:
		return result, &ValidationError{
			Code:        "DJ-402",
			Status:      http.StatusBadRequest,
			LogError:    []string{err.Error()},
			ClientError: []string{"invalid json"},
		}
	}
	return result, nil
}
