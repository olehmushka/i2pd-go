package config

import "context"

type config struct {
	cfg *ConfigStruct
}

func New() Config {
	return &config{}
}

func (c *config) Init(ctx context.Context) error {
	c.cfg = &ConfigStruct{}

	return nil
}

func (c *config) ParseConfig(ctx context.Context, path string) error {
	return nil
}

func (c *config) GetOption(ctx context.Context, optName string, value interface{}) error {
	return nil
}

func (c *config) SetOption(ctx context.Context, optName string, value interface{}) error {
	return nil
}

func (c *config) IsDefault(ctx context.Context, optName string) error {
	return nil
}
