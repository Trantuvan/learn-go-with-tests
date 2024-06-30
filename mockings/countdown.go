package mockings

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	if out == nil {
		return
	}
	for i := 3; i >= 1; i-- {
		fmt.Fprintln(out, i)
		if i == 1 {
			fmt.Fprint(out, "Go!")
		}
	}
}
