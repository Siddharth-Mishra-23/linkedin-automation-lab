package stealth

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type ScrollPlugin struct{}

func (s *ScrollPlugin) Name() string {
	return "Random Scrolling Behavior"
}

func (s *ScrollPlugin) Apply(ctx context.Context) error {
	scrolls := rand.Intn(4) + 2
	for i := 0; i < scrolls; i++ {
		distance := rand.Intn(400) + 100
		fmt.Printf("[Stealth][Scroll] Scrolling %d px\n", distance)
		time.Sleep(time.Duration(rand.Intn(300)+200) * time.Millisecond)
	}
	return nil
}
