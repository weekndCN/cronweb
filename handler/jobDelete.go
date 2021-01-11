package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
	"github.com/weekndCN/cronweb/jobs"
	"github.com/weekndCN/cronweb/logger"
)

// Res .
type Res struct {
	Result string `json:"res"`
}

// HandleDelete delete a job api
func HandleDelete(c *cron.Cron, event jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		name, ok := vars["name"]
		if !ok {
			logger.FromRequest(r).WithError(ErrNotFound).Debugln("name参数不存在")
			BadRequestf(w, "name参数不存在")
			return
		}

		if err := event.Delete(c, name); err != nil {
			InternalError(w, err)
			return
		}

		res := &Res{Result: "Delete Job Success"}

		JSON(w, res, 200)
	}
}
