package testing

import (
	"testing"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	lPlayers league.League
}

func NewStubPlayerStore(scores map[string]int, winCalls []string,
	lPlayers league.League) *StubPlayerStore {
	return &StubPlayerStore{scores: scores, winCalls: winCalls, lPlayers: lPlayers}
}

func (s *StubPlayerStore) GetPlayersScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() league.League {
	return s.lPlayers
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.winCalls) != 1 {
		t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}
	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}
