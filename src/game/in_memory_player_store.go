package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.Mutex{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() (league []Player) {
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return
}
