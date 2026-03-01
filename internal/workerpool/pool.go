package workerpool

import (
	"context"
	"sync"
)

type WorkerFunc[T any, R any] func(context.Context, T) (R, error)

type Pool[T any, R any] struct {
	numWorkers int
	jobs       chan T
	results    chan R
	workerFn   WorkerFunc[T, R]
	wg         sync.WaitGroup
}

func New[T any, R any](numWorkers, bufferSize int, workerFn WorkerFunc[T, R]) *Pool[T, R] {
	return &Pool[T, R]{
		numWorkers: numWorkers,
		jobs:       make(chan T, bufferSize),
		results:    make(chan R, bufferSize),
		workerFn:   workerFn,
	}
}

func (p *Pool[T, R]) Start(ctx context.Context) {
	for i := 1; i <= p.numWorkers; i++ {
		p.wg.Add(1)
		go p.worker(ctx, i)
	}

	go func() {
		p.wg.Wait()
		close(p.results)
	}()
}

func (p *Pool[T, R]) Submit(job T) {
	p.jobs <- job
	IncrementJobsSubmitted()
}

func (p *Pool[T, R]) CloseJobs() {
	close(p.jobs)
}

func (p *Pool[T, R]) Results() <-chan R {
	return p.results
}
