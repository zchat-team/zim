package auth

type GenerateOptions struct {
	Metadata map[string]string
	Scopes   []string
	Issuer   string
	Type     string
}

type GenerateOption func(o *GenerateOptions)

func WithType(t string) GenerateOption {
	return func(o *GenerateOptions) {
		o.Type = t
	}
}

func WithMetadata(md map[string]string) GenerateOption {
	return func(o *GenerateOptions) {
		o.Metadata = md
	}
}

func WithIssuer(i string) GenerateOption {
	return func(o *GenerateOptions) {
		o.Issuer = i
	}
}

func WithScopes(s ...string) GenerateOption {
	return func(o *GenerateOptions) {
		o.Scopes = s
	}
}

func NewGenerateOptions(opts ...GenerateOption) GenerateOptions {
	var options GenerateOptions
	for _, o := range opts {
		o(&options)
	}
	return options
}

type RefreshTokenOptions struct {
	Metadata map[string]string
}

type RefreshTokenOption func(o *RefreshTokenOptions)

func Metadata(md map[string]string) RefreshTokenOption {
	return func(o *RefreshTokenOptions) {
		o.Metadata = md
	}
}

func NewRefreshTokenOptions(opts ...RefreshTokenOption) RefreshTokenOptions {
	var options RefreshTokenOptions
	for _, o := range opts {
		o(&options)
	}
	return options
}
