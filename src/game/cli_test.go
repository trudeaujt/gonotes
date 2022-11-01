package poker_test

import (
	"bytes"
	"github.com/trudeaujt/poker"
	"strings"
	"testing"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("start a game with 5 players and record chris wins from player input", func(t *testing.T) {
		in := strings.NewReader("5\nChris Wins\n")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		want := "Chris"
		if game.FinishedWith != want {
			t.Errorf("wanted %s, got %s", game.FinishedWith, want)
		}
	})
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		in := strings.NewReader("7\n")
		out := &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		got := out.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		if game.StartedWith != 7 {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("it prints an error when a non-numeric value is entered ", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := strings.NewReader("a\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}
		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}
