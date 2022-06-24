package geometry

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Width float64
	Height float64
}


func (t Triangle) Area() float64 {
	return t.Height * t.Width / 2
}