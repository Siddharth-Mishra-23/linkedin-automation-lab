package stealth

import (
	"context"
	"fmt"
	"time"
)

type SchedulePlugin struct {
	StartHour int
	EndHour   int
}

func (s *SchedulePlugin) Name() string {
	return "Activity Scheduling"
}

func (s *SchedulePlugin) Apply(ctx context.Context) error {
	now := time.Now().Hour()
	if now < s.StartHour || now > s.EndHour {
		fmt.Println("[Stealth][Schedule] Outside business hours, delaying actions")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("[Stealth][Schedule] Within business hours")
	}
	return nil
}
