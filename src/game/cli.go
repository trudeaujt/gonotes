package poker

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	out         io.Writer
	alerter     BlindAlerter
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
		out:         out,
		alerter:     alerter,
	}
}

func (c CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	numberOfPlayers, _ := strconv.Atoi(c.readLine())
	c.scheduleBlindAlerts(numberOfPlayers)

	userInput := c.readLine()
	c.playerStore.RecordWin(extractWinner(userInput))
}

func (c CLI) scheduleBlindAlerts(numberOfPlayers int) {
	fmt.Fprintf(os.Stdout, "players: %v\n", numberOfPlayers)
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		c.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (c CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Split(userInput, " ")[0]
}
