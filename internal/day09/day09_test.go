package day09

import (
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	testcases := map[string]int{
		"10 players; last marble is worth 1618 points": 8317,
		"13 players; last marble is worth 7999 points": 146373,
		"17 players; last marble is worth 1104 points": 2764,
		"21 players; last marble is worth 6111 points": 54718,
		"30 players; last marble is worth 5807 points": 37305,
	}
	for input, expected := range testcases {
		actual, _ := Part1(input)
		if actual != strconv.Itoa(expected) {
			t.Errorf("expected %d, actual %s", expected, actual)
		}
	}
}
