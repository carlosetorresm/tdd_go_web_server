package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerScore interface {
	GetPlayersScore(name string) int
}

type PlayerServer struct {
	Store PlayerScore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.Store.GetPlayersScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}
