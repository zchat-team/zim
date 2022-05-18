package server

import (
	"context"
)

type Options struct {
	Id       string
	TcpAddr  string
	WsAddr   string
	Context  context.Context
	NatsAddr string
}

func NewOptions(opts ...Option) Options {
	options := Options{
		//RpcLogic: "zim.logic",
	}

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

func TcpAddr(addr string) Option {
	return func(o *Options) {
		o.TcpAddr = addr
	}
}

func WsAddr(addr string) Option {
	return func(o *Options) {
		o.WsAddr = addr
	}
}

func NatsAddr(addr string) Option {
	return func(o *Options) {
		o.NatsAddr = addr
	}
}
