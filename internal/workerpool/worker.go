package workerpool

import (
	"context"
)

func (p *Pool[T, R]) worker(ctx context.Context, id int) {
	defer p.wg.Done()

	for {
		select {
		case <-ctx.Done():
			IncrementWorkersStopped()
			return

		case job, ok := <-p.jobs:
			if !ok {
				return
			}

			IncrementJobsProcessing()

			result, err := p.workerFn(ctx, job)
			if err != nil {
				IncrementJobsFailed()
				continue
			}

			select {
			case p.results <- result:
				IncrementJobsCompleted()
			case <-ctx.Done():
				return
			}
		}
	}
}
