package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(id int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // signal that this producer is done
	for i := 0; i < 5; i++ {
		result := id*10 + i
		fmt.Printf("Producer %d produced %d\n", id, result)
		results <- result                  // send result to channel
		time.Sleep(time.Millisecond * 100) // simulate work
	}
}

func consumer(results <-chan int, done chan<- bool) {
	for result := range results {
		fmt.Printf("Consumed %d\n", result)
	}

	done <- true // Signal that the consumer is done
}

func main() {
	var wg sync.WaitGroup
	results := make(chan int) // Channel for passing data
	done := make(chan bool)   // Channel to signal consumer completion

	//Start two producer goroutines
	for i := 1; i <= 2; i++ {
		wg.Add(1) // Incement the counter for each producer
		go producer(i, results, &wg)
	}

	// Start the consumer goroutine
	go consumer(results, done)

	wg.Wait()
	close(results)

	<-done
	fmt.Println("All done")

}
