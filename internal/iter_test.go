package internal

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	s := "hello"

	actual := Repeat(3, s)
	expected := []string{"hello", "hello", "hello"}

	if actual[0] != expected[0] && actual[1] != expected[1] && actual[2] != expected[2] {
		t.Fatal()
	}
}

func TestCycle(t *testing.T) {
	s := []string{"hello"}
	ch := Cycle(s)
	expected := "hello"

	for i := 0; i < 5; i++ {
		actual := <-ch
		if actual != expected {
			t.Fatal()
		}
	}
}
