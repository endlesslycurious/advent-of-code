package main

import "fmt"

// Represents a cave of Node in the graph
type Node struct {
	name        string  // Name of the cave
	large       bool    // true if large cave
	connections []*Node // list of connected caves/nodes
}

func ReadInput(filename string) []Node {
	return []Node{}
}

func main() {
	filename := "./2021/12/day12_input.txt"
	input := ReadInput(filename)
	fmt.Println("Read", len(input), "nodes from", filename)

	answer := Part1(input)
	fmt.Println("Part 1 Answer:", answer)
}

func Part1(nodes []Node) int {
	return 0
}
