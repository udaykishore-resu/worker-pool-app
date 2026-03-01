package workerpool

import "github.com/prometheus/client_golang/prometheus"

var (
	jobsSubmitted = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "workerpool_jobs_submitted_total",
			Help: "Total number of submitted jobs",
		},
	)

	jobsProcessing = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "workerpool_jobs_processing_total",
			Help: "Total number of jobs picked by workers",
		},
	)

	jobsCompleted = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "workerpool_jobs_completed_total",
			Help: "Total number of completed jobs",
		},
	)

	jobsFailed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "workerpool_jobs_failed_total",
			Help: "Total number of failed jobs",
		},
	)

	workersStopped = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "workerpool_workers_stopped_total",
			Help: "Total workers stopped due to cancellation",
		},
	)
)

func init() {
	prometheus.MustRegister(
		jobsSubmitted,
		jobsProcessing,
		jobsCompleted,
		jobsFailed,
		workersStopped,
	)
}

func IncrementJobsSubmitted()  { jobsSubmitted.Inc() }
func IncrementJobsProcessing() { jobsProcessing.Inc() }
func IncrementJobsCompleted()  { jobsCompleted.Inc() }
func IncrementJobsFailed()     { jobsFailed.Inc() }
func IncrementWorkersStopped() { workersStopped.Inc() }
