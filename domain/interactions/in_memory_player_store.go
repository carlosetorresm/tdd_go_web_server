package interations

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerScore {
	return &InMemoryPlayerScore{sync.Mutex{}, map[string]int{}}
}

type InMemoryPlayerScore struct {
	mu    sync.Mutex
	store map[string]int
}

func (i *InMemoryPlayerScore) GetPlayersScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerScore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}
