package filesystem

import (
	"encoding/json"
	"io"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() league.League {
	f.Database.Seek(0, io.SeekStart)
	lPlayers, _ := league.NewLeague(f.Database)
	return lPlayers
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	}

	f.Database.Seek(0, io.SeekStart)
	json.NewEncoder(f.Database).Encode(league)
}
