package jobs

import "github.com/robfig/cron/v3"

type (
	// Job .
	Job struct {
		ID        cron.EntryID `json:"id,omitempty"`
		Name      string       `json:"name"`
		Desc      string       `json:"desc"`
		Scheduler string       `json:"scheduler"`
		Created   int64        `json:"created,omitempty"`
		Action    string       `json:"action"`
	}

	// JobCron cron jobs abstract concept
	JobCron interface {
		Find(*cron.Cron, string) (*Job, error)
		Add(*cron.Cron, Job) error
		Delete(*cron.Cron, string) error
		Update(*cron.Cron, string) error
		List() ([]Job, error)
		Count() (int, error)
	}
)
