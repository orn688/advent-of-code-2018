package day13

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// A CartDirection corresponds to a cart character from the input.
type CartDirection rune

const (
	up    CartDirection = '^'
	right CartDirection = '>'
	down  CartDirection = 'v'
	left  CartDirection = '<'
)

var cartDirections = [4]CartDirection{up, right, down, left}
var cartRegex = regexp.MustCompile(`[\^>v<]`)

func directionIndex(dir CartDirection) int {
	for i, otherDir := range cartDirections {
		if dir == otherDir {
			return i
		}
	}
	panic("invalid direction")
}

// A TurnDirection represents the direction in which a cart turns at an
// intersection, relative to its previous absolute direction.
//
// The integer value of each TurnDirection corresponds to the number of indexes
// within the cartDirections array by which a cart's direction gets incremented
// (mod 4) when it makes a turn in that direction.
//
// For example, if a cart is traveling downwards (== cartDirections[2]) on the
// track and it makes a left (== 3) turn, its new index in cartDirections will
// be ((2 + 3) % 4) == 1 == right.
type TurnDirection int

const (
	rightTurn TurnDirection = 1
	leftTurn  TurnDirection = 3
	straight  TurnDirection = 4
)

// Cart represents the state of a cart from the input track at some later point
// in time.
type Cart struct {
	Dir          CartDirection
	NextTurn     TurnDirection
	LastMoveTick int
}

func newCart(dir CartDirection) *Cart {
	return &Cart{
		Dir: dir,
		// First turn is always left.
		NextTurn:     leftTurn,
		LastMoveTick: -1,
	}
}

// A CartTrack represents the state of the entire track after some number of
// "ticks".
type CartTrack struct {
	grid  [][]*Cart
	rows  []string
	ticks int
}

func newCartTrack(rows []string) *CartTrack {
	grid := make([][]*Cart, len(rows))
	for y, row := range rows {
		grid[y] = make([]*Cart, len(row))
	}
	for y, row := range rows {
		for x, char := range row {
			if isDirection(char) {
				grid[y][x] = newCart(CartDirection(char))
			}
		}
		rows[y] = cartRegex.ReplaceAllStringFunc(row, func(char string) string {
			if char == string(up) || char == string(down) {
				return "|"
			} else if char == string(left) || char == string(right) {
				return "-"
			}
			// This shouldn't happen.
			panic("invalid cart track regex")
		})
	}
	return &CartTrack{
		grid:  grid,
		rows:  rows,
		ticks: 0,
	}
}

func (track *CartTrack) playCollisions(collisionCallback func(x, y int) bool) {
	for cartsRemaining := track.cartCount(); cartsRemaining > 1; track.ticks++ {
		for y, row := range track.rows {
			for x := range row {
				cart := track.grid[y][x]
				if cart == nil || cart.LastMoveTick == track.ticks {
					// Either there isn't a cart here, or this is the new
					// location of a cart that we already moved during this
					// tick.
					continue
				}
				nextX, nextY := track.nextCartLocation(x, y, cart)
				track.grid[y][x] = nil
				if track.grid[nextY][nextX] == nil {
					track.grid[nextY][nextX] = cart
				} else {
					// The new location is occupied by another cart; collision!
					track.grid[nextY][nextX] = nil
					stopPlaying := collisionCallback(nextX, nextY)
					if stopPlaying {
						return
					}
					cartsRemaining -= 2
				}
				cart.LastMoveTick = track.ticks
			}
		}
	}
}

func (track *CartTrack) nextCartLocation(x, y int, c *Cart) (int, int) {
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
	switch track.rows[y][x] {
	case '+':
		nextDirIndex := (directionIndex(c.Dir) + int(c.NextTurn)) % len(cartDirections)
		c.Dir = cartDirections[nextDirIndex]
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
			log.Panicf("invalid up/down location: %d,%d", x, y)
		}
	case '-':
		if !(c.Dir == left || c.Dir == right) {
			log.Panicf("invalid left/right location: %d,%d", x, y)
		}
	default:
		log.Panicf("invalid location: %d,%d", x, y)
	}
	return x, y
}

func (track *CartTrack) cartCount() int {
	count := 0
	for _, row := range track.grid {
		for _, cart := range row {
			if cart != nil {
				count++
			}
		}
	}
	return count
}

// Useful for debugging.
func (track *CartTrack) print() {
	for y, row := range track.rows {
		charRow := []rune(row)
		for x := range row {
			cart := track.grid[y][x]
			if cart == nil {
				continue
			}
			charRow[x] = rune(cart.Dir)
		}
		fmt.Println(string(charRow))
	}
}

// Part1 returns the coordinates of the first crash between two carts in the
// track.
func Part1(input string) (string, error) {
	track := parseInput(input)
	collisionX, collisionY := -1, -1
	track.playCollisions(func(x, y int) bool {
		collisionX, collisionY = x, y
		return true
	})
	if collisionX == -1 {
		return "", fmt.Errorf("no collisions detected")
	}
	return fmt.Sprintf("%d,%d", collisionX, collisionY), nil
}

// Part2 returns the location of the last remaining cart at the moment after the
// last collision occurs.
func Part2(input string) (string, error) {
	track := parseInput(input)
	track.playCollisions(func(_, _ int) bool {
		return false
	})
	for y, row := range track.rows {
		for x := range row {
			cart := track.grid[y][x]
			if cart != nil {
				return fmt.Sprintf("%d,%d", x, y), nil
			}
		}
	}
	return "", fmt.Errorf("even number of carts, so no last cart")
}

func parseInput(input string) *CartTrack {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")
	track := newCartTrack(rows)
	return track
}

func isDirection(char rune) bool {
	for _, dir := range cartDirections {
		if char == rune(dir) {
			return true
		}
	}
	return false
}

func getNextTurn(dir TurnDirection) TurnDirection {
	turnDirs := []TurnDirection{leftTurn, straight, rightTurn}
	for i, otherDir := range turnDirs {
		if dir == otherDir {
			return turnDirs[(i+1)%len(turnDirs)]
		}
	}
	panic(fmt.Sprintf("unknown turn direction %c", dir))
}
