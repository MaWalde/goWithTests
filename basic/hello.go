package main

import (
	"fmt"
)

<<<<<<< HEAD
func Hello(name string) string {
	return "Hello, " + name
=======
const englishHelloPrefix = "Hello, "

const (
	french  = "French"
	spanish = "Spanish"
	swedish = "Swedish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	swedishHelloPrefix = "Hej, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case swedish:
		prefix = swedishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
>>>>>>> e590d57 (Refactor, lifting out constants)
}

func main() {
	fmt.Println(Hello("world"))
}
