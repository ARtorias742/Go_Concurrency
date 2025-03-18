package main

import (
	"fmt"
	"time"
)

func main() {

	// Create an unbuffered channel for strings
	ch := make(chan string)

	// Launch a goroutines to send a message
	go func() {
		time.Sleep(time.Second) // Simulate some work
		ch <- "Hello World!"    // Senf the message to the channel
	}()

	// Main goroutine waits to receive the message
	fmt.Println("Waiting for message...")
	msg := <-ch // Receive blocks until the message is sent
	fmt.Println("Received:", msg)
}
