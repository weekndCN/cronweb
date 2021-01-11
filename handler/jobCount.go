package handler

import (
	"net/http"

	"github.com/weekndCN/cronweb/jobs"
	"github.com/weekndCN/cronweb/logger"
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
			InternalError(w, err)
			return
		}
		data := &Count{Num: num}
		JSON(w, data, 200)
	}
}
