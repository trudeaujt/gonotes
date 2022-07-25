package main

import "sync"

type Counter struct {
	//Why are we not using channels here?
	//As stated on https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/sync#when-to-use-locks-over-channels-and-goroutines
	// - Use channels when passing ownership of data.
	// - Use mutexes for managing state.
	//The data is always owned by the counter... so no need for a channel.
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	//Any goroutine calling Inc will acquire the lock on Counter if they are first.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += 1
}

func (c *Counter) Value() int {
	return c.value
}
