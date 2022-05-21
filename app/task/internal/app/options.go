package app

type BeforeFunc func() error

type Options struct {
	Before BeforeFunc
}

func newOptions(opts ...Option) Options {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	return options
}

type Option func(*Options)

func Before(f BeforeFunc) Option {
	return func(o *Options) {
		o.Before = f
	}
}
