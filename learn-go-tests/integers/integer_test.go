package integer

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 4)
	expected := 6
	if sum != expected {
		t.Errorf("expected '%d' for '%d'", expected, sum)
	}
}
