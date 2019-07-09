package day14

import (
	"strconv"
	"strings"
)

const answerLength = 10

// Scoreboard represents the state of the scoreboard of recipes that have been
// tried so far.
type Scoreboard struct {
	elf1, elf2 int
	Scores     []int
}

func newScoreBoard(recipeAttempts int) *Scoreboard {
	var scores []int
	if recipeAttempts < 0 {
		scores = make([]int, 2)
	} else {
		// There may be a trailing recipe, hence the +1.
		scores = make([]int, 2, recipeAttempts+answerLength+1)
	}
	scores[0], scores[1] = 3, 7
	return &Scoreboard{
		elf1:   0,
		elf2:   1,
		Scores: scores,
	}
}

func (b *Scoreboard) isFull() bool {
	return len(b.Scores)+1 >= cap(b.Scores)
}

func (b *Scoreboard) step() {
	score1, score2 := b.Scores[b.elf1], b.Scores[b.elf2]
	sum := score1 + score2
	for _, newCharScore := range strconv.Itoa(sum) {
		newScore, _ := strconv.Atoi(string(newCharScore))
		b.Scores = append(b.Scores, newScore)
	}
	b.elf1 = (b.elf1 + score1 + 1) % len(b.Scores)
	b.elf2 = (b.elf2 + score2 + 1) % len(b.Scores)
}

func (b *Scoreboard) length() int {
	return len(b.Scores)
}

// Part1 returns the scores of the next 10 recipes after the given number of
// recipe attempts.
func Part1(input string) (string, error) {
	recipeAttempts, err := parseInput(input)
	if err != nil {
		return "", err
	}
	scoreboard := newScoreBoard(recipeAttempts)
	for !scoreboard.isFull() {
		scoreboard.step()
	}
	var sb strings.Builder
	for i := recipeAttempts; i < recipeAttempts+answerLength; i++ {
		sb.WriteString(strconv.Itoa(scoreboard.Scores[i]))
	}
	return sb.String(), nil
}

// Part2 returns the number of recipes it takes until the given sequence of
// recipe scores is seen.
func Part2(input string) (string, error) {
	targetSequence := strings.TrimSpace(input)
	if _, err := strconv.Atoi(targetSequence); err != nil {
		return "", err
	}
	scoreboard := newScoreBoard(-1)
	inputLength := len(targetSequence)
	for {
		oldLength := scoreboard.length()
		scoreboard.step()
		newLength := scoreboard.length()
		if newLength < inputLength {
			continue
		}
		lengthDiff := newLength - oldLength // Could be either 1 or 2
		for offset := 0; offset < lengthDiff; offset++ {
			foundMatch := true
			for i := 1; i <= inputLength; i++ {
				expectedDigit, _ := strconv.Atoi(
					string(targetSequence[inputLength-i]))
				actualDigit := scoreboard.Scores[scoreboard.length()-i-offset]
				if actualDigit != expectedDigit {
					foundMatch = false
					break
				}
			}
			if foundMatch {
				recipesBeforeMatch := scoreboard.length() - inputLength - offset
				return strconv.Itoa(recipesBeforeMatch), nil
			}
		}
	}
}

func parseInput(input string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(input))
}

// A simple integer exponentiation function. Optimized for simplicity rather
// than speed.
func pow(base, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= base
	}
	return result
}
