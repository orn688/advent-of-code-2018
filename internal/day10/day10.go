package day10

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/orn688/advent-of-code-2018/internal/util"
)

var lineRegex = regexp.MustCompile(`^position=<( ?)(?P<X>-?\d+), ( ?)(?P<Y>-?\d+)> ` +
	`velocity=<( ?)(?P<Vx>-?\d+), ( ?)(?P<Vy>-?\d+)>$`)

type point struct {
	X  int
	Y  int
	vx int
	vy int
}

func (p *point) step() {
	p.X += p.vx
	p.Y += p.vy
}

func (p *point) stepBack() {
	p.X -= p.vx
	p.Y -= p.vy
}

type pointgrid struct {
	points                 []*point
	minX, maxX, minY, maxY int
}

func newGrid(points []*point) *pointgrid {
	grid := &pointgrid{points, 0, 0, 0, 0}
	grid._refreshBounds()
	return grid
}

func (grid *pointgrid) getMessage(maxSteps int) (string, int) {
	area := grid.bboxArea()
	for stepCount := 0; stepCount < maxSteps; stepCount++ {
		grid.step()
		// We assume the minimum bounding box area occurs when the points
		// converge to form the word. Therefore, as soon as the area starts
		// increasing, we have passed the instant where the message is shown
		// and must backtrack by one step.
		if grid.bboxArea() > area {
			grid.stepBack()
			return grid.asString(), stepCount
		}
		area = grid.bboxArea()
	}
	return "", -1
}

func (grid *pointgrid) asString() string {
	arr := make([][]bool, grid.height())
	for y := range arr {
		arr[y] = make([]bool, grid.width())
	}
	for _, pt := range grid.points {
		arr[pt.Y-grid.minY][pt.X-grid.minX] = true
	}

	outputRows := make([]string, grid.height())
	for i, row := range arr {
		chars := make([]string, len(arr[0]))
		for x, isSet := range row {
			if isSet {
				chars[x] = "#"
			} else {
				chars[x] = " "
			}
		}
		outputRows[i] = strings.Join(chars, "")
	}
	return strings.Join(outputRows, "\n")
}

func (grid *pointgrid) step() {
	for _, pt := range grid.points {
		pt.step()
	}
	grid._refreshBounds()
}

func (grid *pointgrid) stepBack() {
	for _, pt := range grid.points {
		pt.stepBack()
	}
	grid._refreshBounds()
}

func (grid *pointgrid) width() int {
	return grid.maxX - grid.minX + 1
}

func (grid *pointgrid) height() int {
	return grid.maxY - grid.minY + 1
}

func (grid *pointgrid) bboxArea() int {
	return grid.width() * grid.height()
}

func (grid *pointgrid) _refreshBounds() {
	grid.minX, grid.maxX = grid.points[0].X, grid.points[0].X
	grid.minY, grid.maxY = grid.points[0].Y, grid.points[0].Y
	for _, pt := range grid.points {
		if pt.X > grid.maxX {
			grid.maxX = pt.X
		} else if pt.X < grid.minX {
			grid.minX = pt.X
		}
		if pt.Y > grid.maxY {
			grid.maxY = pt.Y
		} else if pt.Y < grid.minY {
			grid.minY = pt.Y
		}
	}
}

// Part1 returns the message that the points form at the moment that they
// converge.
func Part1(input string) (string, error) {
	points, err := parseInput(input)
	if err != nil {
		return "", err
	}

	grid := newGrid(points)
	msg, _ := grid.getMessage(1e9)
	return msg, nil
}

// Part2 returns the number of steps it takes for the points to converge.
func Part2(input string) (string, error) {
	points, err := parseInput(input)
	if err != nil {
		return "", err
	}

	grid := newGrid(points)
	_, steps := grid.getMessage(1e9)
	return strconv.Itoa(steps), nil
}

func parseInput(input string) ([]*point, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	points := make([]*point, len(lines))
	for i, line := range lines {
		req, err := parseLine(line)
		if err != nil {
			return points, err
		}
		points[i] = req
	}
	return points, nil
}

// Convert a line to a requirement, assuming each step name is a length-1
// string of capital ASCII letters.
func parseLine(line string) (pt *point, err error) {
	groups, err := util.CaptureRegexGroups(lineRegex, line)
	if err != nil {
		return pt, err
	}
	x, err := strconv.Atoi(groups["X"])
	y, err := strconv.Atoi(groups["Y"])
	vx, err := strconv.Atoi(groups["Vx"])
	vy, err := strconv.Atoi(groups["Vy"])
	if err != nil {
		return
	}
	pt = &point{
		X:  x,
		Y:  y,
		vx: vx,
		vy: vy,
	}
	return
}
