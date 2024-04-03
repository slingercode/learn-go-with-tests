package structsmethodsinterfaces

import "testing"

func assertShapes(t testing.TB, shape Shape, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("%#v: got %g want %g", shape, got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := rectangle.Perimeter()
	want := 40.0

	assertShapes(t, rectangle, got, want)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		assertShapes(t, test.shape, got, test.want)
	}
}
