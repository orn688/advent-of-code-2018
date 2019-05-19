package day01

import (
	"fmt"
	"strconv"
	"strings"
)

// Part1 returns the frequency at the end.
func Part1(input string) (string, error) {
	diffs, err := parseInput(input)
	if err != nil {
		return "", err
	}
	frequency := 0
	for _, diff := range diffs {
		frequency += diff
	}
	return strconv.Itoa(frequency), nil
}

// Part2 returns the first frequency to be hit twice.
func Part2(input string) (string, error) {
	diffs, err := parseInput(input)
	if err != nil {
		return "", err
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
	return strconv.Itoa(frequency), nil
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
