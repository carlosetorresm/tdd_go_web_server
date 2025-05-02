package filesystem

import (
	"encoding/json"
	"io"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker
	League   league.League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, io.SeekStart)
	lPlayers, _ := league.NewLeague(database)
	return &FileSystemPlayerStore{
		Database: database,
		League:   lPlayers,
	}
}

func (f *FileSystemPlayerStore) GetLeague() league.League {
	return f.League
}

func (f *FileSystemPlayerStore) GetPlayersScore(name string) int {
	player := f.League.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.League.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.League = append(f.League, league.Player{Name: name, Wins: 1})
	}

	f.Database.Seek(0, io.SeekStart)
	json.NewEncoder(f.Database).Encode(f.League)
}
