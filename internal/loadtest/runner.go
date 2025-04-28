package loadtest

import (
	"load_testing/internal/config"
	"sync"
	"time"
)

type TestResults struct {
	TotalRequests      int64
	SuccessfulRequests int64
	FailedRequests     int64
	AverageResponseTime time.Duration
	RequestsPerSecond   float64
}

type Runner struct {
	config *config.Config
}

func NewRunner(cfg *config.Config) *Runner {
	return &Runner{
		config: cfg,
	}
}

func (r *Runner) Run() TestResults {
	resultsCh := make(chan Result, r.config.Workers)
	var wg sync.WaitGroup
	start := time.Now()

	// Start workers
	for i := 0; i < r.config.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker := NewWorker(r.config.URL, resultsCh)
			
			timer := time.NewTimer(r.config.Duration)
			for {
				select {
				case <-timer.C:
					return
				default:
					worker.Run()
				}
			}
		}()
	}

	// Create a goroutine to close results channel when all workers are done
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	// Collect and process results
	var results TestResults
	var totalResponseTime time.Duration

	for result := range resultsCh {
		results.TotalRequests++
		if result.Error == nil && result.StatusCode >= 200 && result.StatusCode < 400 {
			results.SuccessfulRequests++
		} else {
			results.FailedRequests++
		}
		totalResponseTime += result.ResponseTime
	}

	testDuration := time.Since(start)
	results.AverageResponseTime = totalResponseTime / time.Duration(results.TotalRequests)
	results.RequestsPerSecond = float64(results.TotalRequests) / testDuration.Seconds()

	return results
}