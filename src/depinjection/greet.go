package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//bytes.Buffer implements the io.Writer interface
//If we just accept bytes.Buffer as our writer type, our function is not very useful because it cannot print to stdout!
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//We can call our generalized function like this to print to stdout.
//func main() {
//	Greet(os.Stdout, "Elodie")
//}

//We can call our generalized function like this to return a web page!
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
