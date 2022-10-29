package poker_test

import (
	"bytes"
	"fmt"
	"github.com/trudeaujt/poker"
	"strings"
	"testing"
	"time"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("record chris win from player input", func(t *testing.T) {
		in := strings.NewReader("Chris Wins\n")
		playerStore := &poker.StubPlayerStore{}
		stdout := &bytes.Buffer{}
		cli := poker.NewCLI(playerStore, in, stdout, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo Wins\n")
		playerStore := &poker.StubPlayerStore{}
		stdout := &bytes.Buffer{}
		cli := poker.NewCLI(playerStore, in, stdout, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &poker.SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, poker.DummyStdOut, blindAlerter)
		cli.PlayPoker()

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n") // seven players
		blindAlerter := &poker.SpyBlindAlerter{}

		cli := poker.NewCLI(poker.DummyPlayerStore, in, stdout, blindAlerter)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
}

func assertScheduledAlert(t *testing.T, got poker.ScheduledAlert, want poker.ScheduledAlert) {
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
