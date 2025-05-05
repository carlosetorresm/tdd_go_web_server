package cli

import (
	"bufio"
	"io"
	"strings"
	"time"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type PlayerStore interface {
	GetPlayersScore(name string) int
	RecordWin(name string)
	GetLeague() league.League
}

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	alerter     BlindAlerter
}

func NewCLI(playerStore PlayerStore,
	in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{playerStore: playerStore, in: bufio.NewScanner(in), alerter: alerter}
}

func (cli *CLI) PlayPoker() {
	cli.scheduleBlindAlerts()
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
