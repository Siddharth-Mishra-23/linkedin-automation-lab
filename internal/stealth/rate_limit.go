package stealth

import (
	"context"
	"fmt"
	"time"
)

type RateLimitPlugin struct {
	Delay time.Duration
}

func (r *RateLimitPlugin) Name() string {
	return "Rate Limiting & Throttling"
}

func (r *RateLimitPlugin) Apply(ctx context.Context) error {
	fmt.Println("[Stealth][RateLimit] Cooling down for", r.Delay)
	time.Sleep(r.Delay)
	return nil
}
