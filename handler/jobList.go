package handler

import (
	"encoding/json"
	"net/http"

	"github.com/weekndCN/cronweb/jobs"
)

// HandleList list all jobs
func HandleList(event jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("Content-Type", "application/json")
		tasks, err := event.List()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		data, err := json.Marshal(tasks)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
