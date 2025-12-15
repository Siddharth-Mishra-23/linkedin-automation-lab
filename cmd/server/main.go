package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Siddharth-Mishra-23/linkedin-automation-lab/internal/state"
)

var (
	currentStrategy = "normal"
	strategyMu      sync.Mutex
)

func main() {
	fmt.Println("🌐 LinkedIn Automation Lab — Web Dashboard")
	fmt.Println("Listening on http://localhost:8080")

	// -------------------------------------------------
	// Persistent State Tracker (shared across API)
	// -------------------------------------------------
	stateTracker := state.New(20)

	// -------------------------------------------------
	// Static UI
	// -------------------------------------------------
	fs := http.FileServer(http.Dir("./ui/web"))
	http.Handle("/", fs)

	// -------------------------------------------------
	// Health API
	// -------------------------------------------------
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	// -------------------------------------------------
	// Stats API (STEP 7.3)
	// -------------------------------------------------
	http.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(stateTracker.Snapshot())
	})

	// -------------------------------------------------
	// Strategy API (STEP 7.4)
	// -------------------------------------------------

	// Get current strategy
	http.HandleFunc("/api/strategy", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		strategyMu.Lock()
		defer strategyMu.Unlock()

		_ = json.NewEncoder(w).Encode(map[string]string{
			"strategy": currentStrategy,
		})
	})

	// Update strategy
	http.HandleFunc("/api/strategy/set", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var body struct {
			Strategy string `json:"strategy"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if body.Strategy != "slow" &&
			body.Strategy != "normal" &&
			body.Strategy != "aggressive" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		strategyMu.Lock()
		currentStrategy = body.Strategy
		strategyMu.Unlock()

		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":   "updated",
			"strategy": body.Strategy,
		})
	})

	// -------------------------------------------------
	// Start Server
	// -------------------------------------------------
	log.Fatal(http.ListenAndServe(":8080", nil))
}
