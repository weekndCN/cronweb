package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
	"github.com/weekndCN/cronweb/jobs"
	"github.com/weekndCN/cronweb/logger"
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
		data, err := json.Marshal(*job)
		if err != nil {
			logger.FromRequest(r).WithError(err).Debugln("json转换失败")
			InternalError(w, err)
			return
		}

		JSON(w, data, 200)
	}
}
