package main

import "testing"

var input = [][]int{
	{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
	{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
	{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
	{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
	{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
}

func TestGetHeight(t *testing.T) {
	grid := CreateGrid(input)
	inputs := [][2]int{{0, 0}, {grid.width - 1, grid.height - 1}, {1, 4}, {3, 1}}
	expected := []int{2, 8, 8, 7}

	for i, point := range inputs {
		actual := grid.GetHeight(point[0], point[1])

		if actual != expected[i] {
			t.Error("Get", point, "expected", expected, "got", actual)
		}
	}
}

func TestGetNeighboursHeights(t *testing.T) {
	grid := CreateGrid(input)
	inputs := [][2]int{{0, 0}, {grid.width - 1, grid.height - 1}, {1, 4}, {3, 1}, {9, 1}}
	outputs := [][]int{{1, 3}, {7, 9}, {9, 9, 7}, {8, 8, 6, 9}, {2, 2, 0}}

	for i, point := range inputs {
		expected := outputs[i]
		actual := grid.GetNeighboursHeights(point[0], point[1])

		if len(actual) != len(expected) {
			t.Error("Get", point, "expected length", len(expected), "got", len(actual))
		}

		for idx := range expected {
			if actual[idx] != expected[idx] {
				t.Error("Get", point, "expected", expected, "got", actual)
				break
			}
		}
	}
}

func TestPart1(t *testing.T) {
	expected := 15

	actual := Part1(input)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
