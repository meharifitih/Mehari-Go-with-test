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

	prefix := englishHelloPrefix
	switch language {
	case french:
		return frenchHelloPrefix + name
	case spanish:
		return spanishHelloPrefix + name
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
