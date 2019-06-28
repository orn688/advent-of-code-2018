package day10

import (
	"strings"
	"testing"
)

var testInput = `position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>
`

func TestPart1(t *testing.T) {
	expected := `
#   #  ###
#   #   #
#   #   #
#####   #
#   #   #
#   #   #
#   #   #
#   #  ###
	`

	actual, _ := Part1(testInput)
	if !stringEqualsIgnoringWhitespace(actual, expected) {
		t.Errorf("expected:\n%s\n, actual:\n%s\n", expected, actual)
	}
}

func stringEqualsIgnoringWhitespace(str1 string, str2 string) bool {
	lines1, lines2 := splitLines(str1), splitLines(str2)
	if len(lines1) != len(lines2) {
		return false
	}
	for i := range lines1 {
		if strings.TrimSpace(lines1[i]) != strings.TrimSpace(lines2[i]) {
			return false
		}
	}
	return true
}

func splitLines(str string) []string {
	return strings.Split(strings.TrimSpace(str), "\n")
}

func TestPart2(t *testing.T) {
	expected := "3"
	actual, _ := Part2(testInput)
	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}
