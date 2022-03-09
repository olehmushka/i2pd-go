package handlederror

import (
	"context"
	"fmt"
	contexttools "i2pdgo/core/context-tools"
	traceid "i2pdgo/core/trace-id"
)

type HandledError struct {
	Err error
	Msg string

	Code    *Code
	TraceID *traceid.TraceID
}

func Handle(ctx context.Context, err error, msg string, code *Code) error {
	return &HandledError{
		Err:     err,
		Msg:     msg,
		TraceID: contexttools.GetTraceID(ctx),
		Code:    code,
	}
}

func (e *HandledError) Error() string {
	errMsg := "error: <nil>;"
	if e.Err != nil {
		errMsg = fmt.Sprintf("error: %s;", e.Err.Error())
	}

	return fmt.Sprintf("%s msg: %s;", errMsg, e.Msg)
}

func ToHandledError(ctx context.Context, err error) *HandledError {
	if err == nil {
		return nil
	}

	he, ok := err.(*HandledError)
	if !ok {
		return &HandledError{
			Err:     err,
			Msg:     err.Error(),
			TraceID: contexttools.GetTraceID(ctx),
			Code:    InternalErrorCode,
		}
	}

	return he
}

func HandleInternalError(ctx context.Context, err error, msg string) error {
	return Handle(ctx, err, msg, InternalErrorCode)
}
