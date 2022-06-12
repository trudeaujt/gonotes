package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const japaneseHelloPrefix = "ご機嫌よう、"

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

//This is a named return value! A variable called prefix will be created in the function.
//The initial value will depend on the type. For an int, 0; for a string, "".
func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	case "Japanese":
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	//We don't have to explicitly say `return prefix` since we've defined a named return value in the method declaration.
	//return prefix
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
