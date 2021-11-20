package mock

import (
	"bytes"
	"testing"
)

func TestMock(t *testing.T) {
	buffer := &bytes.Buffer{}

	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()

	want := `3
2
1
Go!
`

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("Expected sleep{} to be called %d times", 4)
	}
}
