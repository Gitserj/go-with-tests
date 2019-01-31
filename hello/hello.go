package main

import "fmt"

const helloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	prefix = helloPrefix
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Serj", ""))
}
