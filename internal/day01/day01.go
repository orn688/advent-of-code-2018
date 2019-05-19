package day01

import (
	"fmt"
	"strconv"
	"strings"
)

// Part1 returns the frequency at the end.
func Part1(input string) (int, error) {
	diffs, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	frequency := 0
	for _, diff := range diffs {
		frequency += diff
	}
	return frequency, nil
}

// Part2 returns the frequency at the end.
func Part2(input string) (int, error) {
	diffs, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	seen := make(map[int]bool)
	frequency := 0
	i := 0
	for {
		frequency += diffs[i]
		if seen[frequency] {
			break
		}
		seen[frequency] = true
		i = (i + 1) % len(diffs)
	}
	return frequency, nil
}

func parseInput(input string) ([]int, error) {
	diffs := strings.Split(strings.TrimSpace(input), "\n")
	parsedDiffs := make([]int, len(diffs))
	for i, stringDiff := range diffs {
		diff, err := strconv.Atoi(stringDiff)
		if err != nil {
			return parsedDiffs, fmt.Errorf("invalid input line: %s", err)
		}
		parsedDiffs[i] = diff
	}
	return parsedDiffs, nil
}
