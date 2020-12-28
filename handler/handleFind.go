package handler

import (
	"fmt"
	"net/http"

	"github.com/robfig/cron/v3"
	"github.com/weekndCN/cronweb/jobs"
)

// HandleFind find a job api
func HandleFind(c *cron.Cron, event jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jobName = "佳能定时任务"
		job, err := event.Find(c, jobName)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(job)
	}
}
