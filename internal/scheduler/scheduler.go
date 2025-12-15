package scheduler

import (
	"fmt"
	"time"
)

type Scheduler struct {
	StartHour int
	EndHour   int
}

func New(start, end int) *Scheduler {
	return &Scheduler{
		StartHour: start,
		EndHour:   end,
	}
}

func (s *Scheduler) CanRun() bool {
	hour := time.Now().Hour()
	return hour >= s.StartHour && hour <= s.EndHour
}

func (s *Scheduler) WaitUntilAllowed() {
	for !s.CanRun() {
		fmt.Println("[Scheduler] Outside allowed hours, waiting...")
		time.Sleep(2 * time.Second)
	}
}
