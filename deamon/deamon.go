package deamon

import (
	"context"
	"i2pdgo/config"
	"i2pdgo/core/logger"
)

type deamon struct {
	cfg config.Config
}

func New() Deamon {
	return &deamon{
		cfg: config.New(),
	}
}

func (d *deamon) Init(ctx context.Context) error {
	log := logger.New(ctx)
	log.Info("initing deamon...")

	if err := d.cfg.Init(ctx); err != nil {
		return err
	}

	log.Info("deamon was initted")

	return nil
}

func (d *deamon) Start(ctx context.Context) error {
	log := logger.New(ctx)
	log.Info("starting deamon...")

	return nil
}

func (d *deamon) Stop(ctx context.Context) error {
	log := logger.New(ctx)
	log.Info("stopping deamon...")

	log.Info("deamon was stopped")

	return nil
}
