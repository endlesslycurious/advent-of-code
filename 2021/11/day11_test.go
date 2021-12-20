package main

import "testing"

var input = [][]int{
	{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
	{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
	{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
	{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
	{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
	{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
	{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
	{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
	{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
	{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
}

func TestGetNeighbours(t *testing.T) {
	grid := CreateGrid(input)
	inputs := []Point{{0, 0}, {9, 9}, {5, 5}, {0, 4}, {3, 9}}
	expected := []int{3, 3, 8, 5, 5}

	for i, in := range inputs {
		neighbours := grid.GetNeighbours(in.x, in.y)
		actual := len(neighbours)

		if actual != expected[i] {
			t.Error("Expected", expected[i], "for", in, "got", actual)
		}
	}
}

func TestPart1Short(t *testing.T) {
	expected := 204
	actual := Part1(input, 10)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestPart1Long(t *testing.T) {
	expected := 1656
	actual := Part1(input, Steps)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
