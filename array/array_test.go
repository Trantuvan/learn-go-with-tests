package array

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 6 numbers", func(t *testing.T) {
		numbers := [6]int{1, 2, 3, 4, 5, 10}
		got := Sum(numbers[:]...)
		want := 25

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers...)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	fmt.Printf("%v", want)

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
