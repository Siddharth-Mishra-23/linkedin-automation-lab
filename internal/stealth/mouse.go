package stealth

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type MouseMovementPlugin struct{}

func (m *MouseMovementPlugin) Name() string {
	return "Human-like Mouse Movement"
}

func (m *MouseMovementPlugin) Apply(ctx context.Context) error {
	steps := rand.Intn(5) + 3
	for i := 0; i < steps; i++ {
		fmt.Println("[Stealth][Mouse] Moving cursor naturally")
		time.Sleep(time.Duration(rand.Intn(200)+100) * time.Millisecond)
	}
	return nil
}
