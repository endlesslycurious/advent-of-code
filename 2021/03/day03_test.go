package main

import (
	"testing"
)

var input = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestPart1(t *testing.T) {
	result := 198
	actual := Part1(input)
	if actual != result {
		t.Error("expected ", result, " got ", actual)
	}
}

func TestPart2(t *testing.T) {
	result := 230
	actual := Part2(input)
	if actual != result {
		t.Error("expected ", result, " got ", actual)
	}
}

func TestFilterByBit(t *testing.T) {
	expected := []string{
		"11110",
		"10110",
		"10111",
		"10101",
		"11100",
		"10000",
		"11001",
	}
	actual := FilterByBit(input, '1', 0)

	if len(actual) != len(expected) {
		t.Error("expected length ", len(expected), " got ", len(actual))
	}
}

func TestStrToBinary(t *testing.T) {
	inputs := []string{"00001", "10000", "01010"}
	expected := []int{1, 16, 10}

	for i, input := range inputs {
		actual := StrToBinary(input)

		if actual != expected[i] {
			t.Error("Expected ", expected[i], " got ", actual)
		}
	}
}

func TestCountOnes(t *testing.T) {
	counts := []float64{7, 5, 8, 7, 5}

	for i, expected := range counts {
		actual := CountOnes(input, i)

		if actual != expected {
			t.Error("Expected ", expected, " got ", actual)
		}
	}
}

var benchInput = ReadPowerBits("./input.txt")

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1(benchInput)
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part2(benchInput)
	}
}
