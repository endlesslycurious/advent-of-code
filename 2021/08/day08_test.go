package main

import (
	"testing"
)

var input = []Digit{
	DigitFromLine("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"),
	DigitFromLine("edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc"),
	DigitFromLine("fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg"),
	DigitFromLine("fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb"),
	DigitFromLine("aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea"),
	DigitFromLine("fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb"),
	DigitFromLine("dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe"),
	DigitFromLine("bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef"),
	DigitFromLine("egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb"),
	DigitFromLine("gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"),
}

func TestPart1(t *testing.T) {
	expected := 26

	actual := Part1(input)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestDecode(t *testing.T) {
	input := DigitFromLine("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	expected := map[string]int{
		SortString("acedgfb"): 8,
		SortString("cdfbe"):   5,
		SortString("gcdfa"):   2,
		SortString("fbcad"):   3,
		SortString("dab"):     7,
		SortString("cefabd"):  9,
		SortString("cdfgeb"):  6,
		SortString("eafb"):    4,
		SortString("cagedb"):  0,
		SortString("ab"):      1,
	}

	actual := Decode(input.input)

	if len(expected) != len(actual) {
		t.Error("Expected length", len(expected), "got", len(actual))
	}

	for key, val := range expected {
		ans, exists := actual[key]

		if !exists {
			t.Error("key", key, "not found")
			break
		}

		if ans != val {
			t.Error("Expected", val, "got", ans, "for key", key)
			break
		}
	}
}

func TestPart2ByDigit(t *testing.T) {
	expected := []int{8394, 9781, 1197, 9361, 4873, 8418, 4548, 1625, 8717, 4315}

	if len(expected) != len(input) {
		t.Error("Lengths of input", len(input), "and expected", len(expected), "differ")
	}

	for idx := range input {
		digit := []Digit{input[idx]}
		actual := Part2(digit)

		if actual != expected[idx] {
			t.Error("Expected", expected[idx], "got", actual, "for input", idx)
		}
	}
}

func TestPart2(t *testing.T) {
	expected := 61229

	actual := Part2(input)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
