package main

import (
	"os"

	"github.com/trantuvan/learn-go-with-tests/mockings"
)

func main() {
	// total := integers.Add(2, 2)
	// fmt.Printf("total %d", total)
	mockings.Countdown(os.Stdout)
}
