package structs

import "testing"

func checkArea(t testing.TB, shape Shape, wanted float64) {
	t.Helper()

	got := shape.Area()

	if got != wanted {
		t.Errorf("Expected %.2f got %.2f", wanted, got)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{10.0, 11.0}
		checkArea(t, r, 110.0)
	})

	t.Run("cirlce", func(t *testing.T) {
		c := Circle{10.0}
		checkArea(t, c, 314.1592653589793)
	})
}

func TestMultiArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 12, Height: 6}, 72.0},
		{Circle{Radius: 10}, 314.1592653589793},
		{Triangle{Length: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v Expected %.2f got %.2f", tt.shape, tt.want, got)
		}
	}
}
