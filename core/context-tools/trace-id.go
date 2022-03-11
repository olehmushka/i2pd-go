package contexttools

import (
	traceid "i2pdgo/core/trace-id"

	"context"
	"fmt"
)

type traceIDKey int

const (
	traceIDCtxKey traceIDKey = iota + 100
)

func AppendTraceID(ctx context.Context, traceID *traceid.TraceID) context.Context {
	return context.WithValue(ctx, traceIDCtxKey, traceID.Value())
}

func GetTraceID(ctx context.Context) *traceid.TraceID {
	id := ctx.Value(traceIDCtxKey)
	if id == nil {
		return traceid.GenerateEmptyTraceID()
	}
	return traceid.GenerateTraceIDByID(fmt.Sprintf("%v", id))
}
