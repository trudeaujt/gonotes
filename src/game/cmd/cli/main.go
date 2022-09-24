package main

import (
	"fmt"
	"github.com/trudeaujt/poker"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	fmt.Println("Let's play poker!")
	println("type {name} wins to record a win.")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
