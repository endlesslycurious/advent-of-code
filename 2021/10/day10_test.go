package main

import "testing"

var input = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

func TestFindErrors(t *testing.T) {
	expected := map[rune]int{
		')': 2,
		']': 1,
		'}': 1,
		'>': 1,
	}

	actual := FindErrors(input)

	for bracket, count := range actual {
		expected_count, exists := expected[bracket]

		if !exists {
			t.Error(bracket, "was unexpected")
		}

		if count != expected_count {
			t.Error(bracket, "expected", expected_count, "got", count)
		}
	}
}

func TestPart1(t *testing.T) {
	expected := 26397
	actual := Part1(input)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 288957
	actual := Part2(input)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
