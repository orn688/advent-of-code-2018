package day02

import (
	"errors"
	"strconv"
	"strings"
)

// Part1 returns the checksum of the list of box IDs.
func Part1(input string) (string, error) {
	boxIDs := parseInput(input)
	numWithTripleLetter, numWithDoubleLetter := 0, 0

	for _, boxID := range boxIDs {
		letterCounts := getLetterCounts(boxID)
		hasDoubleLetter, hasTripleLetter := false, false
		for _, count := range letterCounts {
			hasDoubleLetter = hasDoubleLetter || count == 2
			hasTripleLetter = hasTripleLetter || count == 3
		}

		if hasDoubleLetter {
			numWithDoubleLetter++
		}
		if hasTripleLetter {
			numWithTripleLetter++
		}
	}

	checksum := numWithDoubleLetter * numWithTripleLetter

	return strconv.Itoa(checksum), nil
}

// Part2 returns the characters shared by the two box IDs that differ by a
// single character.
func Part2(input string) (string, error) {
	boxIDs := parseInput(input)

	for _, boxID := range boxIDs {
		for _, otherBoxID := range boxIDs {
			diffCount := 0
			for i, char := range boxID {
				if rune(otherBoxID[i]) != char {
					diffCount++
					if diffCount > 1 {
						break
					}
				}
			}

			if diffCount == 1 {
				return sharedLetters(boxID, otherBoxID), nil
			}
		}
	}

	return "", errors.New("invalid input, no nearby box ids")
}

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func getLetterCounts(str string) map[rune]int {
	counts := make(map[rune]int)
	for _, letter := range str {
		if _, exists := counts[letter]; !exists {
			counts[letter] = 0
		}
		counts[letter]++
	}
	return counts
}

func sharedLetters(str1 string, str2 string) string {
	builder := strings.Builder{}
	for i, char := range str1 {
		if rune(str2[i]) == char {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}
