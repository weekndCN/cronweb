package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/weekndCN/cronweb/jobs"
)

// HandleCount Count jobs
func HandleCount(jobs jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("Content-Type", "application/json")
		num, err := jobs.Count()

		if err != nil {
			fmt.Println(err)
		}

		// json Output
		data := struct {
			Num int `json:"num"`
		}{Num: num}
		msg, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(msg)
	}
}
