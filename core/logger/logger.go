package logger

import (
	"context"
	contexttools "i2pdgo/core/context-tools"

	log "github.com/sirupsen/logrus"
)

func New(ctx context.Context) *log.Entry {
	traceID := contexttools.GetTraceID(ctx)

	e := contexttools.GetEnv(ctx)
	if e.IsDev() || e.IsEmpty() || e.IsTest() || e.IsUnknown() {
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}

	return log.WithFields(log.Fields{
		TraceIDField: traceID.Value(),
	})
}
