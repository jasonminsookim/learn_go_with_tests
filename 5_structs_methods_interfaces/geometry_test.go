package geometry

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	t.Run("peimeter of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{5.0, 10.0}

		got := Perimeter(rectangle)
		want := 30.0

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}

func BenchmarkPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Perimeter(Rectangle{10,20})
	}
}

func TestArea(t *testing.T) {

	t.Run("area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{5.0, 10.0}

		got := Area(rectangle)
		want := 50.0

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}


func BenchmarkA(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		Area(Rectangle{Width:5.0, Height:10.0})
	}
}