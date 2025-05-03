package cli_test

import (
	"strings"
	"testing"

	"github.com/carlosetorresm/tdd_go_web_server/cli"
	test "github.com/carlosetorresm/tdd_go_web_server/testing"
)

func TestCli(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &test.StubPlayerStore{}

		cli := cli.NewCLI(store, in)
		cli.PlayPoker()

		test.AssertPlayerWin(t, store, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &test.StubPlayerStore{}

		cli := cli.NewCLI(store, in)
		cli.PlayPoker()

		test.AssertPlayerWin(t, store, "Cleo")
	})
}
