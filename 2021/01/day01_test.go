package main

import (
	"log"
	"testing"
)

var input = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

// Test solution to first problem
func TestPart1(t *testing.T) {
	result := 7
	actual := Part1(input)
	if actual != result {
		log.Fatalln("FAIL! expected 7 got ", actual)
	}
}

// Test solution to second problem
func TestPart2(t *testing.T) {
	result := 5
	actual := Part2(input)
	if actual != result {
		log.Fatalln("FAIL! expected 5 got ", actual)
	}
}

var benchInput = ReadIntsFile("./input.txt")

// Benchmark first problem solution
func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1(benchInput)
	}
}

// Benchmark second problem solution
func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part2(benchInput)
	}
}
