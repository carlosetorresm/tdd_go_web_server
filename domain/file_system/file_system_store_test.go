package filesystem_test

import (
	"io"
	"os"
	"reflect"
	"testing"

	filesystem "github.com/carlosetorresm/tdd_go_web_server/domain/file_system"
	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

func TestFileSystemStore(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[
	{"Name": "Cleo", "Wins":10},
	{"Name": "Chris", "Wins":33}]`)
	defer cleanDatabase()

	store := filesystem.FileSystemPlayerStore{database}

	t.Run("league form reader", func(t *testing.T) {
		got := store.GetLeague()

		want := []league.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}
		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("Get player score", func(t *testing.T) {
		got := store.GetPlayersScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		existingPlayer := "Chris"
		store.RecordWin(existingPlayer)

		got := store.GetPlayersScore(existingPlayer)
		want := 34
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		newPlayer := "Pepper"
		store.RecordWin(newPlayer)

		got := store.GetPlayersScore(newPlayer)
		want := 1
		assertScoreEquals(t, got, want)
	})

}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
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

func assertScoreEquals(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertLeague(t testing.TB, got, want []league.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
