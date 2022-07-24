package _select

import (
	"fmt"
	"net/http"
	"time"
)

//These are not the best way to do this - we can do it concurrently using select!
func RacerWithoutSelect(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration > bDuration {
		return b
	}
	return a
}
func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

//Again, the two above is not the best way to do this!

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	//Recall that you can wait for values to be sent to a channel with the following blocking call:
	//myVar := <-ch
	//Select allows us to wait on multiple channels.
	//The first one to send a value `wins` and the code underneath the case is executed.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	//time.After is a very handy function when using select!
	//time.After returns a chan, and will send a signal down it after the amount of time we define.
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	//why use a struct{} channel and not a bool?
	//chan struct{} is the smallest data type available from a memory perspective.
	//So there is no allocation vs a bool.
	//Since we are just closing and not sending anything on the chan, why allocate anything?
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		//Send a signal into the channel once we have completed the request.
		close(ch)
	}()
	return ch
}
