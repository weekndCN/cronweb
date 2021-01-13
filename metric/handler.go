package metric

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server metric Server
type Server struct {
	metrics http.Handler
}

// NewServer return a metric server instance
func NewServer() *Server {
	return &Server{
		metrics: promhttp.Handler(),
	}
}

// ServeHTTP responds to an http.Request and writes system
// metrics to the response body in plain text format.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.metrics.ServeHTTP(w, r)
}
