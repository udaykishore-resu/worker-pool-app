package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"worker-pool-app/internal/workerpool"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() *trace.TracerProvider {
	tp := trace.NewTracerProvider()
	otel.SetTracerProvider(tp)
	return tp
}

func main() {

	// Context cancellation via OS signal
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// Start Prometheus metrics endpoint
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	// Worker logic
	workerFn := func(ctx context.Context, job int) (int, error) {
		time.Sleep(500 * time.Millisecond)
		return job * 2, nil
	}

	pool := workerpool.New[int, int](3, 10, workerFn)
	pool.Start(ctx)

	// Producer
	go func() {
		for i := 1; i <= 20; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				pool.Submit(i)
			}
		}
		pool.CloseJobs()
	}()

	// Consumer
	for result := range pool.Results() {
		fmt.Println("Received:", result)
	}

	fmt.Println("Graceful shutdown complete")
}
