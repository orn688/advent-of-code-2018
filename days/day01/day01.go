package day01

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (frequency int) {
	diffs := strings.Split(strings.TrimSpace(input), "\n")
	frequency = 0
	for _, stringDiff := range diffs {
		diff, err := strconv.Atoi(stringDiff)
		if err != nil {
			fmt.Sprintln("Invalid input:", err)
			return
		}
		frequency += diff
	}
	return
}

func Part2(input string) int {
	return 1
}
