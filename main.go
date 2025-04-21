package main

import (
	"log"
	"net/http"

	interactions "github.com/carlosetorresm/tdd_go_web_server/domain/interactions"
	"github.com/carlosetorresm/tdd_go_web_server/server"
)

func main() {
	store := interactions.NewInMemoryPlayerStore()
	server := &server.PlayerServer{Store: store}
	log.Fatal(http.ListenAndServe(":5000", server))
}
