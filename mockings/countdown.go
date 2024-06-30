package mockings

import (
	"fmt"
	"io"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer) {
	if out == nil {
		return
	}
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
		if i == 1 {
			fmt.Fprint(out, finalWord)
		}
	}
}
