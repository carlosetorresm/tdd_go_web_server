package league

import (
	"encoding/json"
	"fmt"
	"io"
)

type Player struct {
	Name string `json:"name"`
	Wins int    `json:"wins"`
}

func NewPlayer(name string, wins int) *Player {
	player := new(Player)
	player.Name = name
	player.Wins = wins
	return player
}

type League []Player

func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
