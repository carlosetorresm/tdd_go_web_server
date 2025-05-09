package testing

import "io"

type GameSpy struct {
	StartedWith int
	StartCalled bool
	BlindAlert  []byte

	FinishedWith string
	FinishCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int, to io.Writer) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
	to.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalled = true
	g.FinishedWith = winner
}
