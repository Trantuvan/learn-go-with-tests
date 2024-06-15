package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("word"))
}

func Hello(name string) string {
	return fmt.Sprint(englishHelloPrefix + name)
}
