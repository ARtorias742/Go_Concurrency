package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	for i := 1; i <=5; i++ {
		fmt.Println("Goroutine:", i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go printNumbers(&wg)
	for i := 1; i <= 5; i++ {
		fmt.Println("Main:", i)
		time.Sleep(100 * time.Millisecond)

	}

	wg.Wait()
}