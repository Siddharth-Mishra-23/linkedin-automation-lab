# LinkedIn Automation Lab

> ⚠️ **Educational Purpose Only**
>
> This project is a proof-of-concept designed strictly for educational and technical evaluation purposes.
> It demonstrates browser automation, system design, and anti-detection strategies.
> **It must NOT be used on real LinkedIn accounts or production environments.**

---

## 🚀 Project Overview

LinkedIn Automation Lab is a Go-based browser automation proof-of-concept that simulates
human-like interaction patterns while demonstrating advanced system architecture,
stealth techniques, and observability.

---

## 🧠 Key Features

- Modular Go architecture
- Browser automation using Rod
- Human-like behavior simulation
- Anti-detection (stealth) techniques
- CLI + Web dashboard (dark UI)
- Strategy-based automation profiles
- Real-time analytics and logs

---

## 🧱 Architecture Overview

The system is composed of multiple independent engines:

- Authentication Engine
- Search & Targeting Engine
- Connection & Messaging Engine
- Stealth Engine (plugin-based)
- Strategy Engine (behavior modeling)
- Scheduler & Rate Limiter
- Analytics & Observability
- UI Layer (CLI + Web)

Detailed architecture is available in `/docs`.

---

## ⚙️ Setup Instructions

```bash
git clone <repo-url>
cd linkedin-automation-lab
cp .env.example .env
go mod tidy
go run ./cmd/cli
