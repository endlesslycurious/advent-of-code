package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Represents a cave of Node in the graph
type Node struct {
	name        string  // Name of the cave
	large       bool    // true if large cave
	connections []*Node // list of connected caves/nodes
}

// Create node graph from input lines
func CreateNodeGraph(lines []string) []Node {
	nodeMap := make(map[string]*Node)

	for _, line := range lines {

		caves := strings.Split(line, "-")
		if len(caves) != 2 {
			log.Fatalln("Expected two caves only, not", len(caves))
		}

		// create nodes in map and list
		for _, cave := range caves {
			_, exists := nodeMap[cave]

			if !exists {
				big := cave == strings.ToUpper(cave)
				node := &Node{cave, big, make([]*Node, 0)}
				nodeMap[cave] = node
			}
		}

		// connect nodes
		start := caves[0]
		end := caves[1]

		nodeMap[start].connections = append(nodeMap[start].connections, nodeMap[end])
		nodeMap[end].connections = append(nodeMap[end].connections, nodeMap[start])
	}

	graph := make([]Node, 0, len(nodeMap))

	for _, node := range nodeMap {
		graph = append(graph, *node)
	}

	return graph
}

func ReadInput(filename string) []Node {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	graph := CreateNodeGraph(lines)

	return graph
}

func main() {
	filename := "./2021/12/day12_input.txt"
	graph := ReadInput(filename)
	fmt.Println("Read", len(graph), "nodes from", filename)

	answer := Part1(graph)
	fmt.Println("Part 1 Answer:", answer)
}

func Part1(graph []Node) int {
	return 0
}
