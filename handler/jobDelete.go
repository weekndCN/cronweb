package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
	"github.com/weekndCN/cronweb/jobs"
)

// HandleDelete delete a job api
func HandleDelete(c *cron.Cron, event jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		name, ok := vars["name"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request or not passing name parameter"))
			return
		}

		if err := event.Delete(c, name); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Delete Job Success"))
	}
}
