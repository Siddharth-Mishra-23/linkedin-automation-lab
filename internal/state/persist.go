package state

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type PersistentState struct {
	Date        string `json:"date"`
	Actions     int    `json:"actions"`
	MaxActions  int    `json:"max_actions"`
}

type Tracker struct {
	mu         sync.Mutex
	maxActions int
	state      PersistentState
}

const stateFile = "state.json"

func New(max int) *Tracker {
	t := &Tracker{maxActions: max}
	t.load()
	return t
}

func (t *Tracker) load() {
	t.mu.Lock()
	defer t.mu.Unlock()

	today := time.Now().Format("2006-01-02")

	file, err := os.ReadFile(stateFile)
	if err != nil {
		t.state = PersistentState{
			Date:       today,
			Actions:    0,
			MaxActions: t.maxActions,
		}
		t.save()
		return
	}

	_ = json.Unmarshal(file, &t.state)

	if t.state.Date != today {
		t.state.Date = today
		t.state.Actions = 0
		t.state.MaxActions = t.maxActions
		t.save()
	}
}

func (t *Tracker) save() {
	data, _ := json.MarshalIndent(t.state, "", "  ")
	_ = os.WriteFile(stateFile, data, 0644)
}

func (t *Tracker) CanPerform() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.state.Actions < t.state.MaxActions
}

func (t *Tracker) RecordAction() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.state.Actions++
	t.save()
}

func (t *Tracker) Snapshot() PersistentState {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.state
}
