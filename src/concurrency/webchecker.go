package main

import (
	"fmt"
	"net/http"
)

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	//If we try and write our results directly to the map, there is a chance that a data race condition will be hit.
	//To get around this, we use something called `channels`.
	resultChannel := make(chan result)

	for _, url := range urls {
		//To call a function concurrently, just use `go func()`!
		go func(u string) {
			//doing like this results in a race condition when multiple threads try and write to the map at the same time!
			//results[u] = wc(u)

			//This `<-` is called a send statement.
			//We are sending a result struct for each call to wc to the resultChannel.
			//channel <- result
			resultChannel <- result{u, wc(u)}

			//We need to capture the current context within our anonymous function, or else all of the
			//calls will be using the last URL since our function is so fast...
		}(url) //Just like JS, we can call the anonymous function immediately like so.
	}

	//After we've filled in the channel async with structs of type `result`, we need to add those to our map synchronously.
	for i := 0; i < len(urls); i++ {
		//We use a receive expression to retrieve the values from the channel.
		//variable <- channel
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

func httpCheck(url string) bool {
	r, err := http.Get(url)
	//fmt.Printf("result for %s is %s\n", url, r)
	if err != nil {
		return false
	}
	if r.StatusCode == 200 {
		return true
	}
	return false
}

func main() {
	urls := []string{"http://google.com", "http://facebook.com", "http://go.dev", "http://twitter.com"}
	result := CheckWebsites(httpCheck, urls)
	fmt.Print(result)
}
