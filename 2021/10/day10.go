package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func ReadInput(filename string) []string {
	data := make([]string, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		data = append(data, line)
	}

	return data
}

func main() {
	filename := "./2021/10/day10_input.txt"
	input := ReadInput(filename)
	fmt.Println("Loaded", len(input), "lines from", filename)

	answer := Part1(input)
	fmt.Println("Part 1 Answer:", answer)

	answer = Part2(input)
	fmt.Println("Part 2 Answer:", answer)
}

// scores of incorect closing brackets
var scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

// Map closing to opening brackets
var closeToOpen = map[rune]rune{
	')': '(', // 41 : 40
	']': '[', // 93 : 91
	'}': '{', // 125 : 121
	'>': '<', // 62 : 60
}

// Map open to closing brackets
var openToClose = map[rune]rune{
	'(': ')', // 40 : 41
	'[': ']', // 91 : 93
	'{': '}', // 121 : 125
	'<': '>', // 60 : 62
}

// check line returns true and problem bracket if error otherwise false and zero
func CheckLine(line string) (bool, rune) {
	open := make([]rune, 0, len(line))

	for _, chr := range line {
		if chr == ')' || chr == ']' || chr == '}' || chr == '>' {
			prev := open[len(open)-1]

			if prev != closeToOpen[chr] {
				return true, chr
			}

			open = open[:len(open)-1]
		} else {
			open = append(open, chr)
		}

	}

	return false, 0
}

// Generate map of errors by bracket to occurence count
func FindErrors(input []string) map[rune]int {
	errors := make(map[rune]int)

	for _, line := range input {
		error, chr := CheckLine(line)

		if error {
			errors[chr]++
		}
	}

	return errors
}

func Part1(input []string) int {
	var score int
	errors := FindErrors(input)

	for bracket, count := range errors {
		subscore := count * scores[bracket]
		score += subscore
	}

	return score
}

// Filter out the incomplete lines
func FilterLines(input []string) []string {
	filtered := make([]string, 0, len(input)/2)

	for _, line := range input {
		err, _ := CheckLine(line)
		if !err {
			filtered = append(filtered, line)
		}
	}

	return filtered
}

// Generate the required pattern of closing braces to complete the line
func FixLine(line string) string {
	var fix string
	open := make([]rune, 0)

	// Work out unclosed openning brackets
	for _, chr := range line {
		if chr == ')' || chr == ']' || chr == '}' || chr == '>' {
			open = open[:len(open)-1]
		} else {
			open = append(open, chr)
		}
	}

	// Reverse order of closing brackets to fix line
	for i := len(open) - 1; i >= 0; i-- {
		chr := open[i]
		fix += string(openToClose[chr])
	}

	return fix
}

// score the fix
func ScoreFix(line string) int {
	var score int

	for _, chr := range line {
		score *= 5

		switch chr {
		case ')':
			score += 1
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}

	return score
}

func Part2(input []string) int {
	filtered := FilterLines(input)
	scores := make([]int, 0, len(filtered))

	for _, incomplete := range filtered {
		completed := FixLine(incomplete)

		score := ScoreFix(completed)

		scores = append(scores, score)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}
