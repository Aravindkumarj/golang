package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const tamilHelloPrefix = "Vanakkam, "
const spanish = "Spanish"
const french = "French"
const tamil = "Tamil"

func Hello(name, language string) string {
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
	case tamil:
		prefix = tamilHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("Aravind", "Tamil"))
}
