package day04

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type guardNap struct {
	GuardID   int
	StartTime time.Time
	EndTime   time.Time
}

type guardEvent struct {
	GuardID   int
	EventTime time.Time
	Awake     bool
}

var unknownGuardID = -1

// Part1 returns the id of the guard with the most time spent asleep,
// multiplied by the minute (between 00:00 and 00:50) during which that guard
// was most often asleep.
func Part1(input string) (string, error) {
	naps, err := parseInput(input)
	if err != nil {
		return "", err
	}
	sleepiestGuardID := getSleepiestGuard(naps)
	sleepiestMinute, _ := getSleepiestMinute(naps, sleepiestGuardID)

	return strconv.Itoa(sleepiestGuardID * sleepiestMinute), nil
}

// Part2 returns the ID of the guard who is most frequently asleep at the same
// minute, multiplied by that minute.
func Part2(input string) (string, error) {
	naps, err := parseInput(input)
	if err != nil {
		return "", err
	}
	uniqueGuardIDs := make(map[int]bool)
	for _, nap := range naps {
		uniqueGuardIDs[nap.GuardID] = true
	}
	resultGuardID := -1
	resultMinute := -1
	resultDaysAsleepAtMinute := 0
	for guardID := range uniqueGuardIDs {
		sleepiestMinute, daysAsleepAtMinute := getSleepiestMinute(naps, guardID)
		if daysAsleepAtMinute > resultDaysAsleepAtMinute {
			resultGuardID = guardID
			resultMinute = sleepiestMinute
			resultDaysAsleepAtMinute = daysAsleepAtMinute
		}
	}
	return strconv.Itoa(resultGuardID * resultMinute), nil
}

func parseInput(input string) ([]guardNap, error) {
	naps := []guardNap{}
	events, err := parseEvents(input)
	if err != nil {
		return naps, err
	}
	napInProgress := false
	for _, event := range events {
		if !napInProgress && !event.Awake {
			napInProgress = true
			nap := guardNap{
				GuardID:   event.GuardID,
				StartTime: event.EventTime,
				EndTime:   event.EventTime.Add(time.Duration(-1)), // placeholder
			}
			naps = append(naps, nap)
		}
		if napInProgress && event.Awake {
			// TODO: make sure nap doesn't span multiple days
			napInProgress = false
			naps[len(naps)-1].EndTime = event.EventTime
		}
	}
	return naps, nil
}

func parseEvents(input string) ([]guardEvent, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sort.Strings(lines)
	events := make([]guardEvent, len(lines))
	for i, line := range lines {
		event, err := parseLine(line)
		if err != nil {
			return events, err
		}
		if event.GuardID == unknownGuardID {
			if i == 0 {
				return events, fmt.Errorf("line has unknown guard: %s", line)
			}
			event.GuardID = events[i-1].GuardID
		}
		events[i] = event
	}
	return events, nil
}

func parseLine(line string) (event guardEvent, err error) {
	components := strings.Split(line[1:], "] ")
	if len(components) != 2 {
		err = fmt.Errorf("invalid line: %s", line)
		return
	}
	rawTime, description := components[0], components[1]
	timeFormat := "2006-01-02 15:04"
	eventTime, err := time.Parse(timeFormat, rawTime)
	if err != nil {
		return
	}
	event.EventTime = eventTime
	if err != nil {
		return
	}
	switch description {
	case "wakes up":
		event.Awake = true
		event.GuardID = unknownGuardID
	case "falls asleep":
		event.Awake = false
		event.GuardID = unknownGuardID
	default:
		rawGuardID := strings.TrimSuffix(strings.TrimPrefix(description, "Guard #"), " begins shift")
		guardID, atoiErr := strconv.Atoi(rawGuardID)
		if atoiErr != nil {
			return event, atoiErr
		}
		event.GuardID = guardID
		event.Awake = true
	}
	return
}

func getSleepiestGuard(naps []guardNap) int {
	guardMinutesAsleep := make(map[int]int)
	for _, nap := range naps {
		napDuration := nap.EndTime.Sub(nap.StartTime)
		guardMinutesAsleep[nap.GuardID] += int(napDuration.Minutes())
	}
	sleepiestGuardID := -1
	sleepiestGuardMinutesAsleep := 0
	for guardID, minutesAsleep := range guardMinutesAsleep {
		if minutesAsleep > sleepiestGuardMinutesAsleep {
			sleepiestGuardID = guardID
			sleepiestGuardMinutesAsleep = minutesAsleep
		}
	}
	return sleepiestGuardID
}

// Returns the minute during which the given guard was most often asleep, along
// with the number of days in which the guard was asleep at that minute.
func getSleepiestMinute(naps []guardNap, guardID int) (int, int) {
	asleepMoments := make(map[int]int, 60)
	for _, nap := range naps {
		if nap.GuardID != guardID {
			continue
		}
		// Assumes naps always start and end between 00:00 and 00:59.
		startMinute, endMinute := nap.StartTime.Minute(), nap.EndTime.Minute()
		// Must be up to and *not including* endMinute, because a guard is
		// counted as awake in the minute they wake up.
		for min := startMinute; min < endMinute; min++ {
			asleepMoments[min]++
		}
	}

	sleepiestMinute := 0
	for minute, daysAsleepAtMinute := range asleepMoments {
		if daysAsleepAtMinute > asleepMoments[sleepiestMinute] {
			sleepiestMinute = minute
		}
	}
	return sleepiestMinute, asleepMoments[sleepiestMinute]
}
