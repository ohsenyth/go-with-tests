package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// commented out because of table driven test
// func TestArea(t *testing.T) {

// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		checkArea(t, rectangle, 72.0)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

func TestArea(t *testing.T) {

	areaTests := []struct {
		// shape Shape
		// want  float64

		name    string
		shape   Shape
		hasArea float64
	}{
		// {Rectangle{12, 6}, 72.1},
		// {Circle{10}, 314.1592653589793},
		// {Triangle{12, 6}, 36.0},

		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	// for _, tt := range areaTests {
	// 	got := tt.shape.Area()
	// 	if got != tt.want {
	// 		t.Errorf("got %.2f want %.2f", got, tt.want)
	// 	}
	// }

	// By wrapping each case in a t.Run you will have clearer test output on failures as it will print the name of the case
	// go test -run TestArea/Rectangle
	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})

	}

}
