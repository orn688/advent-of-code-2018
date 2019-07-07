package day07

import (
	"testing"
)

const input = `
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`

func TestPart1(t *testing.T) {
	expected := "CABDFE"
	actual, _ := Part1(input)
	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 15
	actual, _ := timeToComplete(input, 2, 1)
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
