package math

import "testing"

func TestSum(t *testing.T) {
	expected := 11
	sum := Sum([]int{10, -2, 3})
	if sum != expected {
		t.Errorf("Wanted %d but received %d", expected, sum)
	}
}

func TestAdd(t *testing.T) {
	expected := 5
	add := Add(3, 2)
	if add != expected {
		t.Errorf("Wanted %d but received %d", expected, add)
	}
}
