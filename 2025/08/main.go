package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/RyanConnell/adventocode/2025/pkg/parser"
)

func main() {
	// Annoyingly this one doesn't have the "number of connections that need to be applied" as
	// part of the input so we need to swap between 10 (sample) and 1000 (input) for testing this...
	lines := parser.MustReadFile("input/input.txt")
	connectionCount := 1000

	solutionPart1, err := solve(lines, connectionCount)
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
	x, y, z         int
	connections     []*Point
	connectedToRoot bool
}

func (p *Point) distanceTo(target *Point) float64 {
	return math.Sqrt(
		math.Pow(float64(p.x-target.x), 2) +
			math.Pow(float64(p.y-target.y), 2) +
			math.Pow(float64(p.z-target.z), 2),
	)
}

func (p *Point) flood() int {
	if p.connectedToRoot {
		return 0
	}
	p.connectedToRoot = true
	floodCount := 1
	for _, connection := range p.connections {
		floodCount += connection.flood()
	}
	return floodCount
}

type Distance struct {
	p1, p2   *Point
	distance float64
}

func parsePoints(lines []string) ([]*Point, []Distance, error) {
	var points []*Point
	for _, line := range lines {
		values := strings.Split(line, ",")
		x, err := strconv.Atoi(values[0])
		if err != nil {
			return nil, nil, err
		}
		y, err := strconv.Atoi(values[1])
		if err != nil {
			return nil, nil, err
		}
		z, err := strconv.Atoi(values[2])
		if err != nil {
			return nil, nil, err
		}
		point := &Point{x: x, y: y, z: z, connectedToRoot: len(points) == 0}
		points = append(points, point)
	}

	var distances []Distance
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			distances = append(distances, Distance{p1, p2, p1.distanceTo(p2)})
		}
	}

	slices.SortFunc(distances, func(a, b Distance) int {
		return cmp.Compare(a.distance, b.distance)
	})

	return points, distances, nil
}

func solve(lines []string, connectionCount int) (int, error) {
	points, distances, err := parsePoints(lines)
	if err != nil {
		return 0, err
	}
	for _, dist := range distances[:connectionCount] {
		dist.p1.connections = append(dist.p1.connections, dist.p2)
		dist.p2.connections = append(dist.p2.connections, dist.p1)
	}

	visited := make(map[*Point]struct{})
	var circuitSizes []int
	for _, p := range points {
		if size, ok := walkConnections(p, visited); ok {
			circuitSizes = append(circuitSizes, size)
		}
	}
	slices.Sort(circuitSizes)

	return circuitSizes[len(circuitSizes)-1] *
		circuitSizes[len(circuitSizes)-2] *
		circuitSizes[len(circuitSizes)-3], nil
}

func walkConnections(start *Point, visited map[*Point]struct{}) (int, bool) {
	if _, ok := visited[start]; ok {
		return 0, false
	}
	visited[start] = struct{}{}
	size := 1
	for _, connection := range start.connections {
		val, ok := walkConnections(connection, visited)
		if ok {
			size += val
		}
	}
	return size, true
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	points, distances, err := parsePoints(lines)
	if err != nil {
		return 0, err
	}

	rootCount := len(points)
	for _, dist := range distances {
		dist.p1.connections = append(dist.p1.connections, dist.p2)
		dist.p2.connections = append(dist.p2.connections, dist.p1)
		if dist.p1.connectedToRoot {
			rootCount -= dist.p2.flood()
		}
		if dist.p2.connectedToRoot {
			rootCount -= dist.p1.flood()
		}
		if rootCount == 1 {
			return dist.p1.x * dist.p2.x, nil
		}
	}

	return 0, fmt.Errorf("nothing found")
}
