package day12

import (
	"container/list"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

const hasPlant = '#'
const noPlant = '.'
const patternLength uint = 5

type plantPot struct {
	hasPlant bool
	number   int
}

// Part1 returns the sum of the pot numbers of all plants with pots after 20
// generations.
func Part1(input string) (string, error) {
	pots, rules := parseInput(input)
	sum := sumAfterGenerations(pots, rules, 20)
	return strconv.Itoa(sum), nil
}

// Part2 returns the sum of the pot numbers of all plants with pots after 50
// billion generations.
func Part2(input string) (string, error) {
	f, err := os.Create("cpuprofile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	pots, rules := parseInput(input)
	sum := sumAfterGenerations(pots, rules, 50000000000)
	return strconv.Itoa(sum), nil
}

func sumAfterGenerations(pots *list.List, rules []bool, gens int) int {
	type arrangementInstance struct {
		generation     int
		firstPotNumber int
	}
	seenArrangements := make(map[string]arrangementInstance)
	foundLoop := false
	for g := 0; g < gens; g++ {
		firstPot := pots.Front().Value.(plantPot)
		str := potString(pots)
		firstInstance, seen := seenArrangements[str]
		if !foundLoop && seen {
			foundLoop = true

			loopLength := g - firstInstance.generation
			potDiff := firstPot.number - firstInstance.firstPotNumber

			// Simulate skipping to the `gens - 1` generation.
			loopCount := (gens - 1 - g) / loopLength
			g += (loopCount * loopLength)
			for e := pots.Front(); e != nil; e = e.Next() {
				pot := e.Value.(plantPot)
				pot.number += potDiff * loopCount
				e.Value = pot
			}
		} else if !seen {
			seenArrangements[str] = arrangementInstance{
				generation:     g,
				firstPotNumber: firstPot.number,
			}
		}
		pots = nextGeneration(pots, rules)
	}

	sum := 0
	for e := pots.Front(); e != nil; e = e.Next() {
		pot := e.Value.(plantPot)
		if pot.hasPlant {
			sum += pot.number
		}
	}
	return sum
}

func parseInput(input string) (*list.List, []bool) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	initialStateString := strings.TrimLeft(lines[0], "initial state: ")
	initialPots := potsFromString(initialStateString)

	rules := make([]bool, 1<<patternLength)
	for i := 2; i < len(lines); i++ {
		pattern, causesPlant := parseRule(lines[i])
		rules[pattern] = causesPlant
	}

	return initialPots, rules
}

func nextGeneration(pots *list.List, rules []bool) *list.List {
	nextPots := list.New()
	// Assumes there will always be a pot with a plant.
	firstPot := pots.Front().Value.(plantPot)
	lastPot := pots.Back().Value.(plantPot)
	potTwoToRight := pots.Front()
	pattern := 0
	latestPotWithPlant := lastPot.number // High dummy starting value.
	for num := firstPot.number - 2; num <= lastPot.number+2; num++ {
		pattern = updatePattern(pattern, potTwoToRight)
		causesPlant := rules[pattern]

		if causesPlant {
			for i := latestPotWithPlant + 1; i < num; i++ {
				nextPots.PushBack(plantPot{
					hasPlant: false,
					number:   i,
				})
			}
			nextPots.PushBack(plantPot{
				hasPlant: causesPlant,
				number:   num,
			})
			latestPotWithPlant = num
		}

		if potTwoToRight != nil {
			potTwoToRight = potTwoToRight.Next()
		}
	}

	return nextPots
}

func updatePattern(pattern int, potTwoToRight *list.Element) int {
	hasPlant := 0
	if potTwoToRight != nil && potTwoToRight.Value.(plantPot).hasPlant {
		hasPlant = 1
	}
	pattern <<= 1
	pattern &= (1 << patternLength) - 1
	pattern |= hasPlant
	return pattern
}

func potsFromString(stateString string) *list.List {
	pots := list.New()
	for i, char := range stateString {
		pot := plantPot{
			hasPlant: (char == hasPlant),
			number:   i,
		}
		pots.PushBack(pot)
	}
	return pots
}

func parseRule(ruleString string) (int, bool) {
	parts := strings.Split(ruleString, " => ")
	patternString := parts[0]
	pattern := 0
	for _, char := range patternString {
		pattern <<= 1
		if char == hasPlant {
			pattern |= 1
		}
	}
	causesPlant := (parts[1] == string(hasPlant))
	return pattern, causesPlant
}

func potString(pots *list.List) string {
	potStrings := make([]rune, pots.Len())
	minPot := pots.Front().Value.(plantPot).number
	for e := pots.Front(); e != nil; e = e.Next() {
		pot := e.Value.(plantPot)
		value := noPlant
		if pot.hasPlant {
			value = hasPlant
		}
		potStrings[pot.number-minPot] = value
	}
	return string(potStrings)
}
