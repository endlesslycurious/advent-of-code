package main

import (
	"log"
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
		log.Fatalln("FAIL! expected ", result, " got ", actual)
	}
}

func TestPart2(t *testing.T) {
	result := 230
	actual := Part2(input)
	if actual != result {
		log.Fatalln("FAIL! expected ", result, " got ", actual)
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
