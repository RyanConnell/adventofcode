package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RyanConnell/adventocode/2025/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/input.txt")

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

type Machine struct {
	endState uint
	buttons  []uint
}

func parseInput(lines []string) ([]*Machine, error) {
	var machines []*Machine
	for _, line := range lines {
		components := strings.Split(line, " ")

		// Create []bool from end state
		var endState uint
		for i, char := range components[0][1 : len(components[0])-1] {
			if i != 0 {
				endState = endState << 1
			}
			if char == '#' {
				endState |= 1
			}
		}

		// Create []bool for each button
		positionCount := len(components[0]) - 2
		buttons := make([]uint, len(components)-2)
		for i, buttonStr := range components[1 : len(components)-1] {
			numStrs := strings.Split(buttonStr[1:len(buttonStr)-1], ",")
			var button uint
			for _, numStr := range numStrs {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					return nil, err
				}
				button |= 1 << (positionCount - num - 1)
			}
			buttons[i] = button
		}

		machine := &Machine{
			endState: endState,
			buttons:  buttons,
		}
		machines = append(machines, machine)
	}
	return machines, nil

}

func solve(lines []string) (int, error) {
	machines, err := parseInput(lines)
	if err != nil {
		return 0, err
	}
	var total int
	for _, machine := range machines {
		smallest := shortestDistance(machine.endState, machine.buttons)
		total += smallest
	}
	return total, nil
}

func shortestDistance(pattern uint, buttons []uint) int {
	distanceCache := make(map[uint]int)
	for _, button := range buttons {
		distanceCache[button] = 1
	}
	for {
		if distance, ok := distanceCache[pattern]; ok {
			return distance
		}

		for _, buttonA := range buttons {
			for buttonB, distance := range distanceCache {
				next := buttonA ^ buttonB
				otherDistance, ok := distanceCache[next]
				if !ok || (ok && otherDistance > distance+1) {
					distanceCache[buttonA^buttonB] = distance + 1
				}
			}
		}
	}
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	return 0, nil
}
