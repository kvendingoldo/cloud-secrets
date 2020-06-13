package provider

type Provider interface {
	GetSecret(name string)
}

type BaseProvider struct {
}