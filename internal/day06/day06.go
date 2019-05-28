package day06

import (
	"fmt"
	"strconv"
	"strings"
)

var infinity = -1

type point struct {
	X int
	Y int
}

// Part1 returns the largest area, defined as the constant Manhattan
// distance-radius around one of the input points, that is not infinite.
//
// Important idea: keep track of the furthest left, right, top, and bottom
// coordinates. If the area for a given point goes beyond one of those
// coordinates, then its area is infinite.
func Part1(input string) (string, error) {
	points, err := parseInput(input)
	if err != nil {
		return "", err
	}
	grid := plotAreas(points)
	areas := findInfiniteAreas(grid)
	for _, row := range grid {
		for _, id := range row {
			if areas[id] == infinity {
				continue
			}
			areas[id]++
		}
	}
	maxArea := 0
	for _, area := range areas {
		if area > maxArea {
			maxArea = area
		}
	}
	return strconv.Itoa(maxArea), nil
}

// Part2 is unimplemented
func Part2(input string) (string, error) {
	return "", nil
}

func parseInput(input string) ([]point, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	points := make([]point, len(lines))
	for i, line := range lines {
		rawCoords := strings.Split(line, ", ")
		if len(rawCoords) != 2 {
			return points, fmt.Errorf("invalid line: %s", line)
		}
		rawX, rawY := rawCoords[0], rawCoords[1]
		x, err := strconv.Atoi(rawX)
		y, err := strconv.Atoi(rawY)
		if err != nil {
			return points, err
		}
		points[i] = point{x, y}
	}
	return points, nil
}

func plotAreas(points []point) [][]int {
	points = normalize(points)
	grid := plotPoints(points)
	given := makeBoolGrid(len(grid), len(grid[0]))
	for _, pt := range points {
		given[pt.Y][pt.X] = true
	}
	for y := range grid {
		for x := range grid[0] {
			pt := point{x, y}
			grid[y][x] = manhattanSearch(pt, grid, given)
		}
	}
	return grid
}

// Marks the area of all IDs with infinite area as -1.
func findInfiniteAreas(grid [][]int) map[int]int {
	infiniteIDs := make(map[int]int)
	maxY, maxX := len(grid)-1, len(grid[0])-1
	// Traverse the left and right edges.
	for y := range grid {
		leftID := grid[y][0]
		rightID := grid[y][maxX]
		infiniteIDs[leftID] = infinity
		infiniteIDs[rightID] = infinity
	}
	// Traverse the top and bottom edges.
	for x := range grid[0] {
		topID := grid[0][x]
		bottomID := grid[maxY][x]
		infiniteIDs[topID] = infinity
		infiniteIDs[bottomID] = infinity
	}
	return infiniteIDs
}

func normalize(points []point) []point {
	normalized := make([]point, len(points))
	_, _, minX, minY := extremeCoords(points)
	for i, pt := range points {
		normalized[i] = point{pt.X - minX, pt.Y - minY}
	}
	return normalized
}

func plotPoints(points []point) [][]int {
	// Assumes points are already normalized, so minX and minY will be 0.
	maxX, maxY, _, _ := extremeCoords(points)
	grid := makeIntGrid(maxY+1, maxX+1)
	for i, pt := range points {
		row, col := pt.Y, pt.X
		grid[row][col] = i
	}
	return grid
}

// extremeCoords returns the extreme x and y values of the input points. Assumes all x and y
// values are non-negative.
func extremeCoords(points []point) (maxX int, maxY int, minX int, minY int) {
	if len(points) == 0 {
		return
	}
	maxX, maxY = points[0].X, points[0].Y
	minX, minY = maxX, maxY
	for _, pt := range points {
		if pt.X > maxX {
			maxX = pt.X
		} else if pt.X < minX {
			minX = pt.X
		}
		if pt.Y > maxY {
			maxY = pt.Y
		} else if pt.Y < minY {
			minY = pt.Y
		}
	}
	return
}

func makeIntGrid(rows int, columns int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, columns)
	}
	return grid
}

func makeBoolGrid(rows int, columns int) [][]bool {
	grid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, columns)
	}
	return grid
}

// Returns the ID of the given point in the grid that has the closest Manhattan
// distance to the source point.
func manhattanSearch(source point, grid [][]int, given [][]bool) int {
	for dist := 0; true; dist++ {
		for xDist := -dist; xDist <= dist; xDist++ {
			yDists := []int{dist - abs(xDist), -dist + abs(xDist)}
			for _, yDist := range yDists {
				x, y := source.X+xDist, source.Y+yDist
				outOfRange := y < 0 || y >= len(given) || x < 0 || x >= len(given[0])
				if !outOfRange && given[y][x] {
					return grid[y][x]
				}
			}
		}
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
