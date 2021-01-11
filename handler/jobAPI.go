package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
	"github.com/weekndCN/cronweb/jobs"
	"github.com/weekndCN/cronweb/logger"
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
	r.HandleFunc("/list", HandleList(s.jobCron)).Methods("GET")
	r.HandleFunc("/count", HandleCount(s.jobCron)).Methods("GET")
	r.HandleFunc("/{name}", HandleFind(s.cron, s.jobCron)).Methods("GET")
	r.HandleFunc("/{name}", HandleDelete(s.cron, s.jobCron)).Methods("DELETE")
	r.HandleFunc("/add", HandleAdd(s.cron, s.jobCron)).Methods("POST")
	r.HandleFunc("/update", HandleUpdate(s.cron)).Methods("POST")
	return r
}
