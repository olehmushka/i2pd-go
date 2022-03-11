package logger

import (
	"context"
	handlederror "i2pdgo/core/handled-error"

	log "github.com/sirupsen/logrus"
)

func LogError(ctx context.Context, l *log.Entry, err error) {
	he := handlederror.ToHandledError(ctx, err)
	l.WithFields(log.Fields{
		ErrorMessageField: he.Msg,
		ErrorCodeField:    he.Code,
	}).Error(he.Err)
}
