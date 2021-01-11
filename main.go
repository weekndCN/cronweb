package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/robfig/cron/v3"
	"github.com/weekndCN/cronweb/handler"
	"github.com/weekndCN/cronweb/jobs"
)

// initLogging can pass config inital value if need
func initLogging() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	// init log
	initLogging()
	c := cron.New()
	jobs := jobs.New()
	r := handler.NewAPI(c, jobs)
	c.Start()
	defer c.Stop()
	log.Fatal(http.ListenAndServe(":9090", r.Handler()))
}
