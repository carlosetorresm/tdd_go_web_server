package filesystem

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

type FileSystemPlayerStore struct {
	Database *json.Encoder
	League   league.League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	lPlayers, err := league.NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}
	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&Tape{file}),
		League:   lPlayers,
	}, nil
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)
	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}
	return err
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

	f.Database.Encode(f.League)
}
