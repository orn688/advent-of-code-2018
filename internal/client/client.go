package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
)

const cacheDirName = ".aoc_cache"

// GetInput fetches and returns the AoC input for the given day. It maintains a
// local cache of the input for each day (in .aoc_cache/<day>) to avoid making
// redundant requests to the AoC server.
func GetInput(day int) (string, error) {
	input, err := getInputFromCache(day)
	if err != nil {
		input, err = requestInput(day)
		if err != nil {
			return "", err
		}
		saveInputToCache(day, input)
	}
	return input, nil
}

func requestInput(day int) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/2018/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	sessionID := os.Getenv("AOC_SESSION_ID")
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionID})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	} else if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to fetch puzzle input: %s", resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getInputFromCache(day int) (string, error) {
	fileName, err := cachedFileForDay(day)
	if err != nil {
		return "", err
	}
	rawInput, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	input := string(rawInput)
	return input, nil
}

func saveInputToCache(day int, input string) error {
	cacheDir, err := getCacheDir()
	if err != nil {
		return err
	}
	err = os.MkdirAll(cacheDir, 0644)
	if err != nil {
		return err
	}
	fileName, err := cachedFileForDay(day)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, []byte(input), 0644)
}

func cachedFileForDay(day int) (string, error) {
	cacheDir, err := getCacheDir()
	if err != nil {
		return "", err
	}

	fileName := path.Join(cacheDir, strconv.Itoa(day))
	return fileName, nil
}

func getCacheDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Join(currentDir, cacheDirName), nil
}
