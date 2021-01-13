package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
	"github.com/weekndCN/rw-cron/jobs"

	api "github.com/weekndCN/rw-cron/handler/api"
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
	r.HandleFunc("/list", api.HandleList(s.jobCron)).Methods("GET")
	r.HandleFunc("/count", api.HandleCount(s.jobCron)).Methods("GET")
	r.HandleFunc("/{name}", api.HandleFind(s.cron, s.jobCron)).Methods("GET")
	r.HandleFunc("/{name}", api.HandleDelete(s.cron, s.jobCron)).Methods("DELETE")
	r.HandleFunc("/add", api.HandleAdd(s.cron, s.jobCron)).Methods("POST")
	r.HandleFunc("/update", api.HandleUpdate(s.cron)).Methods("POST")
	r.HandleFunc("/health", api.HandleHealth()).Methods("GET")
	return r
}
