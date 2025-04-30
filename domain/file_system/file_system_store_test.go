package filesystem_test

import (
	"reflect"
	"strings"
	"testing"

	filesystem "github.com/carlosetorresm/tdd_go_web_server/domain/file_system"
	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league form reader", func(t *testing.T) {
		database := strings.NewReader(`[
	{"Name": "Cleo", "Wins":10},
	{"Name": "Chris", "Wins":33}]`)

		store := filesystem.FileSystemPlayerStore{database}
		got := store.GetLeague()

		want := []league.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}
		assertLeague(t, got, want)
	})
}

func assertLeague(t testing.TB, got, want []league.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
