package server

type Options struct {
	NatsAddr string
}

func NewOptions(opts ...Option) Options {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	return options
}

type Option func(*Options)

func NatsAddr(addr string) Option {
	return func(o *Options) {
		o.NatsAddr = addr
	}
}
