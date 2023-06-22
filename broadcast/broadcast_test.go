package broadcast

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBroadcast(t *testing.T) {
	b := NewBroadcaster()

	// Create three goroutines that register to receive values
	receiverCount := 3
	var wg sync.WaitGroup
	wg.Add(receiverCount)

	for i := 0; i < receiverCount; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			ch := b.Register()
			value := <-ch
			if value != "Hello, world!" {
				t.Errorf("Received unexpected value: %v", value)
			}
		}(&wg)
	}

	time.Sleep(time.Second) // Wait for goroutines to register

	b.Broadcast("Hello, world!")

	wg.Wait()
}

func TestUnregister(t *testing.T) {
	b := NewBroadcaster()

	ch1 := b.Register()
	ch2 := b.Register()

	b.Unregister(ch1)

	go b.Broadcast("Hello, world!")

	select {
	case msg, ok := <-ch1:
		if ok {
			t.Error("Unregistered channel should not receive a value", msg)
		}
	case msg, ok := <-ch2:
		if !ok {
			t.Error("Second registered channel should receive a value, but is closed", msg)
		}
		fmt.Println("Received value from second channel", msg)
	}

}

func TestConcurrentBroadcast(t *testing.T) {
	b := NewBroadcaster()

	// Create three goroutines that register to receive values
	receiverCount := 3
	var wg sync.WaitGroup
	wg.Add(receiverCount)

	for i := 0; i < receiverCount; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			ch := b.Register()
			value := <-ch
			if value != "Hello, world!" {
				t.Errorf("Received unexpected value: %v", value)
			}
		}(&wg)
	}

	time.Sleep(time.Second) // Wait for goroutines to register

	// Perform concurrent broadcasts
	broadcastCount := 10
	for i := 0; i < broadcastCount; i++ {
		go b.Broadcast("Hello, world!")
	}

	wg.Wait()
}
