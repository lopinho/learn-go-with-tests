package main

import "fmt"

const SPANISH = "Spanish"
const FRENCH = "French"
const PORTUGUESE = "Portuguese"
const ENGLISH_HELLO_PREFIX = "Hello, "
const SPANISH_HELLO_PREFIX = "Hola, "
const FRENCH_HELLO_PREFIX = "Bonjour, "
const PORTUGUESE_HELLO_PREFIX = "Olá, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := greetingPrefix(language)

	return prefix + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case FRENCH:
		prefix = FRENCH_HELLO_PREFIX
	case SPANISH:
		prefix = SPANISH_HELLO_PREFIX
	case PORTUGUESE:
		prefix = PORTUGUESE_HELLO_PREFIX
	default:
		prefix = ENGLISH_HELLO_PREFIX
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
