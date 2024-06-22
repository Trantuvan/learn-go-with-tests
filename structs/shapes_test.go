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
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}

func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()

	if got != want {
		t.Errorf("Shape: %#v got %g want %g", shape, got, want)
	}
}

func checkPerimeter(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Perimeter()

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
