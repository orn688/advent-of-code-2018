package day08

import (
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	testcases := map[string]int{
		"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2": 138,
	}
	for input, expected := range testcases {
		actual, _ := Part1(input)
		if actual != strconv.Itoa(expected) {
			t.Errorf("expected %d, actual %s", expected, actual)
		}
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]int{
		"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2": 66,
	}
	for input, expected := range testcases {
		actual, _ := Part2(input)
		if actual != strconv.Itoa(expected) {
			t.Errorf("expected %d, actual %s", expected, actual)
		}
	}
}
