package state

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type PersistentState struct {
	mu         sync.Mutex
	Date       string `json:"date"`
	Actions    int    `json:"actions"`
	MaxActions int    `json:"max_actions"`
	FilePath   string `json:"-"`
}

func NewPersistent(max int, path string) (*PersistentState, error) {
	s := &PersistentState{
		MaxActions: max,
		FilePath:   path,
	}

	if err := s.load(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *PersistentState) CanPerform() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	today := time.Now().Format("2006-01-02")

	if s.Date != today {
		s.Date = today
		s.Actions = 0
		s.save()
	}

	return s.Actions < s.MaxActions
}

func (s *PersistentState) RecordAction() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Actions++
	return s.save()
}

func (s *PersistentState) load() error {
	data, err := os.ReadFile(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			s.Date = time.Now().Format("2006-01-02")
			s.Actions = 0
			return s.save()
		}
		return err
	}

	return json.Unmarshal(data, s)
}

func (s *PersistentState) save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, data, 0644)
}
