# Go Worker Pool

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)

A production-grade, reusable, context-aware Worker Pool implementation in Go.

This project demonstrates controlled concurrency using the fan-out / fan-in pattern with proper lifecycle management, graceful shutdown, and goroutine leak prevention.

---

## ✨ Features

- Context-driven cancellation (`context.Context`)
- Graceful shutdown
- Fan-out / Fan-in concurrency pattern
- Safe channel lifecycle management
- Buffered job queue with backpressure support
- Reusable worker pool abstraction
- Clean modular project structure (`cmd/`, `internal/`)
- No goroutine leaks

---

## 🏗 Architecture
```bash
Producer → Jobs Channel → Worker Pool → Results Channel → Consumer
                          ↑
                     context cancellation
```
---
### Flow

1. Producer submits jobs.
2. Multiple workers process jobs in parallel (fan-out).
3. Results are aggregated into a single channel (fan-in).
4. Context controls lifecycle and cancellation.
5. Worker pool shuts down gracefully.

---

## 📂 Project Structure
```bash
worker-pool-app/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── workerpool/
│   │   ├── pool.go
│   │   └── worker.go
│   │
│   └── types/
│       └── job.go
│
├── pkg/
│   └── logger/
│       └── logger.go
│
├── go.mod
└── go.sum
```