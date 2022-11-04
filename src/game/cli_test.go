package poker_test

import (
	"bytes"
	"github.com/trudeaujt/poker"
	"strings"
	"testing"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("start a game with 3 players and record chris wins from player input", func(t *testing.T) {
		in := userSends("3", "Chris wins")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		poker.AssertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		poker.AssertGameStartedWith(t, game, 3)
		poker.AssertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, poker.DummyStdOut, game)

		cli.PlayPoker()

		poker.AssertGameStartedWith(t, game, 8)
		poker.AssertFinishCalledWith(t, game, "Cleo")
	})
	t.Run("it prints an error when a non-numeric value is entered ", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := strings.NewReader("a\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		poker.AssertGameStartedWith(t, game, 0)
		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func userSends(strs ...string) *strings.Reader {
	input := ""
	for _, s := range strs {
		input = input + s + "\n"
	}
	return strings.NewReader(input)
}
