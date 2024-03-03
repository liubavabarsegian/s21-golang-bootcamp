package anscombe

import "testing"

func TestMean1(t *testing.T) {
	numsArr := []int{1, 2, 3}
	mean := Mean(numsArr)
	if mean != 2 {
		t.Fatal()
	}
}

func TestMean2(t *testing.T) {
	numsArr := []int{-100, 3, 4, 5}
	mean := Mean(numsArr)
	if mean != -22 {
		t.Fatal()
	}
}

func TestMedianOdd(t *testing.T) {
	numsArr := []int{1, 2, 3}
	median := Median(numsArr)
	if median != 2 {
		t.Fatal()
	}
}

func TestMedianEven(t *testing.T) {
	numsArr := []int{1, 2, 3, 4}
	median := Median(numsArr)
	if median != 2.5 {
		t.Fatal()
	}
}

func TestMode(t *testing.T) {
	numsArr := []int{1, 2, 3, 5, 5, 5, 6, 6, 6}
	mode := Mode(numsArr)
	if mode != 5 {
		t.Fatal()
	}
}

func TestMode2(t *testing.T) {
	numsArr := []int{1}
	mode := Mode(numsArr)
	if mode != 1 {
		t.Fatal()
	}
}
