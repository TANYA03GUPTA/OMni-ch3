package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

// processTask simulates processing a task (e.g., an API call)
func processTask(task Task) {
	fmt.Printf("Processing task %d\n", task.ID)
	// Simulate random processing time
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	fmt.Printf("Finished task %d\n", task.ID)
}

func worker(id int, tasks <- chan int, results chan<- int, wg *sync.WaitGroup){
	defer wg.Done()
	for task := range tasks{
		fmt.Printf("Worker %d: Received task %d\n", id, task)
		time.Sleep(2*time.Second)
		//processtask(task)
		fmt.Printf("Worker %d finished task %d\n",id,task)
		results <- task
	}
}

func main(){
	const numwrk =3
	const numtask = 5

	tasks := make(chan int, numtask)
	results := make(chan int, numtask)

	var wg sync.WaitGroup

	for i := 1;i<= numwrk;i++{
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}
//send task to channel 
	for i := 1;i <= numtask;i++{
		tasks <- i
	}
	close(tasks)
//wait for all workers to finish
	wg.Wait()
	close(results)

	fmt.Println("All tasks completed. Results:")
	for result := range results {
		fmt.Println(result)
	}
	

}