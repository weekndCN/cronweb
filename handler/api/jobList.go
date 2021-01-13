package api

import (
	"net/http"

	req "github.com/weekndCN/rw-cron/handler/request"
	"github.com/weekndCN/rw-cron/jobs"
	"github.com/weekndCN/rw-cron/logger"
)

// HandleList list all jobs
func HandleList(event jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := event.List()
		if err != nil {
			logger.FromRequest(r).WithError(err).Debugln("list jobs失败")
			req.InternalError(w, err)
			return
		}

		req.JSON(w, tasks, 200)
	}
}
