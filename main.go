package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	filesystem "github.com/carlosetorresm/tdd_go_web_server/domain/file_system"
	"github.com/carlosetorresm/tdd_go_web_server/server"
)

const dbFileName = "game.db.json"
const port = 5000

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := filesystem.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}
	server := server.NewPlayerServer(store)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), server); err != nil {
		log.Fatalf("could not liten on port %d %v", port, err)
	}
}
