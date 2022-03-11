package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefixs(language) + name
}

func greetingPrefixs(language string) string {
	prefix := englishHelloPrefix
	switch language {
	case french:
		return frenchHelloPrefix
	case spanish:
		return spanishHelloPrefix
	default:
		return prefix
	}

}

func main() {
	fmt.Println(Hello("me", "am"))
}
