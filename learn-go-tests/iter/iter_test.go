package iter

import (
	"strings"
	"testing"
)

func TestIteration(t *testing.T) {
	repeatCount := 8
	repeated := Repeat("a", repeatCount)
	char := "a"
	expected := strings.Repeat(char, repeatCount)

	if repeated != expected {
		t.Errorf("Expected %q got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
