package smi

import (
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return (math.Pi * c.Radius * c.Radius)
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (0.5 * t.Base * t.Height)
}

type Shape interface {
	Area() float64
}
