package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 5

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1) // Increment the counter for each worker
		go worker(w, jobs, &wg)
	}

	// Send jobs to the workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close the channel to signal that no more jobs will be sent

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All workers have completed their jobs")
}