package main

import (
	"os"
	"time"

	"github.com/trantuvan/learn-go-with-tests/mockings"
)

// type defaultSleeper struct{}

// func (s defaultSleeper) Sleep() {
// 	time.Sleep(time.Second)
// }

func main() {
	// total := integers.Add(2, 2)
	// fmt.Printf("total %d", total)
	sleeper := mockings.ConfigurableSleeper{Duration: time.Second, Sleepfunc: time.Sleep}
	mockings.Countdown(os.Stdout, sleeper)
}
