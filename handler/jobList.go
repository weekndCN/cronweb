package handler

import (
	"encoding/json"
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

		data, err := json.Marshal(tasks)
		if err != nil {
			logger.FromRequest(r).WithError(err).Debugln("Json格式化失败")
			InternalError(w, err)
			return
		}
		JSON(w, data, 200)
	}
}
