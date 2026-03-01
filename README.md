# Go Worker Pool

![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)

A production-grade, reusable, generic, context-aware Worker Pool implementation in Go.

This project demonstrates controlled concurrency using the fan-out / fan-in pattern, complete lifecycle management, graceful shutdown via OS signals, Prometheus instrumentation, and goroutine leak prevention.
---

## ✨ Features

- Generic `WorkerPool[T, R]` (Go Generics)
- Context-driven cancellation (`context.Context`)
- OS signal–based graceful shutdown (`SIGINT, SIGTERM`)
- Fan-out / Fan-in concurrency pattern
- Safe channel lifecycle management
- Buffered job queue with backpressure support
- Prometheus metrics instrumentation
- Reusable worker function abstraction
- Clean modular project structure (`cmd/, internal/, pkg/`)
- No goroutine leaks
- Production-ready observability endpoint (`/metrics`)

---

## 🏗 Architecture
```bash
Producer → Jobs Channel → Worker Pool → Results Channel → Consumer
                          ↑
                 Context + OS Signal Cancellation
```
---
## 🏗 Distributed Architecture
```plain text
Producer Service
      ↓
   Kafka Topic
      ↓
Worker Pool Consumers (Multiple Pods)
      ↓
Result Topic / DB
```
### Flow

1. Application starts with signal.NotifyContext.
2. Prometheus metrics endpoint is exposed.
3. Producer submits jobs into buffered queue.
4. Multiple workers process jobs concurrently (fan-out).
5. Results are aggregated into a single channel (fan-in).
6. On:
    - Job completion
    - Context timeout
    - OS signal (Ctrl+C)
    → Workers terminate gracefully.
7. Results channel closes automatically after all workers exit.
8. Application shuts down cleanly with zero goroutine leaks.

---

## 📂 Project Structure
```bash
worker-pool-app/
│
├── cmd/
│   └── app/
│       └── main.go          # Application entrypoint + signal handling
│
├── internal/
│   ├── workerpool/
│   │   ├── pool.go          # Generic WorkerPool[T, R]
│   │   ├── worker.go        # Worker execution logic
│   │   └── metrics.go       # Prometheus instrumentation
│   │
│   └── types/
│       └── job.go           # Generic job definition
│
├── pkg/
│   └── logger/
│       └── logger.go        # Reusable logger package
│
├── go.mod
└── go.sum
```
---
## ⚙️ Internal Concurrency Flow
```code
main()
  ├── signal.NotifyContext()
  ├── Start()
  │     ├── worker 1
  │     ├── worker 2
  │     └── worker N
  │
  ├── Submit jobs
  ├── CloseJobs()
  └── range Results()
```
---
## 🧠 Generic Worker Pool Design
```go
type WorkerFunc[T any, R any] func(context.Context, T) (R, error)

type Pool[T any, R any] struct {
    numWorkers int
    jobs       chan T
    results    chan R
    workerFn   WorkerFunc[T, R]
}
```

This makes the pool reusable for:
- CPU-bound workloads
- I/O-bound tasks
- API calls
- File processing
- Database operations
- Streaming pipelines

## 📊 Prometheus Metrics
Exposed at:
```code
http://localhost:2112/metrics
```

**Available Metrics**
| Metric                             | Description                      |
| ---------------------------------- | -------------------------------- |
| `workerpool_jobs_submitted_total`  | Total submitted jobs             |
| `workerpool_jobs_processing_total` | Jobs picked by workers           |
| `workerpool_jobs_completed_total`  | Successfully completed jobs      |
| `workerpool_jobs_failed_total`     | Failed jobs                      |
| `workerpool_workers_stopped_total` | Workers stopped via cancellation |

## 🔐 Graceful Shutdown
The application listens for:
- SIGINT
- SIGTERM
When triggered:
- Context is cancelled
- Workers stop accepting new work
- In-flight jobs finish safely
- Results channel closes automatically
- Application exits cleanly

No dangling goroutines. No channel panics.