package day13

import (
	"testing"
)

var input = `
/->-\
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/
`

func TestPart1(t *testing.T) {
	expected := "7,3"
	actual, _ := Part1(input)
	if actual != expected {
		t.Errorf("expected %s actual %s", expected, actual)
	}
}

func TestPart2(t *testing.T) {

}
