package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Siddharth-Mishra-23/linkedin-automation-lab/internal/state"
)

/* ---------------- GLOBAL STATE ---------------- */

// Strategy state
var (
	currentStrategy = "normal"
	strategyMu      sync.Mutex
)

// Stealth plugin state
var (
	stealthMu sync.Mutex
	stealthPlugins = map[string]bool{
		"mouse":       true,
		"timing":      true,
		"fingerprint": true,
		"scroll":      true,
		"typing":      true,
		"hover":       true,
		"schedule":    true,
		"ratelimit":   true,
	}
)

// Live logs
var (
	logMu   sync.Mutex
	logsBuf []string
)

/* ---------------- HELPERS ---------------- */

func addLog(msg string) {
	logMu.Lock()
	defer logMu.Unlock()

	if len(logsBuf) > 200 {
		logsBuf = logsBuf[len(logsBuf)-200:]
	}

	logsBuf = append(logsBuf, msg)
}

/* ---------------- MAIN ---------------- */

func main() {
	fmt.Println("🌐 LinkedIn Automation Lab — Web Dashboard")
	fmt.Println("Listening on http://localhost:8080")

	addLog("Server started")

	// -------------------------------------------------
	// Persistent State Tracker (STEP 7.3)
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
		addLog("Health check requested")
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	// -------------------------------------------------
	// Stats API
	// -------------------------------------------------
	http.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
		addLog("Stats requested")
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(stateTracker.Snapshot())
	})

	// -------------------------------------------------
	// Strategy APIs
	// -------------------------------------------------

	http.HandleFunc("/api/strategy", func(w http.ResponseWriter, r *http.Request) {
		addLog("Strategy fetched")
		w.Header().Set("Content-Type", "application/json")

		strategyMu.Lock()
		defer strategyMu.Unlock()

		_ = json.NewEncoder(w).Encode(map[string]string{
			"strategy": currentStrategy,
		})
	})

	http.HandleFunc("/api/strategy/set", func(w http.ResponseWriter, r *http.Request) {
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

		addLog("Strategy updated to " + body.Strategy)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":   "updated",
			"strategy": body.Strategy,
		})
	})

	// -------------------------------------------------
	// Stealth APIs
	// -------------------------------------------------

	http.HandleFunc("/api/stealth", func(w http.ResponseWriter, r *http.Request) {
		addLog("Stealth plugins fetched")
		w.Header().Set("Content-Type", "application/json")

		stealthMu.Lock()
		defer stealthMu.Unlock()

		_ = json.NewEncoder(w).Encode(stealthPlugins)
	})

	http.HandleFunc("/api/stealth/set", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var body struct {
			Plugin  string `json:"plugin"`
			Enabled bool   `json:"enabled"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		stealthMu.Lock()
		defer stealthMu.Unlock()

		if _, ok := stealthPlugins[body.Plugin]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		stealthPlugins[body.Plugin] = body.Enabled
		addLog("Stealth plugin toggled: " + body.Plugin)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"status":  "updated",
			"plugin":  body.Plugin,
			"enabled": body.Enabled,
		})
	})

	// -------------------------------------------------
	// Live Logs API
	// -------------------------------------------------

	http.HandleFunc("/api/logs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		logMu.Lock()
		defer logMu.Unlock()

		_ = json.NewEncoder(w).Encode(logsBuf)
	})

	// -------------------------------------------------
	// Export Logs API
	// -------------------------------------------------

	http.HandleFunc("/api/logs/export", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set(
			"Content-Disposition",
			"attachment; filename=automation_logs.txt",
		)

		logMu.Lock()
		defer logMu.Unlock()

		for _, line := range logsBuf {
			fmt.Fprintln(w, line)
		}
	})

	// -------------------------------------------------
	// Start Server
	// -------------------------------------------------

	log.Fatal(http.ListenAndServe(":8080", nil))
}
