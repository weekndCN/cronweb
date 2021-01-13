package api

import (
	"fmt"
	"net/http"

	req "github.com/weekndCN/rw-cron/handler/request"
)

// HandleLog fetch log latest
func HandleLog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req.NotImplemented(w, fmt.Errorf("NotImplemented"))
	}
}

// HandleLogs fetch logs specified range
func HandleLogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req.NotImplemented(w, fmt.Errorf("NotImplemented"))
	}
}
