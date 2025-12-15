package stealth

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type TypingPlugin struct{}

func (t *TypingPlugin) Name() string {
	return "Realistic Typing Simulation"
}

func (t *TypingPlugin) Apply(ctx context.Context) error {
	chars := rand.Intn(8) + 5
	for i := 0; i < chars; i++ {
		fmt.Println("[Stealth][Typing] Typing character")
		time.Sleep(time.Duration(rand.Intn(200)+80) * time.Millisecond)
	}
	return nil
}
