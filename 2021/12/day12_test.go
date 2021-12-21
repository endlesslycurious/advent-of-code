package main

import (
	"testing"
)

var input = []string{
	"start-A",
	"start-b",
	"A-c",
	"A-b",
	"b-d",
	"A-end",
	"b-end",
}

func TestCreateGraphSize(t *testing.T) {
	graph := CreateNodeGraph(input)

	actual := len(graph)
	expected := 6

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCreateGraphNames(t *testing.T) {
	graph := CreateNodeGraph(input)
	expected := []string{"start", "A", "b", "c", "d", "end"}

	for _, name := range expected {
		found := false
		for _, node := range graph {
			if node.name == name {
				found = true
				break
			}
		}

		if !found {
			t.Error("Expected", name, "not found")
			break
		}
	}
}

func TestCreateGraphLarge(t *testing.T) {
	graph := CreateNodeGraph(input)
	expected := map[string]bool{
		"start": false,
		"A":     true,
		"b":     false,
		"c":     false,
		"d":     false,
		"end":   false,
	}

	for _, node := range graph {
		large, exists := expected[node.name]

		if !exists {
			t.Error("Unexpected", node.name, "node")
		}

		if node.large != large {
			t.Error("Expected node", node.name, "large", large, "got", node.large)
		}
	}
}

func TestCreateGraphConnectionsCount(t *testing.T) {
	graph := CreateNodeGraph(input)
	expected := map[string]int{
		"start": 2,
		"A":     4,
		"b":     4,
		"c":     1,
		"d":     1,
		"end":   2,
	}

	for _, node := range graph {
		count, exists := expected[node.name]

		if !exists {
			t.Error("Unexpected", node.name, "node")
		}

		actual := len(node.connections)
		if actual != count {
			t.Error("Expected node", node.name, "connection", count, "got", actual)
		}
	}
}

func TestPart1(t *testing.T) {
	expected := 10
	graph := CreateNodeGraph(input)
	actual := Part1(graph)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
