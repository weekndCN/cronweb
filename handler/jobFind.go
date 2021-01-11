package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
	"github.com/weekndCN/cronweb/jobs"
)

// HandleFind find a job api
func HandleFind(c *cron.Cron, event jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name, ok := vars["name"]

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request or Bad Parameter"))
			return
		}

		job, err := event.Find(c, name)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		data, err := json.Marshal(*job)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Marshal failed"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
