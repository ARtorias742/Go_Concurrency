package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(500 * time.Millisecond) // Simulate processing time
		result := task * 2
		results <- result
	}
	wg.Done()
}

func main() {
	const numWorkers = 3
	const numTasks = 20
	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, results, &wg)
	}
	// Send tasks
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks)
	// Wait for workers to finish and close results channel
	go func() {
		wg.Wait()
		close(results)
	}()
	// Collect results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
