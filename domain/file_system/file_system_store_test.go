package filesystem_test

import (
	"testing"

	filesystem "github.com/carlosetorresm/tdd_go_web_server/domain/file_system"
	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
	test "github.com/carlosetorresm/tdd_go_web_server/testing"
)

func TestFileSystemStore(t *testing.T) {
	database, cleanDatabase := test.CreateTempFile(t, `[
	{"Name": "Cleo", "Wins":10},
	{"Name": "Chris", "Wins":33}]`)
	defer cleanDatabase()

	store, err := filesystem.NewFileSystemPlayerStore(database)
	test.AssertNoError(t, err)

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := test.CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := filesystem.NewFileSystemPlayerStore(database)
		test.AssertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		got := store.GetLeague()

		want := league.League{
			*league.NewPlayer("Chris", 33),
			*league.NewPlayer("Cleo", 10),
		}
		test.AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		test.AssertLeague(t, got, want)
	})

	t.Run("Get player score", func(t *testing.T) {
		got := store.GetPlayersScore("Chris")
		want := 33
		test.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		existingPlayer := "Chris"
		store.RecordWin(existingPlayer)

		got := store.GetPlayersScore(existingPlayer)
		want := 34
		test.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		newPlayer := "Pepper"
		store.RecordWin(newPlayer)

		got := store.GetPlayersScore(newPlayer)
		want := 1
		test.AssertScoreEquals(t, got, want)
	})

}
