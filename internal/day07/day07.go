package day07

import (
	"container/heap"
	"fmt"
	"regexp"
	"strings"

	"github.com/orn688/advent-of-code-2018/internal/util"
)

var lineRegex = regexp.MustCompile(`^Step (?P<Dependency>[A-Z]+) must be ` +
	`finished before step (?P<Depender>[A-Z]+) can begin.$`)

type requirement struct {
	// Dependency must be completed before Depender.
	Depender   string
	Dependency string
}

type dependencyGraph map[string][]string

// Part1 is unimplemented
func Part1(input string) (string, error) {
	reqs, err := parseInput(input)
	if err != nil {
		return "", err
	}
	graph := buildGraph(reqs)
	order, err := topoOrder(graph)
	if err != nil {
		return "", err
	}
	return strings.Join(order, ""), nil
}

// Part2 is unimplemented
func Part2(input string) (string, error) {
	return "", nil
}

func parseInput(input string) ([]requirement, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	reqs := make([]requirement, len(lines))
	for i, line := range lines {
		req, err := parseLine(line)
		if err != nil {
			return reqs, err
		}
		reqs[i] = req
	}
	return reqs, nil
}

// Convert a line to a requirement, assuming each step name is a string of
// capital ASCII letters.
func parseLine(line string) (req requirement, err error) {
	groups, err := util.CaptureRegexGroups(lineRegex, line)
	if err != nil {
		return req, err
	}
	req = requirement{
		Depender:   groups["Depender"],
		Dependency: groups["Dependency"],
	}
	return
}

func buildGraph(reqs []requirement) dependencyGraph {
	graph := make(map[string][]string)
	for _, req := range reqs {
		graph[req.Dependency] = append(graph[req.Dependency], req.Depender)
	}
	return graph
}

func topoOrder(graph dependencyGraph) ([]string, error) {
	ordering := make([]string, len(graph))

	dependencyCounts := getDependencyCounts(graph)
	stepsWithoutDependencies := getStepsWithoutDependencies(dependencyCounts)

	for len(stepsWithoutDependencies) > 0 {
		step := heap.Pop(&stepsWithoutDependencies).(string)
		for _, dependent := range graph[step] {
			dependencyCounts[dependent]--
			if dependencyCounts[dependent] == 0 {
				heap.Push(&stepsWithoutDependencies, dependent)
			}
		}
		ordering = append(ordering, step)
	}

	if ordering[len(ordering)-1] == "" {
		// We got to a point where there are no uncompleted steps whose
		// dependencies are all completed, but not all steps have been put
		// into the ordering - i.e., a cycle.
		return []string{}, fmt.Errorf("cycle detected")
	}
	return ordering, nil
}

func getDependencyCounts(graph dependencyGraph) map[string]int {
	dependencyCounts := make(map[string]int, len(graph))
	for step := range graph {
		dependencyCounts[step] = 0
	}
	for _, dependentSteps := range graph {
		for _, step := range dependentSteps {
			dependencyCounts[step]++
		}
	}
	return dependencyCounts
}

func getStepsWithoutDependencies(depCounts map[string]int) util.StringHeap {
	stepsWithoutDependencies := make(util.StringHeap, 0)
	for step, dependencyCount := range depCounts {
		if dependencyCount == 0 {
			stepsWithoutDependencies.Push(step)
		}
	}
	// Use a heap so that steps are ordered alphabetically if there's a "tie".
	heap.Init(&stepsWithoutDependencies)
	return stepsWithoutDependencies
}
