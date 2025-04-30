package filesystem

import (
	"io"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

type FileSystemPlayerStore struct {
	Database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []league.Player {
	f.Database.Seek(0, io.SeekStart)
	lPlayers, _ := league.NewLeague(f.Database)
	return lPlayers
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}
