package stealth

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type TimingPlugin struct{}

func (t *TimingPlugin) Name() string {
	return "Randomized Timing"
}

func (t *TimingPlugin) Apply(ctx context.Context) error {
	delay := time.Duration(rand.Intn(500)+300) * time.Millisecond
	fmt.Println("[Stealth][Timing] Waiting for", delay)
	time.Sleep(delay)
	return nil
}
