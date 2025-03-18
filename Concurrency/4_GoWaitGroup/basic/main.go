package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this worker is done
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter for each worker
		go worker(i, &wg)
	}
	wg.Wait() // Block until all workers call Done()
	fmt.Println("All workers completed")
}
