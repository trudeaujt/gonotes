package main

import (
	"github.com/trudeaujt/poker"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5001", server); err != nil {
		log.Fatalf("could not listen on port 5001 %v", err)
	}
}

//func main() {
//	//Handler is the interface that we need to implement in order to make a server.
//	//Usually, we would creat a struct and make it implement the interface by implementing its own serveHTTP method.
//	//But in this case, we don't have state yet, so it doesn't make sense to use a struct.
//	//That is where HandlerFunc comes in.
//	//It's an adapter to allow the use of normal functions as HTTP handlers without needing to create a struct.
//	//Looking at the documentation, HandlerFunc already implements the ServeHTTP method.
//	server := NewPlayerServer(NewInMemoryPlayerStore())
//	log.Fatal(http.ListenAndServe(":5001", server))
//}
