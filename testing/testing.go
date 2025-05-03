package testing

import (
	"os"
	"reflect"
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

func AssertScoreEquals(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func AssertLeague(t testing.TB, got, want league.League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpFile.Write([]byte(initialData))
	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}
	return tmpFile, removeFile
}

func AssertStatus(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("did not get correct status,got status %d want %d", got, want)
	}
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong,got %q, want %q", got, want)
	}
}
