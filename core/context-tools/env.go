package contexttools

import (
	"i2pdgo/core/env"

	"context"
	"fmt"
)

type envKey int

const (
	envCtxKey envKey = iota + 200
)

func AppendEnv(ctx context.Context, e *env.Env) context.Context {
	return context.WithValue(ctx, envCtxKey, e.Value())
}

func GetEnv(ctx context.Context) *env.Env {
	e := ctx.Value(envCtxKey)
	if e == nil {
		return env.CreateEmptyEnv()
	}
	return env.CreateEnvByName(fmt.Sprintf("%v", e))
}
