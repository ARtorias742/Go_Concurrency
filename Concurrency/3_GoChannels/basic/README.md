### How It Works:
Channel Creation: ch := make(chan string) creates an unbuffered channel for strings.
Sender Goroutine: A goroutine is started that sleeps for one second (to simulate work) and then sends "Hello, World!" to the channel using ch <-.
Receiver: The main goroutine prints "Waiting for message..." and then blocks on <-ch, waiting for the sender to provide data.
Synchronization: The receive operation waits until the send occurs, demonstrating how unbuffered channels synchronize goroutines.
This example illustrates the basic mechanics of channels: sending and receiving data, with built-in synchronization.

Understanding Unbuffered Channels
In the example above, we used an unbuffered channel. Here’s what that means:

Send Operation (ch <- value): Blocks the sender until a receiver is ready to take the value.
Receive Operation (<-ch): Blocks the receiver until a sender provides a value.
This blocking behavior ensures that the sender and receiver are synchronized. Unbuffered channels are perfect for scenarios where you need one goroutine to wait for another, as seen in the basic example.

Advanced Example: Fan-In with Multiple Producers
Now, let’s explore a more advanced use case: the fan-in pattern. In this pattern, multiple producer goroutines send data to a single channel, and one consumer goroutine collects all the data. We’ll use a sync.WaitGroup to coordinate the producers and close the channel when they’re done.

Scenario:
Producers: Three goroutines, each sending five integers to a shared channel.
Consumer: The main goroutine receives and prints all messages.
Synchronization: A WaitGroup ensures the channel is closed only after all producers finish.