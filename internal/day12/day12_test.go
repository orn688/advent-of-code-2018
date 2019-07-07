package day12

import (
	"strconv"
	"testing"
)

const input = `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
`

func TestPart1(t *testing.T) {
	expected := 325
	actual, _ := Part1(input)
	if actual != strconv.Itoa(expected) {
		t.Errorf("expected %d, actual %s", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	// This expected value wasn't given by AoC, but is produced by a version of
	// the program that produced the correct value given the real input.
	expected := 999999999374
	actual, _ := Part2(input)
	if actual != strconv.Itoa(expected) {
		t.Errorf("expected %d, actual %s", expected, actual)
	}
}
