package handler

import (
	"net/http"

	"github.com/robfig/cron/v3"
)

// HandleDelete delete a job api
func HandleDelete(c *cron.Cron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
