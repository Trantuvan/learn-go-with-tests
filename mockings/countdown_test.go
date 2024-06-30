package mockings

import (
	"bytes"
	"testing"
)

type spySleeper struct {
	calls int
}

func (s *spySleeper) Sleep() {
	s.calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &spySleeper{}
	AnotherCountDown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.calls != countdownStart {
		t.Errorf("not enough calls to sleeper, want %d got %d", countdownStart, spySleeper.calls)
	}
}
