package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/robfig/cron/v3"
	"github.com/weekndCN/rw-cron/jobs"
	"github.com/weekndCN/rw-cron/logger"
)

// HandleAdd add job api
func HandleAdd(c *cron.Cron, event jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var job jobs.Job
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			InternalError(w, err)
			logger.FromRequest(r).WithError(err).Debugln("无法解析body的内容")
			return
		}

		err = json.Unmarshal(data, &job)
		if err != nil {
			logger.FromRequest(r).WithError(err).Debugln("Json数据格式或者参数错误")
			BadRequest(w, err)
			return
		}

		if job.Name == "" || job.Scheduler == "" || job.Action == "" {
			logger.FromRequest(r).WithError(err).Debugln("Json数据格式或者参数错误")
			BadRequestf(w, "Json数据格式或者参数错误")
			return
		}

		err = event.Add(c, job)
		if err != nil {
			logger.FromRequest(r).WithError(err).Debugln("任务添加失败")
			InternalError(w, err)
			return
		}
		JSON(w, "任务添加成功", 200)
	}
}
