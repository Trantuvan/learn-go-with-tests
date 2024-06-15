package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("word"))
}

func Hello(name string) string {
	if len(strings.TrimSpace(name)) == 0 {
		return "Hello, World"
	}
	return fmt.Sprint(englishHelloPrefix + name)
}
