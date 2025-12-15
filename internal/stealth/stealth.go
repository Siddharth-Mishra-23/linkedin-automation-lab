package stealth

import "context"

type Plugin interface {
	Name() string
	Apply(ctx context.Context) error
}

type Engine interface {
	Register(plugin Plugin)
	ApplyAll(ctx context.Context) error
}
