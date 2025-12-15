package strategy

import "time"

type DefaultEngine struct {
	profile Profile
}

func NewDefaultEngine(profile Profile) *DefaultEngine {
	return &DefaultEngine{profile: profile}
}

func (e *DefaultEngine) Current() Profile {
	return e.profile
}

func PredefinedProfiles() map[string]Profile {
	return map[string]Profile{
		"slow": {
			Name:             "Slow & Careful",
			MinDelay:         2 * time.Second,
			MaxDelay:         5 * time.Second,
			MaxActionsPerDay: 10,
		},
		"normal": {
			Name:             "Normal User",
			MinDelay:         1 * time.Second,
			MaxDelay:         3 * time.Second,
			MaxActionsPerDay: 20,
		},
		"aggressive": {
			Name:             "Aggressive Networker",
			MinDelay:         500 * time.Millisecond,
			MaxDelay:         1500 * time.Millisecond,
			MaxActionsPerDay: 35,
		},
	}
}
