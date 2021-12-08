package main

import "testing"

func TestPart1(t *testing.T) {
	var input []Digit // todo fill in input
	expected := 26

	actual := Part1(input)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
