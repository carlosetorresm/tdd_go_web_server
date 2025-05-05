package cli_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/carlosetorresm/tdd_go_web_server/cli"
	test "github.com/carlosetorresm/tdd_go_web_server/testing"
)

func TestCli(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &test.StubPlayerStore{}
		var dummySpyAlerter = &test.SpyBlindAlerter{}

		cli := cli.NewCLI(store, in, dummySpyAlerter)
		cli.PlayPoker()

		test.AssertPlayerWin(t, store, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &test.StubPlayerStore{}
		var dummySpyAlerter = &test.SpyBlindAlerter{}

		cli := cli.NewCLI(store, in, dummySpyAlerter)
		cli.PlayPoker()

		test.AssertPlayerWin(t, store, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &test.StubPlayerStore{}
		blindAlerter := &test.SpyBlindAlerter{}

		cli := cli.NewCLI(store, in, blindAlerter)
		cli.PlayPoker()

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

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= 1 {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}
				got := blindAlerter.Alerts[i]
				assertScheduledAlert(t, got, want)

			})
		}
	})
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
