package main

import (
	"fmt"
	"strings"
)

const spanish = "spanish"
const french = "french"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return findPrefix(lang) + name
}

func findPrefix(lang string) (prefix string) {
	l := strings.ToLower(lang)
	switch l {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
