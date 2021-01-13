package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/weekndCN/rw-cron/jobs"
)

// JobCount metric job count
func JobCount(c jobs.JobCron) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "cron_jobs_count",
			Help: "Total number of jobs.",
		}, func() float64 {
			num, _ := c.Count()
			return float64(num)
		}),
	)
}
