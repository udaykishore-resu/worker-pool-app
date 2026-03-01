package workerpool

import (
	"context"
	"testing"
)

func BenchmarkWorkerPool(b *testing.B) {
	workerFn := func(ctx context.Context, job int) (int, error) {
		return job * 2, nil
	}

	pool := New[int, int](10, 1000, workerFn)
	ctx := context.Background()
	pool.Start(ctx)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.Submit(i)
	}

	pool.CloseJobs()
	for range pool.Results() {
	}
}
