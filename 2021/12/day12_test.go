package main

import "testing"

var input = []string{
	"start-A",
	"start-b",
	"A-c",
	"A-b",
	"b-d",
	"A-end",
	"b-end",
}

func TestPart1(t *testing.T) {
	expected := 10
	actual := Part1([]Node{})

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
