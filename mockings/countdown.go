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

/*
order of exec:
write - sleep - write until countdownStart - sleep - Go!
*/
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

/*
order of exec: sleep - sleep - sleep -> write - write - write -> Go!
*/
func AnotherCountDown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		sleeper.Sleep()
	}
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}
