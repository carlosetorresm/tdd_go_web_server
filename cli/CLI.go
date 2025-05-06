package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type BlindAlerterFunc func(duration time.Duration, amount int)

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}

type PlayerStore interface {
	GetPlayersScore(name string) int
	RecordWin(name string)
	GetLeague() league.League
}

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game *Game
}

func NewCLI(in io.Reader, out io.Writer, game *Game) *CLI {
	return &CLI{in: bufio.NewScanner(in), out: out, game: game}
}

const PlayerPrompt = "Please enter the number of players: "

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayerInput := cli.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayerInput, "\n"))

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)
	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
