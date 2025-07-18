package main

import (
	"fmt"
)

const french = "French"
const spanish = "Spanish"
const swedish = "Swedish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const swedishHelloPrefix = "Hej, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case swedish:
		prefix = swedishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world, ", ""))
}
