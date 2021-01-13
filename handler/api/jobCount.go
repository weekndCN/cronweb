package api

import (
	"net/http"

	req "github.com/weekndCN/rw-cron/handler/request"
	"github.com/weekndCN/rw-cron/jobs"
	"github.com/weekndCN/rw-cron/logger"
)

// Count job numbers
type Count struct {
	Num int `json:"num"`
}

// HandleCount Count jobs
func HandleCount(jobs jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		num, err := jobs.Count()
		if err != nil {
			logger.FromRequest(r).WithError(err).Debugln("job count failed")
			req.InternalError(w, err)
			return
		}
		data := &Count{Num: num}
		req.JSON(w, data, 200)
	}
}
