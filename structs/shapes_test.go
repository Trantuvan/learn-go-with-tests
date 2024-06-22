package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		want := 40.0
		checkPerimeter(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 62.83185307179586
		checkPerimeter(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		want := 72.0
		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.159265358979
		checkArea(t, circle, want)
	})
}

func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func checkPerimeter(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Perimeter()

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
