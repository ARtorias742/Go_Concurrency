// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var counter int

// func increment(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	counter++ // Race condition here
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go increment(&wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("Final counter value:", counter) // May not be 10
// }

// // go run -race main.go


package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter) // Always 10
}