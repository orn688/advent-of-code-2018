package day11

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Part1 returns the max-sum 3x3 square in the grid.
func Part1(input string) (string, error) {
	serialNumber, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return "", err
	}

	x, y := max3x3Square(serialNumber, 300, 300)
	return fmt.Sprintf("%d,%d", x, y), nil
}

func max3x3Square(serialNumber, width, height int) (maxSquareX, maxSquareY int) {
	grid := makeGrid(width, height)
	setPowerLevels(grid, serialNumber)

	maxSquareSum := math.MinInt32
	for y := 0; y < height-2; y++ {
		for x := 0; x < width-2; x++ {
			squareSum := calculateSquareSum(grid, x, y, 3)
			if squareSum > maxSquareSum {
				maxSquareSum = squareSum
				maxSquareX = x + 1
				maxSquareY = y + 1
			}
		}
	}
	return
}

// Part2 returns the max-sum square of any size in the grid.
func Part2(input string) (string, error) {
	serialNumber, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return "", err
	}

	x, y, size := maxSquare(serialNumber, 300, 300)
	return fmt.Sprintf("%d,%d,%d", x, y, size), nil
}

func maxSquare(serialNumber, width, height int) (maxSquareX, maxSquareY, maxSquareSize int) {
	grid := makeGrid(width, height)
	setPowerLevels(grid, serialNumber)

	maxSquareSum := math.MinInt32
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			maxSize := width - y
			if x > y {
				maxSize = width - x
			}
			for size := 1; size <= maxSize; size++ {
				squareSum := calculateSquareSum(grid, x, y, size)
				if squareSum > maxSquareSum {
					maxSquareSum = squareSum
					maxSquareX = x + 1
					maxSquareY = y + 1
					maxSquareSize = size
				}
			}
		}
	}
	return
}

func makeGrid(width, height int) [][]int {
	grid := make([][]int, height)
	for y := range grid {
		grid[y] = make([]int, width)
	}
	return grid
}

func setPowerLevels(grid [][]int, serialNumber int) {
	for yZeroIndexed := range grid {
		y := yZeroIndexed + 1
		for xZeroIndexed := range grid[0] {
			x := xZeroIndexed + 1
			rackID := x + 10
			power := ((rackID * y) + serialNumber) * rackID
			hundredsDigit := (power / 100) % 10
			grid[yZeroIndexed][xZeroIndexed] = hundredsDigit - 5
		}
	}
}

func calculateSquareSum(grid [][]int, leftX, topY, size int) (sum int) {
	for y := topY; y < topY+size; y++ {
		for x := leftX; x < leftX+size; x++ {
			sum += grid[y][x]
		}
	}
	return
}
