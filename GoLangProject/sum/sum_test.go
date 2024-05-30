package sum

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	expected := 15

	if result != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, expected)
	}
}
