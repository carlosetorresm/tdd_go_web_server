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
	in          *bufio.Scanner
}

func NewCLI(playerStore PlayerStore,
	in io.Reader) *CLI {
	return &CLI{playerStore: playerStore, in: bufio.NewScanner(in)}
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
