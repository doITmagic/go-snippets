package broadcast

import (
	"fmt"
	"sync"
	"time"
)

func receive(ch <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	value := <-ch
	fmt.Println("Received:", value)
}

// ExampleBroadcast Create un example for the use of broadcast package
func ExampleBroadcast() {
	b := NewBroadcaster()

	var wg sync.WaitGroup

	wg.Add(3)
	go receive(b.Register(), &wg)
	go receive(b.Register(), &wg)
	go receive(b.Register(), &wg)

	// Wait for goroutines to register
	time.Sleep(time.Second)

	b.Broadcast("Hello, world!")

	wg.Wait()
	fmt.Println("All goroutines have received the broadcast.")

	// Output:
	// Received: Hello, world!
	// Received: Hello, world!
	// Received: Hello, world!
	// All goroutines have received the broadcast.
}
