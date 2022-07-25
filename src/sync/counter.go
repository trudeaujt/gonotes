package main

import "sync"

type Counter struct {
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
