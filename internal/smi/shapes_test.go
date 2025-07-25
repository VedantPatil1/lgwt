package smi

import (
	"math"
	"testing"
)

func TestPerimiter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimiter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10.0}, (math.Pi * 100)},
		{Triangle{12.0, 6.0}, 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("Got %g Expected %g", got, tt.want)
		}
	}
}
