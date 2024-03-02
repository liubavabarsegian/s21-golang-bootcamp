package anscombe

import "testing"

func TestMean(t *testing.T) {
	numsArr := []int{1, 2, 3}
	mean := Mean(numsArr)
	if mean != 2 {
		t.Fatal()
	}
}
