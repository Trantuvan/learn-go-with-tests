package main

import (
	"fmt"

	integers "github.com/trantuvan/learn-go-with-tests/integers"
)

func main() {
	total := integers.Add(2, 2)
	fmt.Printf("total %d", total)
}
