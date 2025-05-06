package cli_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/carlosetorresm/tdd_go_web_server/cli"
	test "github.com/carlosetorresm/tdd_go_web_server/testing"
)

var dummystore = &test.StubPlayerStore{}
var dummyblindAlerter = &test.SpyBlindAlerter{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCli(t *testing.T) {
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &test.SpyBlindAlerter{}
		game := cli.NewGame(blindAlerter, dummystore)

		cli.NewCLI(in, stdout, game).PlayPoker()

		got := stdout.String()
		want := cli.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		cases := []test.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		checkSchedulingCases(t, cases, blindAlerter)

	})

	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &test.SpyBlindAlerter{}
		game := cli.NewGame(blindAlerter, dummystore)

		game.Start(5)
		cases := []test.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}
		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &test.SpyBlindAlerter{}
		game := cli.NewGame(blindAlerter, dummystore)

		game.Start(7)

		cases := []test.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

}

func TestGame_Finish(t *testing.T) {
	store := &test.StubPlayerStore{}
	game := cli.NewGame(dummyblindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	test.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(t *testing.T, cases []test.ScheduledAlert, blindAlerter *test.SpyBlindAlerter) {
	t.Helper()
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			assertScheduledAlert(t, got, want)
		})
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
