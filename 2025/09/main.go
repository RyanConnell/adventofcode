package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/RyanConnell/adventocode/2025/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/sample.txt")

	solutionPart1, err := solve(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
}

/// Part 1 \\\

type Point struct {
	x, y int
}

func (p *Point) rectangleSize(target *Point) int {
	return int(1+math.Abs(float64(p.x-target.x))) * int(1+math.Abs(float64(p.y-target.y)))
}

func solve(lines []string) (int, error) {
	var points []*Point
	for _, line := range lines {
		numStrs := strings.Split(line, ",")
		x, err := strconv.Atoi(numStrs[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(numStrs[1])
		if err != nil {
			return 0, err
		}
		points = append(points, &Point{x: x, y: y})
	}

	var largest int
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			rectSize := p1.rectangleSize(p2)
			if rectSize > largest {
				largest = rectSize
			}
		}
	}
	return largest, nil
}

/// Part 2 \\\

type Line struct {
	start, end *Point
}

const (
	VERTICAL = iota
	HORIZONTAL
)

func minmax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func (l *Line) intersects(o *Line) bool {
	if l.start == o.start || l.start == o.end || l.end == o.start || l.end == o.end {
		// Don't count intersection if we have a point on the line.
		return false
	}
	switch l.direction() {
	case VERTICAL:
		if o.direction() == VERTICAL && o.start.x != l.start.x {
			// If both are vertical and not on same X they can't intersect
			return false
		}
		minX, maxX := minmax(o.start.x, o.end.x)
		if l.start.x >= minX && l.start.y <= maxX {
			minY, maxY := minmax(l.start.y, l.end.y)
			return o.start.y >= minY && o.start.y <= maxY
		}
	case HORIZONTAL:
		if o.direction() == HORIZONTAL && o.start.y != l.start.y {
			// If both are horizontal and not on same Y they can't intersect
			return false
		}
		minX, maxX := minmax(l.start.x, l.end.x)
		if o.start.x >= minX && o.start.y <= maxX {
			minY, maxY := minmax(o.start.y, o.end.y)
			return l.start.y >= minY && o.start.y <= maxY
		}
	}
	return false
}

func (l *Line) direction() int {
	if l.start.x == l.end.x {
		return VERTICAL
	}
	return HORIZONTAL
}

func solvePart2(input []string) (int, error) {
	var points []*Point
	for _, line := range input {
		numStrs := strings.Split(line, ",")
		x, err := strconv.Atoi(numStrs[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(numStrs[1])
		if err != nil {
			return 0, err
		}
		points = append(points, &Point{x: x, y: y})
	}

	var lines []*Line
	last := points[len(points)-1]
	for _, p := range points {
		lines = append(lines, &Line{last, p})
		last = p
	}

	var largest int
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			if !validRectangle(p1, p2, lines) {
				continue
			}
			rectSize := p1.rectangleSize(p2)
			if rectSize > largest {
				largest = rectSize
			}
		}
	}
	return largest, nil
}

func validRectangle(p1, p2 *Point, lines []*Line) bool {
	// TODO: Count how many lines we intersect per point
	fmt.Println("ValidRectangle: ", p1, p2)

	points := []*Point{&Point{p2.x, p1.y}, &Point{p1.x, p2.y}}
	for _, point := range points {
		testLines := []*Line{
			{point, &Point{point.x, 0}},
			{point, &Point{point.x, 100}},
			{point, &Point{0, point.y}},
			{point, &Point{100, point.y}},
		}
		for _, testLine := range testLines {
			var intersectionCounter int
			for _, line := range lines {
				if testLine.intersects(line) {
					intersectionCounter++
				}
			}
			fmt.Printf("\t%v => %d intersections\n", testLine, intersectionCounter)
		}
	}

	return false
}
