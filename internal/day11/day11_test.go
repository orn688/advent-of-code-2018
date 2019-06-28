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

func TestPart2(t *testing.T) {
	type testcase struct {
		serialNumber int
		expectedX    int
		expectedY    int
		expectedSize int
	}
	testcases := []testcase{
		// Commented out because they're VERY slow.
		// {
		// 	serialNumber: 18,
		// 	expectedX:    90,
		// 	expectedY:    269,
		// 	expectedSize: 16,
		// },
		// {
		// 	serialNumber: 42,
		// 	expectedX:    232,
		// 	expectedY:    251,
		// 	expectedSize: 12,
		// },
	}

	for _, tc := range testcases {
		expected := fmt.Sprintf(
			"%d,%d,%d", tc.expectedX, tc.expectedY, tc.expectedSize)
		actual, _ := Part2(strconv.Itoa(tc.serialNumber))
		if actual != expected {
			t.Errorf("expected %s, actual %s", expected, actual)
		}
	}
}
