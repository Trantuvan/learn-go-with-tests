package mockings

import (
	"fmt"
	"io"
	"time"
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
		time.Sleep(time.Second)
		if i == 1 {
			fmt.Fprint(out, finalWord)
		}
	}
}
