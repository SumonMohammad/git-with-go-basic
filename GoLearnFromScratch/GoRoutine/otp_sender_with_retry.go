package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type OTPJob struct {
	User     string
	MaxRetry int
}

func otpSender(id int, jobs <-chan OTPJob, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		for attempt := 1; attempt <= job.MaxRetry+1; attempt++ {
			fmt.Printf("The %d worker sending the otp to %s (Attempt %d)\n", id, job.User, attempt)

			time.Sleep(500 * time.Millisecond)

			if rand.Intn(100) < 30 {
				if attempt <= job.MaxRetry {
					fmt.Printf("Worker %d failed %s, retrying (%d/%d)\n", id, job.User, attempt, job.MaxRetry)
					time.Sleep(1000 * time.Millisecond)
					continue
				} else {
					fmt.Printf("Worker %d failed %s, after retrying (%d/%d)\n", id, job.User, job.MaxRetry, job.MaxRetry)
					results <- fmt.Sprintf("❌ %s: Failed after retries", job.User)
					break
				}
			} else {
				fmt.Printf("Worker %d successfully sent the OTP to %s\n", id, job.User)
				results <- fmt.Sprintf("✅ %s: OTP sent successfully", job.User)
				break
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numOfSender := 3
	maxRetry := 2
	var wg sync.WaitGroup

	users := []string{"Alice", "Bob", "Carol", "Dave", "Eve", "Frank", "Grace", "Heidi", "Ivan", "Judy"}
	jobs := make(chan OTPJob, len(users))
	results := make(chan string, len(users))

	// Start workers
	for i := 1; i <= numOfSender; i++ {
		wg.Add(1)
		go otpSender(i, jobs, results, &wg)
	}

	// Send jobs
	for _, user := range users {
		jobs <- OTPJob{
			User:     user,
			MaxRetry: maxRetry,
		}
	}
	close(jobs)

	// Wait for workers to finish and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for res := range results {
		fmt.Println(res)
	}
}
