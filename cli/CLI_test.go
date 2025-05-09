package cli_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/carlosetorresm/tdd_go_web_server/cli"
	test "github.com/carlosetorresm/tdd_go_web_server/testing"
)

var dummyStdOut = &bytes.Buffer{}

func userSends(messages ...string) *strings.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func TestCli(t *testing.T) {

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &test.GameSpy{}
		stdout := &bytes.Buffer{}

		player := "Chris"
		in := userSends("3", "Chris wins")
		cli.NewCLI(in, stdout, game).PlayPoker()

		assertMessagesSentToUser(t, stdout, cli.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, player)
	})

	t.Run("record 'Cleo' win from user input", func(t *testing.T) {
		game := &test.GameSpy{}

		player := "Cleo"
		in := userSends("1", "Cleo wins")
		cli.NewCLI(in, dummyStdOut, game).PlayPoker()

		assertGameStartedWith(t, game, 1)
		assertFinishCalledWith(t, game, player)
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &test.GameSpy{}
		stdOut := &bytes.Buffer{}
		in := userSends("notNumber")

		cli.NewCLI(in, stdOut, game).PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdOut, cli.PlayerPrompt, cli.BadPlayerInputErrMsg)

	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &test.GameSpy{}
		stdOut := &bytes.Buffer{}
		in := userSends("1", "Cleo is a cereal-Killer")
		cli.NewCLI(in, stdOut, game).PlayPoker()

		assertGameStartedWith(t, game, 1)
		assertMessagesSentToUser(t, stdOut, cli.PlayerPrompt, cli.BadWinnerInputMsg)
	})

}

func assertGameNotStarted(t *testing.T, game *test.GameSpy) {
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertGameStartedWith(t *testing.T, game *test.GameSpy, numberOfPlayers int) {
	t.Helper()
	if game.StartedWith != numberOfPlayers {
		t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
	}
}

func assertFinishCalledWith(t *testing.T, game *test.GameSpy, winner string) {
	t.Helper()
	if game.FinishedWith != winner {
		t.Errorf("expected finish called with %q but got %q", winner, game.FinishedWith)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertScheduledAlert(t testing.TB, got, want test.ScheduledAlert) {
	t.Helper()
	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}
