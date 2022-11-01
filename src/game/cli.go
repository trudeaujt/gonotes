package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

/*
CLI is just concerned with:
- constructing TexasHoldem with its existing dependencies, and
- Interpreting user input as method invocations for TexasHoldem
*/
type CLI struct {
	in  *bufio.Scanner
	out io.Writer
	/*
		From a domain perspective, we want to:
		  - Start a TexasHoldem, indicating how many people are playing
		  - Finish a TexasHoldem, declaring the winner
		The new TexasHoldem type encapsulates this for us.

		With this, we have passed BlindAlerter and PlayerStore to TexasHoldem since it is now responsible for alerting as well as storing results.
	*/
	/*
		Extracting this out into an interface means that we can use mocks to test that the Game functionality is being properly called!
	*/
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number."

func (c CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)

	numberOfPlayersInput := c.readLine()
	numberOfPlayers, err := strconv.Atoi(numberOfPlayersInput)

	if err != nil {
		fmt.Fprint(c.out, BadPlayerInputErrMsg)
		return
	}

	c.game.Start(numberOfPlayers)

	winnerInput := c.readLine()
	winner := extractWinner(winnerInput)
	c.game.Finish(winner)
}

func (c CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Split(userInput, " ")[0]
}
