package app

import (
	"context"
)

type Options struct {
	Name     string
	Version  string
	Metadata map[string]string
	Context  context.Context
}

func newOptions(opts ...Option) Options {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	return options
}

type Option func(*Options)

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

func Version(version string) Option {
	return func(o *Options) {
		o.Version = version
	}
}

func Metadata(md map[string]string) Option {
	return func(o *Options) {
		o.Metadata = md
	}
}
