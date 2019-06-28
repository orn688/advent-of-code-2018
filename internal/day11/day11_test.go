package day11

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	type testcase struct {
		serialNumber int
		expectedX    int
		expectedY    int
	}
	testcases := []testcase{
		{
			serialNumber: 18,
			expectedX:    33,
			expectedY:    45,
		},
		{
			serialNumber: 42,
			expectedX:    21,
			expectedY:    61,
		},
	}

	for _, tc := range testcases {
		expected := fmt.Sprintf("%d,%d", tc.expectedX, tc.expectedY)
		actual, _ := Part1(strconv.Itoa(tc.serialNumber))
		if actual != expected {
			t.Errorf("expected %s, actual %s", expected, actual)
		}
	}
}
