package day13

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

type cartDirection rune

const (
	up    cartDirection = '^'
	right cartDirection = '>'
	down  cartDirection = 'v'
	left  cartDirection = '<'
)

var allDirections = []cartDirection{up, right, down, left}
var cartRegex = regexp.MustCompile(`[\^>v<]`)

func directionIndex(dir cartDirection) int {
	for i, otherDir := range allDirections {
		if dir == otherDir {
			return i
		}
	}
	panic("invalid direction")
}

type turnDirection int

const (
	rightTurn turnDirection = 1
	leftTurn  turnDirection = 3
	straight  turnDirection = 4
)

type coord struct {
	X int
	Y int
}

type cart struct {
	Dir          cartDirection
	NextTurn     turnDirection
	LastMoveTime int
}

func newCart(dir cartDirection) *cart {
	return &cart{
		Dir: dir,
		// First turn is always left.
		NextTurn:     leftTurn,
		LastMoveTime: math.MinInt32,
	}
}

// Part1 returns the coordinates of the first crash between two carts.
func Part1(input string) (string, error) {
	cartsGrid, mapRows := parseInput(input)
	x, y := getCollisionLocation(cartsGrid, mapRows)
	return fmt.Sprintf("%d,%d", x, y), nil
}

func getCollisionLocation(cartsGrid [][]*cart, mapRows []string) (int, int) {
	for currentTime := 0; ; currentTime++ {
		for y, row := range mapRows {
			for x := range row {
				c := cartsGrid[y][x]
				if c == nil || c.LastMoveTime == currentTime {
					continue
				}
				nextX, nextY := nextCartLocation(x, y, c, mapRows)
				if cartsGrid[nextY][nextX] != nil {
					return nextX, nextY
				}
				cartsGrid[y][x] = nil
				cartsGrid[nextY][nextX] = c
				c.LastMoveTime = currentTime
			}
		}
	}
}

// Part2 is unimplemented
func Part2(input string) (string, error) {
	return "", nil
}

func parseInput(input string) ([][]*cart, []string) {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")
	cartsGrid := initCartsGrid(rows)
	for y, row := range rows {
		for x, char := range row {
			if isDirection(char) {
				cartsGrid[y][x] = newCart(cartDirection(char))
			}
		}
		rows[y] = cartRegex.ReplaceAllStringFunc(row, func(char string) string {
			if char == string(up) || char == string(down) {
				return "|"
			} else if char == string(left) || char == string(right) {
				return "-"
			}
			// This shouldn't happen.
			panic("invalid regex")
		})
	}
	return cartsGrid, rows
}

func initCartsGrid(rows []string) [][]*cart {
	grid := make([][]*cart, len(rows))
	for y, row := range rows {
		grid[y] = make([]*cart, len(row))
	}
	return grid
}

func isDirection(char rune) bool {
	for _, dir := range allDirections {
		if char == rune(dir) {
			return true
		}
	}
	return false
}

func nextCartLocation(x, y int, c *cart, mapRows []string) (int, int) {
	// 1. Move cart one step in CURRENT direction
	switch c.Dir {
	case up:
		y--
	case down:
		y++
	case left:
		x--
	case right:
		x++
	}
	// 2. Choose new direction
	switch mapRows[y][x] {
	case '+':
		nextDirIndex := (directionIndex(c.Dir) + int(c.NextTurn)) % len(allDirections)
		c.Dir = allDirections[nextDirIndex]
		c.NextTurn = getNextTurn(c.NextTurn)
	case '/':
		switch c.Dir {
		case up:
			c.Dir = right
		case down:
			c.Dir = left
		case left:
			c.Dir = down
		case right:
			c.Dir = up
		}
	case '\\':
		switch c.Dir {
		case up:
			c.Dir = left
		case down:
			c.Dir = right
		case left:
			c.Dir = up
		case right:
			c.Dir = down
		}
	case '|':
		if !(c.Dir == down || c.Dir == up) {
			e := fmt.Sprintf("invalid up/down location: %d,%d", x, y)
			panic(e)
		}
	case '-':
		if !(c.Dir == left || c.Dir == right) {
			e := fmt.Sprintf("invalid left/right location: %d,%d", x, y)
			panic(e)
		}
	default:
		e := fmt.Sprintf("invalid location: %d,%d", x, y)
		panic(e)
	}
	return x, y
}

func getNextTurn(dir turnDirection) turnDirection {
	turnDirs := []turnDirection{leftTurn, straight, rightTurn}
	for i, otherDir := range turnDirs {
		if dir == otherDir {
			return turnDirs[(i+1)%len(turnDirs)]
		}
	}
	panic("invalid turn direction")
}

// Useful for debugging.
func printCarts(cartsGrid [][]*cart, mapRows []string) {
	for y, row := range mapRows {
		charRow := []rune(row)
		for x := range row {
			c := cartsGrid[y][x]
			if c == nil {
				continue
			}
			charRow[x] = rune(c.Dir)
		}
		fmt.Println(string(charRow))
	}
}
