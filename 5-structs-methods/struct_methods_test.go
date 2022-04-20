package main

import "testing"

func TestPerimeter(t *testing.T) {
	retangle := Retangle{10.0, 10.0}
	got := Perimeter(retangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Retangle", Retangle{12, 6}, 72},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36},
	}

	for _, at := range areaTests {
		t.Run(at.name, func(t *testing.T) {
			got := at.shape.Area()
			if got != at.want {
				t.Errorf("%#v -> got: %.2f, want: %.2f", at.shape, got, at.want)
			}
		})
	}

}
