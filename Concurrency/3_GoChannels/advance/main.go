package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	var wg sync.WaitGroup

	// Start three producer goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1) // Increment WaitGroup for each producer
		go func(id int) {
			defer wg.Done() // Signal completion when done
			for j := 0; j < 5; j++ {
				// Send values like 0-4 for id=0, 10-14 for id=1, 20-24 for id = 2
				ch <- id*10 + j
			}
		}(i)
	}

	// Launch a goroutine to close the channel after all producers finish
	go func() {
		wg.Wait()
		close(ch) // Close the channel to signal no more data
	}()

	// Consumer : Recieve all message from the channel
	for msg := range ch {
		fmt.Println(msg)
	}
}
