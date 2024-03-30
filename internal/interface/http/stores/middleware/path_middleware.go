package stores_middleware

import (
	"fmt"
	"net/http"
)

func ValidatePath(next http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "1" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "id is required"}`))
			return
		}
		fmt.Println("Validating id: ", id)
		next.ServeHTTP(w, r)
	}
}
