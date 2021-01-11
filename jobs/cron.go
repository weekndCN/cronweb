package jobs

import (
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/weekndCN/cronweb/dingtalk"
)

var log = logrus.New()

// Jobs store all jobs in-memory
type Jobs struct {
	sync.Mutex
	Tasks map[string]Job
}

// New return a new job hub
func New() JobCron {
	return &Jobs{
		Tasks: make(map[string]Job),
	}
}

// List methods list all jobs
func (c *Jobs) List() (*[]Job, error) {
	task := []Job{}
	c.Lock()
	for _, v := range c.Tasks {
		task = append(task, v)
	}
	c.Unlock()
	return &task, nil
}

// Find find a job
func (c *Jobs) Find(cron *cron.Cron, jobName string) (*Job, error) {
	// job datastore search
	task, ok := c.Tasks[jobName]
	if !ok {
		return nil, fmt.Errorf("%s not in-memory", jobName)
	}

	entries := cron.Entries()
	for i := range entries {
		if task.ID == entries[i].ID {
			return &task, nil
		}
	}

	return nil, fmt.Errorf("%s not in jobs queue", jobName)
}

// Add  add a job
func (c *Jobs) Add(cron *cron.Cron, job Job) error {
	robot := dingtalk.NewRobot(fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", job.Webhook))
	if _, ok := c.Tasks[job.Name]; ok {
		return fmt.Errorf("不能存在重复Job名称")
	}

	job.Created = time.Now().Unix()
	id, err := cron.AddFunc(job.Scheduler, func() {
		msg, err := HTTPGet(job.Action)
		if err != nil {
			log.Println(err)
		}
		level := "Info"
		// if http request time(gap time) greater than job timeout
		if msg.Duration > job.Timeout && err == nil {
			level = "Warning"
		}

		if err != nil || msg.StatusCode >= 400 {
			level = "Error"
		}

		if err == nil && 200 <= msg.StatusCode && msg.StatusCode <= 300 {
			level = "Success"
		}

		switch job.Alert {
		case "Always":
			text := dingtalk.MsgText("#0DAD51", job.Name, msg.Name, level, msg.Start.Format(time.UnixDate), string(msg.Body), msg.Duration, msg.StatusCode)
			if err != nil {
				text = dingtalk.MsgText("#FF0000", job.Name, msg.Name, level, msg.Start.Format(time.UnixDate), err.Error(), msg.Duration, msg.StatusCode)
			}
			robot.SendMarkdown(msg.Name, text, nil, false)
		case "Success":
			if level == "Success" {
				text := dingtalk.MsgText("#0DAD51", job.Name, msg.Name, level, msg.Start.Format(time.UnixDate), string(msg.Body), msg.Duration, msg.StatusCode)
				robot.SendMarkdown(msg.Name, text, nil, false)
			}
		case "Failed":
			if level == "Error" {
				text := dingtalk.MsgText("#FF0000", job.Name, msg.Name, level, msg.Start.Format(time.UnixDate), string(msg.Body), msg.Duration, msg.StatusCode)
				robot.SendMarkdown(msg.Name, text, nil, false)
			}
		default:
			break
		}

		log.Printf("job:%s task:%s run with %s statusCode: %d\n", job.Name, msg.Name, level, msg.StatusCode)
	})

	if err != nil {
		fmt.Println(err)
	}
	// set job id
	job.ID = id
	// add job to job in-memory data store
	c.Lock()
	c.Tasks[job.Name] = job
	c.Unlock()
	return nil
}

// Delete delete a job from job store
func (c *Jobs) Delete(cron *cron.Cron, name string) error {
	c.Lock()
	// remove from in-memory datastore if matched
	task, ok := c.Tasks[name]
	if ok {
		delete(c.Tasks, name)
	}
	jobID := task.ID
	c.Unlock()

	if !ok {
		return fmt.Errorf("job name not exists in jobs queue")
	}
	// remove from job queue
	cron.Remove(jobID)
	return nil
}

// Update update a job in current job store
func (c *Jobs) Update(cron *cron.Cron, name string) error {
	return nil
}

// Count retrun jobs number
func (c *Jobs) Count() (int, error) {
	c.Lock()
	num := len(c.Tasks)
	c.Unlock()
	return num, nil
}
