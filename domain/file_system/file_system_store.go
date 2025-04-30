package filesystem

import (
	"encoding/json"
	"io"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

type FileSystemPlayerStore struct {
	Database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []league.Player {
	var lPlayers []league.Player
	json.NewDecoder(f.Database).Decode(&lPlayers)
	return lPlayers
}
