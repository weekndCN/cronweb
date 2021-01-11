package logger

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type loggerKey struct{}

// L alias logger entry
var L = logrus.NewEntry(logrus.StandardLogger())

// WithContext returns a new context with the provided logger. Use in
// combination with logger.WithField(s) for great effect.
func WithContext(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// FromContext receiver current logger from the context.
// if not available. use default
func FromContext(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey{})
	if logger == nil {
		return L
	}

	return logger.(*logrus.Entry)
}

// FromRequest receiver the current logger from the request.
// if no avalable. use default
func FromRequest(r *http.Request) *logrus.Entry {
	return FromContext(r.Context())
}
