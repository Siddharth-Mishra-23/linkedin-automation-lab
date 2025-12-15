package stealth

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type HoverPlugin struct{}

func (h *HoverPlugin) Name() string {
	return "Mouse Hovering & Wandering"
}

func (h *HoverPlugin) Apply(ctx context.Context) error {
	hovers := rand.Intn(3) + 2
	for i := 0; i < hovers; i++ {
		fmt.Println("[Stealth][Hover] Hovering over element")
		time.Sleep(time.Duration(rand.Intn(400)+200) * time.Millisecond)
	}
	return nil
}
