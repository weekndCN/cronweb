package main

import (
	"log"
	"net/http"

	"github.com/robfig/cron/v3"
	"github.com/weekndCN/cronweb/handler"
	"github.com/weekndCN/cronweb/jobs"
)

func main() {
	c := cron.New()
	jobs := jobs.New()
	r := handler.NewAPI(c, jobs)
	c.Start()
	defer c.Stop()
	log.Fatal(http.ListenAndServe(":9090", r.Handler()))
}
