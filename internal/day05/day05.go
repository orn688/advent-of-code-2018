package day05

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Part1 returns the length of the collapsed version of the input polymer.
func Part1(input string) (string, error) {
	units, err := parseInput(input)
	if err != nil {
		return "", err
	}

	collapse(units)
	return strconv.Itoa(units.Len()), nil
}

// Part2 returns the shortest possible polymer length after removing all
// instances of any one chosen character (case-insensitive) and then collapsing
// the polymer.
func Part2(input string) (string, error) {
	baseUnits, err := parseInput(input)
	if err != nil {
		return "", err
	}
	minLength := baseUnits.Len()
	for char := 'a'; char <= 'z'; char++ {
		units := copy(baseUnits)
		removeChar(units, char)
		collapse(units)
		if units.Len() < minLength {
			minLength = units.Len()
		}
	}
	return strconv.Itoa(minLength), nil
}

func parseInput(input string) (*list.List, error) {
	units := list.New()
	for _, char := range strings.TrimSpace(input) {
		if !('a' <= char && char <= 'z') && !('A' <= char && char <= 'Z') {
			return units, fmt.Errorf("invalid character: '%c'", char)
		}
		units.PushBack(char)
	}
	return units, nil
}

// Collapses the polymer in-place.
func collapse(units *list.List) {
	current := units.Front()
	for current != nil {
		next := current.Next()
		if next == nil {
			break
		}
		if canCollapse(current, next) {
			newCurrent := current.Prev()
			if newCurrent == nil {
				newCurrent = next.Next()
			}
			units.Remove(next)
			units.Remove(current)
			current = newCurrent
		} else {
			current = next
		}
	}
}

func canCollapse(left *list.Element, right *list.Element) bool {
	leftChar, rightChar := left.Value.(int32), right.Value.(int32)
	// The lower and upper versions of character are 32 codepoints apart.
	// For example, A = 65 and a = 97.
	return abs(leftChar-rightChar) == 32
}

func abs(x int32) int32 {
	if x < 0 {
		return -1 * x
	}
	return x
}

func copy(lst *list.List) *list.List {
	listCopy := list.New()
	for element := lst.Front(); element != nil; element = element.Next() {
		listCopy.PushBack(element.Value)
	}
	return listCopy
}

// Remove all instances of char in place (case-insensitive).
func removeChar(units *list.List, char rune) {
	element := units.Front()
	upper, lower := unicode.ToUpper(char), unicode.ToLower(char)
	for element != nil {
		next := element.Next()
		value := element.Value.(rune)
		if value == upper || value == lower {
			units.Remove(element)
		}
		element = next
	}
}
