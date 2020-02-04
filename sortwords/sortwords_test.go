package sortwords

import (
	"testing"
)




var expected = [5]string{
	"cows",
	"cows",
	"fight",
	"only",
	"only",
}

func TestInsertionSort(t *testing.T) {
	var input = [5]string{
		"only",
		"cows",
		"fight",
		"cows",
		"only",
	}
	InsertionSort(&input)
	if input != expected {
		t.Errorf("Array not sorted: %v", input)
	}
}


func TestSwap(t *testing.T) {
	a := [5]string{"one", "two", "three", "four", "five"}
	Swap(&a, 2, 4)
	expected := [5]string{"one", "two", "five", "four", "three"}
	if a != expected {
		t.Errorf("Problem with Swap: %v", a)
	}
}
