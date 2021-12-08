package main

import (
	"testing"
)

func TestFindMax(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	expected := 16

	actual := FindMax(input)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestFindMaxPanics(t *testing.T) {
	input := []int{16, -1, 2, 0, 4, 2, 7, 1, 2, 14}

	defer func() {
		if r := recover(); r == nil {
			t.Error("FindMax did not panic")
		}
	}()

	FindMax(input)
}

var input = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func TestPart1(t *testing.T) {
	expected := 37

	actual := Part1(input)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 168

	actual := Part2(input)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
