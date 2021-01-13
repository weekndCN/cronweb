package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
	"github.com/weekndCN/rw-cron/handler/api/jobAdd"
	"github.com/weekndCN/rw-cron/handler/api/jobCount"
	"github.com/weekndCN/rw-cron/handler/api/jobDelete"
	"github.com/weekndCN/rw-cron/handler/api/jobList"
	"github.com/weekndCN/rw-cron/handler/api/jobFind"
	"github.com/weekndCN/rw-cron/handler/api/jobLog"
	"github.com/weekndCN/rw-cron/logger"
)

// Server job server
type Server struct {
	cron    *cron.Cron
	jobCron jobs.JobCron
}

// NewAPI .
func NewAPI(cron *cron.Cron, jobs jobs.JobCron) Server {
	return Server{
		cron:    cron,
		jobCron: jobs,
	}
}

// Handler endpoints handler
func (s Server) Handler() http.Handler {
	r := mux.NewRouter()
	r.Use(logger.Middleware)
	r.HandleFunc("/list", jobList.HandleList(s.jobCron)).Methods("GET")
	r.HandleFunc("/count", jobCount.HandleCount(s.jobCron)).Methods("GET")
	r.HandleFunc("/{name}", jobFind.HandleFind(s.cron, s.jobCron)).Methods("GET")
	r.HandleFunc("/{name}", jobDelete.HandleDelete(s.cron, s.jobCron)).Methods("DELETE")
	r.HandleFunc("/add", jobAdd.HandleAdd(s.cron, s.jobCron)).Methods("POST")
	r.HandleFunc("/update", jobUpdate.HandleUpdate(s.cron)).Methods("POST")
	r.HandleFunc("/health", HandleHealth()).Methods("GET")
	return r
}
