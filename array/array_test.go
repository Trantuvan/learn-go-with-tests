package array

import "testing"

func TestSum(t *testing.T) {
	numbers := [6]int{1, 2, 3, 4, 5, 10}
	got := Sum(numbers[:]...)
	want := 25

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
