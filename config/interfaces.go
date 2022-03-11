package config

import "context"

type Config interface {
	Init(ctx context.Context) error
	ParseConfig(ctx context.Context, path string) error
	GetOption(ctx context.Context, optName string, value interface{}) error
	SetOption(ctx context.Context, optName string, value interface{}) error
	IsDefault(ctx context.Context, optName string) error
}
