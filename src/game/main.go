package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	//Handler is the interface that we need to implement in order to make a server.
	//Usually, we would creat a struct and make it implement the interface by implementing its own serveHTTP method.
	//But in this case, we don't have state yet, so it doesn't make sense to use a struct.
	//That is where HandlerFunc comes in.
	//It's an adapter to allow the use of normal functions as HTTP handlers without needing to create a struct.
	//Looking at the documentation, HandlerFunc already implements the ServeHTTP method.
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))
}
