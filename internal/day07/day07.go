package day07

import (
	"container/heap"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/orn688/advent-of-code-2018/internal/util"
)

var lineRegex = regexp.MustCompile(`^Step (?P<Dependency>[A-Z]) must be ` +
	`finished before step (?P<Depender>[A-Z]) can begin.$`)

// A Requirement is a 2-tuple made up of a step name (the Depender) and the
// name of a step that it depends on (the Dependency).
type requirement struct {
	// Dependency must be completed before Depender.
	Depender   string
	Dependency string
}

type stepInProgress struct {
	Step     string
	TimeLeft int
}

// A dependencyGraph is an adjacency list mapping step names to the names of
// steps that depend on them.
// ReverseAlphaOrder indicates whether a step named B should be built before a
// step named A, all else being equal
type dependencyGraph struct {
	adjList                 map[string][]string
	stepsReadyToBuild       util.StringHeap
	unbuiltDependencyCounts map[string]int
	reverseAlphaOrder       bool
}

func newDependencyGraph(reqs []requirement, reverse bool) dependencyGraph {
	graph := dependencyGraph{
		adjList:                 make(map[string][]string),
		stepsReadyToBuild:       make(util.StringHeap, 0),
		unbuiltDependencyCounts: make(map[string]int, len(reqs)),
		reverseAlphaOrder:       reverse,
	}
	for _, req := range reqs {
		graph.adjList[req.Dependency] = append(
			graph.adjList[req.Dependency], req.Depender,
		)
	}
	graph.setDependencyCounts()
	graph.setStepsReadyToBuild()
	return graph
}

// Set the number of dependencies for each build step.
func (graph *dependencyGraph) setDependencyCounts() {
	for step := range graph.adjList {
		graph.unbuiltDependencyCounts[step] = 0
	}
	for _, dependentSteps := range graph.adjList {
		for _, step := range dependentSteps {
			graph.unbuiltDependencyCounts[step]++
		}
	}
}

// Initialize the steps with no dependencies. This must be called after the
// dependency counts have been initialized by calling setDependencyCounts.
func (graph *dependencyGraph) setStepsReadyToBuild() {
	for step, dependencyCount := range graph.unbuiltDependencyCounts {
		if dependencyCount == 0 {
			graph.markAsReadyToBuild(step)
		}
	}
}

func (graph *dependencyGraph) markAsReadyToBuild(step string) {
	priority := int(step[0])
	if graph.reverseAlphaOrder {
		priority *= -1
	}
	heap.Push(&graph.stepsReadyToBuild, util.HeapElement{
		Value:    step,
		Priority: int(step[0]),
	})
}

func (graph *dependencyGraph) stepWasBuilt(step string) {
	for _, dependent := range graph.adjList[step] {
		graph.unbuiltDependencyCounts[dependent]--
		if graph.unbuiltDependencyCounts[dependent] == 0 {
			graph.markAsReadyToBuild(dependent)
		}
	}
}

// The second return value is the flag to indicate whether a buildable step was
// returned.
func (graph *dependencyGraph) getNextStep() (string, bool) {
	if len(graph.stepsReadyToBuild) == 0 {
		return "", false
	}
	element := heap.Pop(&graph.stepsReadyToBuild).(util.HeapElement)
	return element.Value, true
}

// Part1 returns a topological ordering of the steps, based on the requirements
// in the input.
func Part1(input string) (string, error) {
	reqs, err := parseInput(input)
	if err != nil {
		return "", err
	}
	graph := newDependencyGraph(reqs, false)
	ordering := make([]string, len(graph.adjList))

	for i := 0; i < len(ordering); i++ {
		nextStep, stepAvailable := graph.getNextStep()
		if !stepAvailable {
			// We got to a point where there are no uncompleted steps whose
			// dependencies are all completed, but not all steps have been put
			// into the ordering - i.e., a cycle.
			return "", fmt.Errorf("cycle detected")
		}
		ordering[i] = nextStep
		graph.stepWasBuilt(nextStep)
	}

	return strings.Join(ordering, ""), nil
}

// Part2 returns the time it would take workerCount workers to complete the
// steps.
func Part2(input string) (string, error) {
	reqs, err := parseInput(input)
	if err != nil {
		return "", err
	}
	graph := newDependencyGraph(reqs, true)

	workerCount := 5
	currentSteps := make([]stepInProgress, workerCount)
	complete := false
	time := -1
	for !complete {
		time++
		// Assume, then verify
		complete = true
		for i := range currentSteps {
			if currentSteps[i].TimeLeft == 0 {
				oldStep := currentSteps[i].Step
				if oldStep != "" {
					graph.stepWasBuilt(oldStep)
				}
				nextStep, stepAvailable := graph.getNextStep()
				if stepAvailable {
					complete = false
					currentSteps[i] = stepInProgress{
						Step:     nextStep,
						TimeLeft: stepDuration(nextStep) - 1,
					}
				} else {
					currentSteps[i].Step = ""
				}
			} else {
				currentSteps[i].TimeLeft--
				complete = false
			}
		}
	}

	for _, dependencyCount := range graph.unbuiltDependencyCounts {
		if dependencyCount > 0 {
			return "", fmt.Errorf("cycle detected")
		}
	}

	return strconv.Itoa(time), nil
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

// Convert a line to a requirement, assuming each step name is a length-1
// string of capital ASCII letters.
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

// Assumes the step has length 1 and is A-Z.
func stepDuration(step string) int {
	durationOfStepA := 61
	return durationOfStepA + (int(step[0]) - int('A'))
}
