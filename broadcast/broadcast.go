// Package broadcast  package that provides an abstraction to broadcast values to multiple goroutines
package broadcast

import (
	"sync"
)

type Broadcaster struct {
	channels   []chan interface{}
	register   chan chan interface{}
	unregister chan chan interface{}
	closeOnce  sync.Once
	mu         sync.RWMutex // Mutex to protect access to channels
}

func NewBroadcaster() *Broadcaster {
	b := &Broadcaster{
		channels:   make([]chan interface{}, 0),
		register:   make(chan chan interface{}),
		unregister: make(chan chan interface{}),
	}
	go b.run()
	return b
}

func (b *Broadcaster) Register() chan interface{} {
	ch := make(chan interface{})
	b.register <- ch
	return ch
}

func (b *Broadcaster) Unregister(ch chan interface{}) {
	b.unregister <- ch
}

func (b *Broadcaster) Broadcast(value interface{}) {
	b.mu.RLock()
	broadcastChannels := make([]chan interface{}, len(b.channels))
	copy(broadcastChannels, b.channels)
	b.mu.RUnlock()

	for _, ch := range broadcastChannels {
		select {
		case ch <- value:
		default:
		}
	}
}

func (b *Broadcaster) run() {
	for {
		select {
		case ch := <-b.register:
			b.mu.Lock()
			b.channels = append(b.channels, ch)
			b.mu.Unlock()
		case ch := <-b.unregister:
			b.mu.Lock()
			for i, c := range b.channels {
				if c == ch {
					b.channels = append(b.channels[:i], b.channels[i+1:]...)
					close(ch)
					break
				}
			}
			b.mu.Unlock()
		}
	}
}
