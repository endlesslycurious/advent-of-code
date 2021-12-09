package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func ReadInput(filename string) []Digit {
	data := make([]Digit, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		digit := DigitFromLine(line)

		data = append(data, digit)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Loaded ", len(data), " digits from ", filename)

	return data
}

func main() {
	input := ReadInput("./2021/08/input.txt")

	answer := Part1(input)
	fmt.Println("Part 1 Answer:", answer)

	answer = Part2(input)
	fmt.Println("Part 2 Answer:", answer)
}

type Digit struct {
	input  []string // input values
	output []string // output values
}

func SortString(in string) string {
	raw := []rune(in)

	sort.Slice(raw, func(i, j int) bool { return raw[i] < raw[j] })

	return string(raw)
}

func ProcessStrings(line string) []string {
	output := make([]string, 0)

	for _, raw := range strings.Split(line, " ") {
		output = append(output, SortString(raw))
	}

	return output
}

func DigitFromLine(line string) Digit {
	sections := strings.Split(line, "|")

	if len(sections) != 2 {
		panic("Should only be two sections per line")
	}

	input := ProcessStrings(strings.TrimSpace(sections[0]))
	output := ProcessStrings(strings.TrimSpace(sections[1]))

	return Digit{input, output}
}

const (
	OneSegments   = 2 // Character 1 requires 2 segments
	FourSegments  = 4 // Character 4 requires 4 segments
	SevenSegments = 3 // Character 7 requires 3 segments
	EightSegments = 7 // Character 8 requires 7 segments
)

func Part1(digits []Digit) int {
	freq := make(map[int]int, 9)

	for _, digit := range digits {
		for _, output := range digit.output {
			segments := len(output)
			freq[segments]++
		}
	}

	count := freq[OneSegments]
	count += freq[FourSegments]
	count += freq[SevenSegments]
	count += freq[EightSegments]

	return count
}

func Diff(a, b string) int {
	var diff int

	if len(a) < len(b) {
		a, b = b, a
	}

	for _, chr := range a {
		if !strings.Contains(b, string(chr)) {
			diff++
		}
	}

	return diff
}

func FindDiff(pattern string, delta int, candidates []string) string {
	var found string
	for idx, signal := range candidates {
		if signal == "" {
			continue
		}

		if Diff(signal, pattern) == delta {
			candidates[idx] = ""
			found = signal
			break
		}
	}

	if len(found) == 0 {
		log.Fatalln("Signal with delta", delta, "from", pattern, "not found in", candidates)
	}

	return found
}

func Decode(input []string) map[string]int {
	lookup := make(map[string]int, len(input))
	remaining := make(map[int][]string, len(input)-4)

	var four, seven, eight string

	// identify 1, 4, 7 & 8 by length first
	for _, signal := range input {
		length := len(signal)
		switch length {
		case OneSegments:
			lookup[signal] = 1
		case FourSegments:
			lookup[signal] = 4
			four = signal
		case SevenSegments:
			lookup[signal] = 7
			seven = signal
		case EightSegments:
			lookup[signal] = 8
			eight = signal
		default:
			remaining[length] = append(remaining[length], signal)
		}
	}

	three := FindDiff(seven, 2, remaining[5])
	lookup[three] = 3

	five := FindDiff(four, 2, remaining[5])
	lookup[five] = 5

	two := FindDiff(four, 3, remaining[5])
	lookup[two] = 2

	six := FindDiff(four, 3, remaining[6])
	lookup[six] = 6

	nine := FindDiff(four, 2, remaining[6])
	lookup[nine] = 9

	zero := FindDiff(eight, 1, remaining[6])
	lookup[zero] = 0

	return lookup
}

func Part2(digits []Digit) int {
	var sum int

	for _, digit := range digits {
		// decode input into a lookup table
		lookup := Decode(digit.input)

		// process output using lookup
		output := lookup[digit.output[0]] * 1000
		output += lookup[digit.output[1]] * 100
		output += lookup[digit.output[2]] * 10
		output += lookup[digit.output[3]] * 1

		sum += output
	}

	return sum
}
