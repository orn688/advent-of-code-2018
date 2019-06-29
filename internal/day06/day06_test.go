package day06

import (
	"strconv"
	"testing"
)

var input = `
1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`

func TestPart1(t *testing.T) {
	expected := 17
	actual, _ := Part1(input)
	if actual != strconv.Itoa(expected) {
		t.Errorf("expected %d, actual %s", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 16
	actual, _ := centralAreaSize(input, 32)
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
