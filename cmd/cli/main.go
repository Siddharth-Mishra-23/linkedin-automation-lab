package main

import (
	"context"
	"fmt"
    "time"

	"github.com/Siddharth-Mishra-23/linkedin-automation-lab/internal/browser"
	"github.com/Siddharth-Mishra-23/linkedin-automation-lab/internal/stealth"
)

func main() {
	fmt.Println("🚀 LinkedIn Automation Lab (CLI)")
	fmt.Println("Mode: Mock Browser")

	ctx := context.Background()

	stealthEngine := stealth.NewEngine()
	stealthEngine.Register(&stealth.MouseMovementPlugin{})
	stealthEngine.Register(&stealth.TimingPlugin{})
	stealthEngine.Register(&stealth.FingerprintPlugin{})
	stealthEngine.Register(&stealth.ScrollPlugin{})
	stealthEngine.Register(&stealth.TypingPlugin{})
	stealthEngine.Register(&stealth.HoverPlugin{})
	stealthEngine.Register(&stealth.SchedulePlugin{
		StartHour: 9,
		EndHour:   18,
	})
	stealthEngine.Register(&stealth.RateLimitPlugin{
		Delay: 800 * time.Millisecond,
	})


	if err := stealthEngine.ApplyAll(ctx); err != nil {
		fmt.Println("Stealth engine failed:", err)
		return
	}

	br := &browser.MockBrowser{}

	if err := br.Start(ctx); err != nil {
		fmt.Println("Failed to start browser:", err)
		return
	}
	defer br.Close()

	page, err := br.NewPage(ctx)
	if err != nil {
		fmt.Println("Failed to create page:", err)
		return
	}

	page.Goto("https://www.linkedin.com/login")
	page.Type("#username", "demo@example.com")
	page.Type("#password", "********")
	page.Click("button[type=submit]")
	page.Wait(1000)

	fmt.Println("✅ Mock automation flow completed")
}
