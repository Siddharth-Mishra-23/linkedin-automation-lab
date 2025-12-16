🚀 LinkedIn Automation Lab

## 🎥 Demo Video

A short walkthrough demonstrating setup, execution, and key features of the LinkedIn Automation Lab:

(https://drive.google.com/drive/folders/1j5UaaGjZFJgrzBSSiFqeG7lNVnhqIIDF?usp=sharing)



⚠️ Educational & Demonstration Purpose Only

This project is a technical proof-of-concept built to demonstrate system design, browser automation architecture, runtime configuration, and observability.

❌ It is NOT intended for real LinkedIn usage, production deployment, or violation of any platform’s terms of service.

📌 Overview

LinkedIn Automation Lab is a Go-based automation framework with a real-time web dashboard that simulates human-like interaction behavior.

The project focuses on:

clean backend architecture

modular behavior strategies

stealth & anti-detection concepts

persistent state tracking

live observability through logs & metrics

This is not a script — it is a system.

✨ Key Features
🧠 Strategy-Driven Behavior

Switch between automation profiles at runtime:

Slow & Careful

Normal User

Aggressive

Each strategy controls delays, rate limits, and action pacing.

🛡️ Stealth Plugin System

Modular stealth plugins that can be toggled live:

Mouse movement simulation

Typing behavior

Scroll randomness

Hover behavior

Timing jitter

Browser fingerprint masking

Schedule enforcement

Rate limiting

Plugins are runtime-configurable via the dashboard.

⏱️ Daily Rate Limiting & Persistence

Tracks daily actions

Enforces maximum limits

Automatically persists state

Prevents unsafe overuse

📊 Real-Time Web Dashboard (Dark UI)

Neon dark theme UI

Live system status

Strategy selector

Stealth plugin toggles

Daily usage progress bar

Live streaming logs

📜 Live Observability

In-memory log buffer

Real-time UI updates

Backend event visibility

Debug-friendly design

🧱 Architecture
linkedin-automation-lab/
├── cmd/
│   ├── cli/        # CLI-based automation runner
│   └── server/     # Web dashboard backend
├── internal/
│   ├── browser/    # Browser abstraction (mock/real)
│   ├── stealth/    # Stealth plugin system
│   ├── strategy/   # Behavior strategies
│   ├── scheduler/  # Time-based execution control
│   ├── state/      # Persistent daily state tracking
│   └── analytics/  # Metrics & snapshots
├── ui/
│   └── web/        # Dark themed dashboard (HTML/CSS/JS)
├── configs/
├── scripts/
└── README.md
⚙️ Tech Stack

Language: Go (Golang)

Backend: net/http

Frontend: HTML, CSS, Vanilla JS

State: File-backed persistence

Concurrency: Goroutines + Mutexes

UI Theme: Dark Neon / Glassmorphism

🚀 Getting Started
1️⃣ Clone the repository
git clone https://github.com/Siddharth-Mishra-23/linkedin-automation-lab.git
cd linkedin-automation-lab

2️⃣ Run the Web Dashboard
go run ./cmd/server

Expected output:
🌐 LinkedIn Automation Lab — Web Dashboard
Listening on http://localhost:8080

3️⃣ Open the Dashboard

If running locally:
http://localhost:8080

If running on GitHub Codespaces:

Open Ports

Click Open Browser on port 8080

🧪 Available APIs
| Endpoint            | Description          |
| ------------------- | -------------------- |
| `/api/health`       | System health        |
| `/api/stats`        | Daily usage stats    |
| `/api/strategy`     | Get current strategy |
| `/api/strategy/set` | Update strategy      |
| `/api/stealth`      | Get stealth plugins  |
| `/api/stealth/set`  | Toggle plugins       |
| `/api/logs`         | Live system logs     |

🧑‍💻 CLI Mode (Mock Automation)
go run ./cmd/cli

Runs a mock browser automation flow using:

strategy engine

stealth plugins

scheduler

state tracker

🎯 What This Project Demonstrates

This project was built to showcase:

Backend system design

Runtime configuration without restarts

Clean modular architecture

Thread-safe state handling

Observability & monitoring concepts

UI-driven control of backend behavior

This is interview-grade, not tutorial-grade.

📸 Demo

Live dashboard with real-time updates

Strategy switching

Plugin toggling

Log streaming

(Screenshots / demo video can be added here)

📜 Disclaimer

This project is strictly for:

learning

academic evaluation

portfolio demonstration

The author does not endorse real-world automation of LinkedIn or similar platforms.

👤 Author

Siddharth Mishra
B.Tech CSE
GitHub: @Siddharth-Mishra-23

⭐ Final Note

If you’re reviewing this project:

It was intentionally designed to reflect real backend systems,
not a one-off automation script.
