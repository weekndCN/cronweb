package handler

import (
	"io"
	"net/http"
)

// HandleHealth handle system health
func HandleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "OK")
	}
}
