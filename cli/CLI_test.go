package cli_test

import (
	"testing"

	"github.com/carlosetorresm/tdd_go_web_server/cli"
	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	lPlayers league.League
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

func TestCli(t *testing.T) {
	playerStore := &StubPlayerStore{}
	cli := &cli.CLI{playerStore}
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

}
