package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Siddharth-Mishra-23/linkedin-automation-lab/internal/state"
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
	// Start Server
	// -------------------------------------------------
	log.Fatal(http.ListenAndServe(":8080", nil))
}
