package utils

import (
	"mocks_rabbitmq/mocks"
	"sync"
)

const listenerSlots = 128

// ErrBroadcast enables broadcast an error channel to various listener channels
type ErrBroadcast struct {
	sync.Mutex // Protects listeners
	listeners  []chan<- mocks.Error
	c          chan mocks.Error
}

// NewErrBroadcast creates a broadcast object for push errors to subscribed channels
func NewErrBroadcast() *ErrBroadcast {
	b := &ErrBroadcast{
		c:         make(chan mocks.Error),
		listeners: make([]chan<- mocks.Error, 0, listenerSlots),
	}

	go func() {
		for {
			select {
			case e := <-b.c:
				b.spread(e)
			}
		}
	}()

	return b
}

// Add a new listener
func (b *ErrBroadcast) Add(c chan<- mocks.Error) {
	b.Lock()
	b.listeners = append(b.listeners, c)
	b.Unlock()
}

// Delete the listener
func (b *ErrBroadcast) Delete(c chan<- mocks.Error) {
	i, ok := b.findIndex(c)
	if !ok {
		return
	}
	b.Lock()
	b.listeners[i] = b.listeners[len(b.listeners)-1]
	b.listeners[len(b.listeners)-1] = nil
	b.listeners = b.listeners[:len(b.listeners)-1]
	b.Unlock()
}

// Write to subscribed channels
func (b *ErrBroadcast) Write(err mocks.Error) {
	b.c <- err
}

func (b *ErrBroadcast) spread(err mocks.Error) {
	b.Lock()
	for _, l := range b.listeners {
		l <- err
	}
	b.Unlock()
}

func (b *ErrBroadcast) findIndex(c chan<- mocks.Error) (int, bool) {
	b.Lock()
	defer b.Unlock()

	for i := range b.listeners {
		if b.listeners[i] == c {
			return i, true
		}
	}
	return -1, false
}
