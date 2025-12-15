package stealth

import (
	"context"
	"fmt"
)

type DefaultEngine struct {
	plugins []Plugin
}

func NewEngine() *DefaultEngine {
	return &DefaultEngine{
		plugins: []Plugin{},
	}
}

func (e *DefaultEngine) Register(plugin Plugin) {
	e.plugins = append(e.plugins, plugin)
}

func (e *DefaultEngine) ApplyAll(ctx context.Context) error {
	for _, plugin := range e.plugins {
		fmt.Println("[Stealth] Applying:", plugin.Name())
		if err := plugin.Apply(ctx); err != nil {
			return err
		}
	}
	return nil
}
