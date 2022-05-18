package server

type Options struct {
	Id       string
	TcpAddr  string
	WsAddr   string
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
