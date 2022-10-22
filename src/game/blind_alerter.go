package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// Any type can implement an interface, not just structs!
// Here we are implementing our interface with a function.
// With this, users of our interface will be able to implement it with just a single
// function, instead of having to create an empty struct type.
type BlindAlerterFunc func(duration time.Duration, amount int)

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

// We then create a function with the same signature as BlindAlerterFunc to print to stdout.
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
