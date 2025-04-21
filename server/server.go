package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerScore interface {
	GetPlayersScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerScore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.Store.GetPlayersScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
