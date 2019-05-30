package day03

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/orn688/advent-of-code-2018/internal/util"
)

type coordinate struct {
	X int
	Y int
}

type fabricClaim struct {
	ID       int
	LeftDist int
	TopDist  int
	Width    int
	Height   int
}

func (c fabricClaim) allCoordinates() []coordinate {
	coordinates := make([]coordinate, c.Width*c.Height)

	for yOffset := 0; yOffset < c.Height; yOffset++ {
		for xOffset := 0; xOffset < c.Width; xOffset++ {
			x, y := xOffset+c.LeftDist, yOffset+c.TopDist
			index := xOffset + (yOffset * c.Width)
			coordinates[index] = coordinate{x, y}
		}
	}

	return coordinates
}

// Part1 returns the number of disputed squares (>= 2 claims).
func Part1(input string) (string, error) {
	claims, err := parseInput(input)
	if err != nil {
		return "", err
	}

	claimCounts := getClaimCounts(claims)

	numDisputedSquares := 0
	for _, count := range claimCounts {
		if count > 1 {
			numDisputedSquares++
		}
	}

	return strconv.Itoa(numDisputedSquares), nil
}

// Part2 returns the ID of the only claim that doesn't overlap with any others.
func Part2(input string) (string, error) {
	claims, err := parseInput(input)
	if err != nil {
		return "", err
	}

	claimCounts := getClaimCounts(claims)

	for _, claim := range claims {
		conflict := false
		for _, coord := range claim.allCoordinates() {
			if claimCounts[coord] > 1 {
				conflict = true
				break
			}
		}

		if !conflict {
			return strconv.Itoa(claim.ID), nil
		}
	}

	return "", errors.New("no matching claim in input")
}

func parseInput(input string) ([]*fabricClaim, error) {
	regex := regexp.MustCompile(`^#(?P<ClaimID>\d+) @ ` +
		`(?P<LeftDist>\d+),(?P<TopDist>\d+): ` +
		`(?P<Width>\d+)x(?P<Height>\d+)$`)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	claims := make([]*fabricClaim, len(lines))

	for i, line := range lines {
		groups, err := util.CaptureRegexGroups(regex, line)
		if err != nil {
			return claims, fmt.Errorf("invalid line: %s", line)
		}
		claimID, err := strconv.Atoi(groups["ClaimID"])
		leftDist, err := strconv.Atoi(groups["LeftDist"])
		topDist, err := strconv.Atoi(groups["TopDist"])
		width, err := strconv.Atoi(groups["Width"])
		height, err := strconv.Atoi(groups["Height"])
		if err != nil {
			return claims, fmt.Errorf("invalid data in line: %s", line)
		}

		claims[i] = &fabricClaim{
			ID:       claimID,
			LeftDist: leftDist,
			TopDist:  topDist,
			Width:    width,
			Height:   height,
		}
	}

	return claims, nil
}

func getClaimCounts(claims []*fabricClaim) map[coordinate]int {
	claimCounts := make(map[coordinate]int)

	for _, claim := range claims {
		for _, coord := range claim.allCoordinates() {
			claimCounts[coord]++
		}
	}

	return claimCounts
}
