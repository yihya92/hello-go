package main

import (
	"context"
	"fmt"
	"regexp"
	"sync"
	"time"
)

type Job struct {
	ID    int
	Email string
}

type Result struct {
	JobID int
	Email string
	Valid bool
	Error error
}

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

// Worker with cancellation support
func worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			// Context cancelled
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

			// Simulate processing time
			time.Sleep(5 * time.Millisecond)

			valid := emailRegex.MatchString(job.Email)

			select {
			case results <- Result{
				JobID: job.ID,
				Email: job.Email,
				Valid: valid,
				Error: nil,
			}:
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	const totalJobs = 10000
	const workerCount = 10

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	jobs := make(chan Job)
	results := make(chan Result)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, results, &wg)
	}

	// Producer
	go func() {
		defer close(jobs)
		for i := 1; i <= totalJobs; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- Job{
				ID:    i,
				Email: fmt.Sprintf("user%d@example.com", i),
			}:
			}
		}
	}()

	// Close results after workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	processed := 0
	for range results {
		processed++
	}

	fmt.Println("Processed:", processed)

	if ctx.Err() != nil {
		fmt.Println("Stopped due to:", ctx.Err())
	}
}
