package main

import (
	"log"
	"net/http"

	"github.com/carlosetorresm/tdd_go_web_server/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
