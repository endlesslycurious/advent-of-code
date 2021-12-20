package main

import (
	"fmt"
	"testing"
)

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

		for _, neighbour := range neighbours {
			if in.x == neighbour.x && in.y == neighbour.y {
				t.Error("Original point in neighbours list")
			}
		}
	}
}

var afterStep1 = [][]int{
	{6, 5, 9, 4, 2, 5, 4, 3, 3, 4},
	{3, 8, 5, 6, 9, 6, 5, 8, 2, 2},
	{6, 3, 7, 5, 6, 6, 7, 2, 8, 4},
	{7, 2, 5, 2, 4, 4, 7, 2, 5, 7},
	{7, 4, 6, 8, 4, 9, 6, 5, 8, 9},
	{5, 2, 7, 8, 6, 3, 5, 7, 5, 6},
	{3, 2, 8, 7, 9, 5, 2, 8, 3, 2},
	{7, 9, 9, 3, 9, 9, 2, 2, 4, 5},
	{5, 9, 5, 7, 9, 5, 9, 6, 6, 5},
	{6, 3, 9, 4, 8, 6, 2, 6, 3, 7},
}

var afterStep2 = [][]int{
	{8, 8, 0, 7, 4, 7, 6, 5, 5, 5},
	{5, 0, 8, 9, 0, 8, 7, 0, 5, 4},
	{8, 5, 9, 7, 8, 8, 9, 6, 0, 8},
	{8, 4, 8, 5, 7, 6, 9, 6, 0, 0},
	{8, 7, 0, 0, 9, 0, 8, 8, 0, 0},
	{6, 6, 0, 0, 0, 8, 8, 9, 8, 9},
	{6, 8, 0, 0, 0, 0, 5, 9, 4, 3},
	{0, 0, 0, 0, 0, 0, 7, 4, 5, 6},
	{9, 0, 0, 0, 0, 0, 0, 8, 7, 6},
	{8, 7, 0, 0, 0, 0, 6, 8, 4, 8},
}

var afterStep3 = [][]int{
	{0, 0, 5, 0, 9, 0, 0, 8, 6, 6},
	{8, 5, 0, 0, 8, 0, 0, 5, 7, 5},
	{9, 9, 0, 0, 0, 0, 0, 0, 3, 9},
	{9, 7, 0, 0, 0, 0, 0, 0, 4, 1},
	{9, 9, 3, 5, 0, 8, 0, 0, 6, 3},
	{7, 7, 1, 2, 3, 0, 0, 0, 0, 0},
	{7, 9, 1, 1, 2, 5, 0, 0, 0, 9},
	{2, 2, 1, 1, 1, 3, 0, 0, 0, 0},
	{0, 4, 2, 1, 1, 2, 5, 0, 0, 0},
	{0, 0, 2, 1, 1, 1, 9, 0, 0, 0},
}

func CheckGrid(grid Grid, stepGrid Grid, step int, t *testing.T) bool {
	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			expected := stepGrid.GetEnergy(x, y)
			actual := grid.GetEnergy(x, y)

			if actual != expected {
				t.Error("Step", step, "at", x, y, "expected", expected, "got", actual)

				fmt.Println("    Expected Step", step, "State")
				stepGrid.Print()

				fmt.Println("    Actual Step", step, "State")
				grid.Print()
				return false
			}
		}
	}

	return true
}

func TestGridUpdateStep1(t *testing.T) {
	grid := CreateGrid(input)
	expected := []Grid{CreateGrid(afterStep1), CreateGrid(afterStep2), CreateGrid(afterStep3)}

	for i, stepGrid := range expected {
		grid.Update()
		ok := CheckGrid(grid, stepGrid, i+1, t)

		if !ok {
			break
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
