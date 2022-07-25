package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})
	t.Run("it runs safe concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		//A waitgroup is a convenient way of synchronizing concurrent processes.
		//It will wait for a collection of goroutines to finish.
		var wg sync.WaitGroup
		//First we pass in the number of executions we are looking for.
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				//Then after each execution we let the waitgroup know.
				wg.Done()
			}()
		}
		//It can also be used to block until all goroutines are finished.
		wg.Wait()

		assertCount(t, counter, wantedCount)
	})
}

//If we run go vet on our code, we get an error because we are copying our mutex after the first use.
//Without this, when we pass our Counter (by value) to assertCount it will try and create a copy of the mutex.
func NewCounter() *Counter {
	return &Counter{}
}

//To solve the above issue, we just pass in a pointer to our Counter instead.
//This has the side effect of showing readers that it would be better not to initialize the type themselves.
func assertCount(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
