package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

// Server job server
type Server struct {
	cron *cron.Cron
}

// NewAPI .
func NewAPI(cron *cron.Cron) Server {
	return Server{
		cron: cron,
	}
}

// Handler endpoints handler
func (s Server) Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/list", HandleList(s.cron))
	r.HandleFunc("/find", HandleFind(s.cron))
	r.HandleFunc("/delete", HandleDelete(s.cron))
	r.HandleFunc("/add", HandleAdd(s.cron))
	r.HandleFunc("/Update", HandleUpdate(s.cron))
	return r
}
