package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
	"github.com/weekndCN/rw-cron/jobs"
	"github.com/weekndCN/rw-cron/logger"
)

// HandleFind find a job api
func HandleFind(c *cron.Cron, event jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name, ok := vars["name"]

		if !ok {
			logger.FromRequest(r).WithError(ErrNotFound).Debugln("name参数不存在")
			BadRequestf(w, "name参数不存在")
			return
		}

		job, err := event.Find(c, name)
		if err != nil {
			logger.FromRequest(r).WithError(err).Debugln("job获取失败")
			BadRequest(w, err)
			return
		}

		JSON(w, job, 200)
	}
}
