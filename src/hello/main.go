package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	prefix := englishHelloPrefix
	if name == "" {
		name = "World"
	}
	switch lang {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
