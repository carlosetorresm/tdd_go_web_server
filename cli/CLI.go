package cli

import (
	"bufio"
	"io"
	"strings"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

type PlayerStore interface {
	GetPlayersScore(name string) int
	RecordWin(name string)
	GetLeague() league.League
}

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func NewCLI(playerStore PlayerStore,
	in io.Reader) *CLI {
	return &CLI{playerStore: playerStore, in: in}
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.in)
	reader.Scan()
	cli.playerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
