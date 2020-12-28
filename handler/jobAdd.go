package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/robfig/cron/v3"
	"github.com/weekndCN/cronweb/jobs"
)

// HandleAdd add job api
func HandleAdd(c *cron.Cron, event jobs.JobCron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var job jobs.Job
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("无法解析body的内容"))
			return
		}
		err = json.Unmarshal(data, &job)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Json数据格式或者参数错误"))
			return
		}

		err = event.Add(c, job)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("任务添加成功"))
	}
}
