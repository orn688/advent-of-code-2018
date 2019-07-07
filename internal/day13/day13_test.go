package day13

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `
/->-\
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/
`
	expected := "7,3"
	actual, _ := Part1(input)
	if actual != expected {
		t.Errorf("expected %s actual %s", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	input := `
/>-<\
|   |
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/
`
	expected := "6,4"
	actual, _ := Part2(input)
	if actual != expected {
		t.Errorf("expected %s actual %s", expected, actual)
	}
}
