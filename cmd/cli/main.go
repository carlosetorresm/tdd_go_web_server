package main

import (
	"fmt"
	"log"
	"os"

	"github.com/carlosetorresm/tdd_go_web_server/cli"
	filesystem "github.com/carlosetorresm/tdd_go_web_server/domain/file_system"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := cli.NewGame(cli.BlindAlerterFunc(cli.StdOutAlerter), store)
	cli.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
