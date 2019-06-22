package day09

import (
	"container/list"
	"regexp"
	"strconv"

	"github.com/orn688/advent-of-code-2018/internal/util"
)

var inputRegex = regexp.MustCompile(`(?P<PlayerCount>\d+) players; last ` +
	`marble is worth (?P<LastMarble>\d+) points`)

// Part1 returns the max score of any player after playing the game with the
// given number of players and marbles.
func Part1(input string) (string, error) {
	playerCount, lastMarble, err := parseInput(input)
	if err != nil {
		return "", err
	}
	scores := playGame(playerCount, lastMarble)
	return strconv.Itoa(max(scores)), nil
}

// Part2 is unimplemented
func Part2(input string) (string, error) {
	return "", nil
}

func playGame(playerCount, lastMarble int) []int {
	marbles := list.New()
	scores := make([]int, playerCount)
	var currentMarble *list.Element
	for marble := 0; marble <= lastMarble; marble++ {
		if currentMarble == nil {
			currentMarble = marbles.PushBack(marble)
		} else if marble%23 == 0 {
			player := marble % playerCount
			scores[player] += marble
			var nextCurrentMarble *list.Element
			for i := 0; i < 7; i++ {
				nextCurrentMarble = currentMarble
				currentMarble = currentMarble.Prev()
				if currentMarble == nil {
					currentMarble = marbles.Back()
				}
			}
			scores[player] += marbles.Remove(currentMarble).(int)
			currentMarble = nextCurrentMarble
		} else {
			currentMarble = currentMarble.Next()
			if currentMarble == nil {
				currentMarble = marbles.Front()
			}
			currentMarble = marbles.InsertAfter(marble, currentMarble)
		}
	}
	return scores
}

func parseInput(input string) (playerCount int, lastMarble int, err error) {
	groups, err := util.CaptureRegexGroups(inputRegex, input)
	if err != nil {
		return
	}
	playerCount, err = strconv.Atoi(groups["PlayerCount"])
	lastMarble, err = strconv.Atoi(groups["LastMarble"])
	return
}

func max(nums []int) (maximum int) {
	maximum = nums[0]
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
	}
	return
}
