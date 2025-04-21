package main

import (
	"log"
	"net/http"

	"github.com/carlosetorresm/tdd_go_web_server/server"
)

type InMemoryPlayerScore struct{}

func (i *InMemoryPlayerScore) GetPlayersScore(name string) int {
	return 123
}

func main() {
	server := &server.PlayerServer{Store: &InMemoryPlayerScore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
