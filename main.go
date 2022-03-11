package main

import (
	"context"
	contexttools "i2pdgo/core/context-tools"
	"i2pdgo/core/env"
	"i2pdgo/core/logger"
	traceid "i2pdgo/core/trace-id"
	"i2pdgo/deamon"
	"os"
)

func initCtxValues() context.Context {
	ctx := context.Background()
	ctx = contexttools.AppendTraceID(ctx, traceid.GenerateTraceIDByID("system_trace_id"))
	ctx = contexttools.AppendEnv(ctx, env.CreateEnvByName(os.Getenv(env.EnvVar)))

	return ctx
}

func main() {
	ctx := initCtxValues()
	log := logger.New(ctx)

	d := deamon.New()

	if err := d.Init(ctx); err != nil {
		logger.LogError(ctx, log, err)
		if stopErr := d.Stop(ctx); stopErr != nil {
			logger.LogError(ctx, log, stopErr)
		}
		os.Exit(0)

		return
	}

	if err := d.Start(ctx); err != nil {
		logger.LogError(ctx, log, err)
		if stopErr := d.Stop(ctx); stopErr != nil {
			logger.LogError(ctx, log, stopErr)
		}
		os.Exit(1)

		return
	}

	if err := d.Stop(ctx); err != nil {
		logger.LogError(ctx, log, err)
		if stopErr := d.Stop(ctx); stopErr != nil {
			logger.LogError(ctx, log, stopErr)
		}
		os.Exit(1)

		return
	}

	os.Exit(0)
}
