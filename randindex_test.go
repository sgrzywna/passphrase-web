package main

import "testing"

func TestRandIndex(t *testing.T) {
	const maxNumber = 10
	rnd := NewRandIndex()
	res := rnd.RandInt32(maxNumber)
	if res < 0 || res > maxNumber {
		t.Errorf("expected value in range<0, %d>, got %d", maxNumber, res)
	}
}
