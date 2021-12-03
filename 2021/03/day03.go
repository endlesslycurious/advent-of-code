package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadPowerBits(filename string) []string {
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

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Loaded ", len(data), " readings from ", filename)

	return data
}

func main() {
	input := ReadPowerBits("./2021/03/input.txt")

	answer := Part1(input)
	fmt.Println("Part 1 answer: ", answer)

	answer = Part2(input)
	fmt.Println("Part 2 answer: ", answer)
}

func Part1(input []string) int {
	oneFreq := make(map[int]int, 5)
	total := len(input)

	// Generate freqency map of 1s in input, storing count of matching lines
	for _, data := range input {
		for i, bit := range data {
			if bit == '1' {
				oneFreq[i]++
			}
		}
	}

	// bits indexed from zero not one
	var bits = len(oneFreq) - 1
	var gamma, epsilon int

	// Use 1s frequency map to work out most common value per column/bit
	// to assemble gamma & epsilon
	for i, count := range oneFreq {
		val := 1 << (bits - i)

		if count > total/2 {
			gamma = gamma | val
		} else {
			epsilon = epsilon | val
		}
	}

	return gamma * epsilon
}

// Return list of origs with filter set at specified index
func FilterByBit(origs []string, filter byte, index int) []string {
	filtered := make([]string, 0)

	for _, orig := range origs {
		if orig[index] == filter {
			filtered = append(filtered, orig)
		}
	}

	return filtered
}

// Convert binary string e.g. "01010" to integer
func StrToBinary(in string) int {
	var result int

	for i, val := range in {
		pow := len(in) - 1 - i
		if val == '1' {
			result = result | (1 << pow)
		}
	}

	return result
}

// Count all the inputs with 1 at specified index
func CountOnes(input []string, index int) float64 {
	var count float64

	for _, data := range input {
		if data[index] == '1' {
			count++
		}
	}

	return count
}

// Filter by One decision logic type
type filterByOneDecision func(x, y float64) bool

// Gamma log filtering decision logic
func filterGamma(count, half float64) bool {
	return count >= half
}

// Epsilon log filtering decision logic
func filterEpsilon(count, half float64) bool {
	return count < half
}

// Filter log using supplied filter decision logic
func filterLog(log []string, filter filterByOneDecision) string {
	candidates := log
	bits := len(log[0])

	for i := 0; i < bits; i++ {
		oneCount := CountOnes(candidates, i)
		half := float64(len(candidates)) / 2.0

		if filter(oneCount, half) {
			candidates = FilterByBit(candidates, '1', i)
		} else {
			candidates = FilterByBit(candidates, '0', i)
		}

		if len(candidates) == 1 {
			break
		}
	}

	return candidates[0]
}

func Part2(input []string) int {
	gammaStr := filterLog(input, filterGamma)
	gamma := StrToBinary(gammaStr)

	epsilonStr := filterLog(input, filterEpsilon)
	epsilon := StrToBinary(epsilonStr)

	return gamma * epsilon
}
