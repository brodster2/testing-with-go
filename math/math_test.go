package math

import "testing"

func TestSum(t *testing.T) {
	tSum := Sum([]int{5, 3, 2})
	if tSum != 10 {
		t.Errorf("Expected 10 but got %d", tSum)
	}
}
