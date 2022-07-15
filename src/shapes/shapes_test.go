package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	//We are declaring an anonymous struct, areaTests.
	//We declare a slice of these structs by using []struct with two fields, the shape and the want.
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{ //then we fill the slice with the test cases.
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{
			name:  "Triangle",
			shape: Triangle{12, 6},
			want:  36.0,
		}, //we can optionally name the fields... but intellij does this for us automatically
	}

	//We can iterate over our test cases just like any other slice, using the struct fields to run our tests!
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			//We can use %#v to print out our struct with the values in its fields, so we can see what is being tested.
			t.Errorf("%#v got %g, want %g", tt.shape, got, tt.want)
		}
	}

}
