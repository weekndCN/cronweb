package jobs

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

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
func (c *Jobs) List() ([]Job, error) {
	task := []Job{}
	c.Lock()
	for _, v := range c.Tasks {
		task = append(task, v)
	}
	c.Unlock()
	return task, nil
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

	if _, ok := c.Tasks[job.Name]; ok {
		return fmt.Errorf("不能存在重复Job名称")
	}

	job.Created = time.Now().Unix()

	id, err := cron.AddFunc(job.Scheduler, func() {
		log.Println(job.Action)
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
