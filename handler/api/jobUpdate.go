package api

import (
	"fmt"
	"net/http"

	"github.com/robfig/cron/v3"
	req "github.com/weekndCN/rw-cron/handler/request"
)

// HandleUpdate update a exsit job
func HandleUpdate(c *cron.Cron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req.NotImplemented(w, fmt.Errorf("NotImplemented"))
	}
}
