package main

import (
	"fmt"
	"sync"
)

func main() {
	// Shared variable
	var current int
	iterations := 10

	// Mutex to ensure atomic operations on shared data
	var mu sync.Mutex
	wg := new(sync.WaitGroup)
	current = 8

	for i := 0; i < iterations; i++ {
		wg.Add(1)

		go func() {
			// Lock the mutex before modifying shared data
			mu.Lock()

			// Safe to modify the shared 'current' variable
			current++
			fmt.Println(current)

			// Unlock the mutex after modification is done
			mu.Unlock()

			wg.Done()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
