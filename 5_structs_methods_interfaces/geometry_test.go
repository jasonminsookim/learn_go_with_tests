package geometry

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{5.0, 10.0}

		got := rectangle.Perimeter()
		want := 30.0

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}

		got := circle.Perimeter()
		want := 62.83185307179586

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}

func BenchmarkPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		(Rectangle{10, 20}).Perimeter()
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Width: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

func BenchmarkArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rectangle{Width: 5.0, Height: 10.0}.Area()
	}
}
