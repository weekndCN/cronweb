package handler

import (
	"net/http"

	"github.com/weekndCN/cronweb/jobs"
	"github.com/weekndCN/cronweb/logger"
)

// HandleList list all jobs
func HandleList(event jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := event.List()
		if err != nil {
			logger.FromRequest(r).WithError(err).Debugln("list jobs失败")
			InternalError(w, err)
			return
		}

		JSON(w, tasks, 200)
	}
}
