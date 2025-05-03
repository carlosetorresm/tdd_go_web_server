package main

import (
	"fmt"
	"log"
	"net/http"

	filesystem "github.com/carlosetorresm/tdd_go_web_server/domain/file_system"
	"github.com/carlosetorresm/tdd_go_web_server/server"
)

const dbFileName = "game.db.json"
const port = 5000

func main() {
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := server.NewPlayerServer(store)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), server); err != nil {
		log.Fatalf("could not listen on port %d %v", port, err)
	}
}
