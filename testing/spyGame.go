package testing

import "io"

type GameSpy struct {
	StartedWith  int
	StartCalled  bool
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int, to io.Writer) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}
