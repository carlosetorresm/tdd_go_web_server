package inmemoryserver

func NewInMemoryPlayerStore() *InMemoryPlayerScore {
	return &InMemoryPlayerScore{map[string]int{}}
}

type InMemoryPlayerScore struct {
	Store map[string]int
}

func (i *InMemoryPlayerScore) GetPlayersScore(name string) int {
	return i.Store[name]
}

func (i *InMemoryPlayerScore) RecordWin(name string) {
	i.Store[name]++
}
