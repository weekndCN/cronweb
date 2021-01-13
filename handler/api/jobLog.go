package api

import (
	"fmt"
	"net/http"
)

// HandleLog fetch log latest
func HandleLog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NotImplemented(w, fmt.Errorf("NotImplemented"))
	}
}

// HandleLogs fetch logs specified range
func HandleLogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NotImplemented(w, fmt.Errorf("NotImplemented"))
	}
}
