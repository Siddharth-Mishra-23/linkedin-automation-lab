package strategy

import "time"

type Profile struct {
	Name             string
	MinDelay         time.Duration
	MaxDelay         time.Duration
	MaxActionsPerDay int
}

type Engine interface {
	Current() Profile
}
