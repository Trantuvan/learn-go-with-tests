package mockings

import (
	"fmt"
	"io"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

func Countdown(out io.Writer, sleeper Sleeper) {
	if out == nil {
		return
	}
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
		if i == 1 {
			fmt.Fprint(out, finalWord)
		}
	}
}
