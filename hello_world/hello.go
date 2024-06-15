package hello_world

import (
	"fmt"
	"strings"
)

const (
	spanish            = "spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name, lan string) string {
	if len(strings.TrimSpace(name)) == 0 {
		return "Hello, World"
	}

	switch strings.ToLower(lan) {
	case spanish:
		return fmt.Sprint(spanishHelloPrefix, name)
	default:
		return fmt.Sprint(englishHelloPrefix, name)
	}
}
