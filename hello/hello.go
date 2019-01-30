package main

import "fmt"

const helloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == spanish {
		return spanishHelloPrefix + name
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Serj", ""))
}
