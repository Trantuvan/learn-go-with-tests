package mockings

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type spySleeper struct {
	calls int
}

func (s *spySleeper) Sleep() {
	s.calls++
}

type spyCountdownOps struct {
	calls []string
}

func (s *spyCountdownOps) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *spyCountdownOps) Write([]byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &spySleeper{}
		Countdown(buffer, spySleeper)

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
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &spyCountdownOps{}
		Countdown(spySleeperPrinter, spySleeperPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(want, spySleeperPrinter.calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrinter.calls)
		}
	})
}

type spyTime struct {
	durationSlept time.Duration
}

func (s *spyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &spyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
