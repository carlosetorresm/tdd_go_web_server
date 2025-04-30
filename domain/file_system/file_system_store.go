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
