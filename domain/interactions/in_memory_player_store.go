package interations

import (
	"sync"

	"github.com/carlosetorresm/tdd_go_web_server/server"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{sync.Mutex{}, map[string]int{}}
}

type InMemoryPlayerStore struct {
	mu    sync.Mutex
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayersScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []server.Player {
	//TODO implement
	return nil
}
