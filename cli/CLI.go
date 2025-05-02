package cli

import (
	"io"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
)

type PlayerStore interface {
	GetPlayersScore(name string) int
	RecordWin(name string)
	GetLeague() league.League
}

type CLI struct {
	PlayerStore PlayerStore
	In          io.Reader
}

func (cli *CLI) PlayPoker() {
	cli.PlayerStore.RecordWin("Chris")
}
