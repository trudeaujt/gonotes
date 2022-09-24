package poker_test

import (
	"github.com/trudeaujt/poker"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from player input", func(t *testing.T) {
		in := strings.NewReader("Chris Wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo Wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
