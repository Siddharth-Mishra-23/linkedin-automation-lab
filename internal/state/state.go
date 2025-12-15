package state

import "sync"

type State struct {
	mu          sync.Mutex
	actionsDone int
	maxActions  int
}

func New(max int) *State {
	return &State{maxActions: max}
}

func (s *State) CanPerform() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.actionsDone < s.maxActions
}

func (s *State) RecordAction() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.actionsDone++
}
