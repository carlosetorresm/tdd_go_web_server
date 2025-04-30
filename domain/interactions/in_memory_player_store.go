package interations

import (
	"sync"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
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

func (i *InMemoryPlayerStore) GetLeague() (lPlayers []league.Player) {
	for name, wins := range i.store {
		newPlayer := league.Player{Name: name, Wins: wins}
		lPlayers = append(lPlayers, newPlayer)
	}
	return
}
