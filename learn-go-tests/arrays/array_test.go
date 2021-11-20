package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("Expected %d got %d given %v", expected, got, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{4, 3})
	expected := []int{3, 7}

	checkSum(t, got, expected)

}

func TestSumAllTails(t *testing.T) {
	t.Run("slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{4, 3})
		expected := []int{2, 3}

		checkSum(t, got, expected)
	})

	t.Run("empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 3})
		expected := []int{0, 3}
		checkSum(t, got, expected)

	})

}

func checkSum(t testing.TB, got, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %d got %d", expected, got)
	}
}
