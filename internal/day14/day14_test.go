package day14

import (
	"testing"
)

func assertCorrect(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("got '%s', expected '%s'", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	t.Run("5 recipes", func(t *testing.T) {
		result, _ := Part1("5")
		assertCorrect(t, result, "0124515891")
	})

	t.Run("9 recipes", func(t *testing.T) {
		result, _ := Part1("9")
		assertCorrect(t, result, "5158916779")
	})

	t.Run("18 recipes", func(t *testing.T) {
		result, _ := Part1("18")
		assertCorrect(t, result, "9251071085")
	})

	t.Run("2018 recipes", func(t *testing.T) {
		result, _ := Part1("2018")
		assertCorrect(t, result, "5941429882")
	})
}

func TestPart2(t *testing.T) {
	t.Run("sequence 101", func(t *testing.T) {
		result, _ := Part2("101")
		assertCorrect(t, result, "2")
	})

	t.Run("sequence 01", func(t *testing.T) {
		result, _ := Part2("01")
		assertCorrect(t, result, "3")
	})

	t.Run("sequence 01245", func(t *testing.T) {
		result, _ := Part2("01245")
		assertCorrect(t, result, "5")
	})

	t.Run("sequence 51589", func(t *testing.T) {
		result, _ := Part2("51589")
		assertCorrect(t, result, "9")
	})

	t.Run("sequence 92510", func(t *testing.T) {
		result, _ := Part2("92510")
		assertCorrect(t, result, "18")
	})

	t.Run("sequence 59414", func(t *testing.T) {
		result, _ := Part2("59414")
		assertCorrect(t, result, "2018")
	})
}
