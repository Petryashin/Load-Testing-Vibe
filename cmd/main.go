package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"load_testing/internal/config"
	"load_testing/internal/loadtest"
)

func main() {
	// Parse command line flags
	url := flag.String("url", "", "Target URL to test")
	workers := flag.Int("workers", 10, "Number of concurrent workers")
	duration := flag.Duration("duration", 60*time.Second, "Test duration")
	flag.Parse()

	if *url == "" {
		log.Fatal("URL is required")
	}

	// Create configuration
	cfg := &config.Config{
		URL:      *url,
		Workers:  *workers,
		Duration: *duration,
	}

	// Create and run load test
	runner := loadtest.NewRunner(cfg)
	results := runner.Run()

	// Print results
	fmt.Printf("\nLoad Test Results:\n")
	fmt.Printf("Total Requests: %d\n", results.TotalRequests)
	fmt.Printf("Successful Requests: %d\n", results.SuccessfulRequests)
	fmt.Printf("Failed Requests: %d\n", results.FailedRequests)
	fmt.Printf("Average Response Time: %v\n", results.AverageResponseTime)
	fmt.Printf("Requests/Second: %.2f\n", results.RequestsPerSecond)
}