package loadtest

import (
	"net/http"
	"time"
)

type Result struct {
	StatusCode   int
	ResponseTime time.Duration
	Error        error
	Timestamp    time.Time
}

type Worker struct {
	client    *http.Client
	url       string
	resultsCh chan<- Result
}

func NewWorker(url string, resultsCh chan<- Result) *Worker {
	return &Worker{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		url:       url,
		resultsCh: resultsCh,
	}
}

func (w *Worker) Run() {
	start := time.Now()
	resp, err := w.client.Get(w.url)
	duration := time.Since(start)

	result := Result{
		ResponseTime: duration,
		Timestamp:    time.Now(),
	}

	if err != nil {
		result.Error = err
	} else {
		result.StatusCode = resp.StatusCode
		resp.Body.Close()
	}

	w.resultsCh <- result
}
