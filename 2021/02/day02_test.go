package main

import (
	"testing"
)

// Test input from problem
var input = []Instruction{
	{Forward, 5},
	{Down, 5},
	{Forward, 8},
	{Up, 3},
	{Down, 8},
	{Forward, 2},
}

// Test part one solution
func TestPart1(t *testing.T) {
	result := 150
	actual := Part1(input)
	if actual != result {
		t.Error("expected ", result, " got ", actual)
	}
}

var benchInput = LoadInstructions("./day02_input.txt")

// Benchmark part one solution
func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1(benchInput)
	}
}

// Test part two solution
func TestPart2(t *testing.T) {
	result := 900
	actual := Part2(input)
	if actual != result {
		t.Error("expected ", result, " got ", actual)
	}
}

// Benchmark part two solution
func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part2(benchInput)
	}
}
