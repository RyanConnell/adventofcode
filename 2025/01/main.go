package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/RyanConnell/adventocode/2025/pkg/parser"
)

const MAXSIZE = 100

func main() {
	lines := parser.MustReadFile("input/input.txt")

	solutionPart1, err := solve(lines, true)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solve(lines, false)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
}

/// Part 1 & 2 \\\

func solve(lines []string, exactZero bool) (int, error) {
	var totalZeros int
	position := 50
	for _, line := range lines {
		direction := 1
		if line[0] == 'L' {
			direction = -1
		}

		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}

		var zeros int
		position, zeros = move(position, distance*direction, exactZero)
		totalZeros += zeros
	}
	return totalZeros, nil
}

func move(position, distance int, exactZero bool) (int, int) {
	var zeros int

	// Shift to zero
	if distance < 0 { // Move left
		distanceToZero := position
		if distance+distanceToZero <= 0 {
			distance += distanceToZero
			if !exactZero && distanceToZero != 0 {
				zeros++
			}
			position = 0
		}
	} else { // Move right
		distanceToZero := MAXSIZE - position
		if distance >= distanceToZero {
			distance -= distanceToZero
			if !exactZero && distanceToZero != 0 {
				zeros++
			}
			position = 0
		}
	}

	// Modulus
	if !exactZero {
		zeros += int(math.Abs(float64(distance))) / MAXSIZE
	}
	distance %= MAXSIZE

	// Add remainder
	if distance < 0 {
		position += distance
		if position < 0 {
			position += MAXSIZE
		}
	} else {
		position += distance
		if position >= MAXSIZE {
			position -= MAXSIZE
		}
	}
	if exactZero && position == 0 {
		zeros = 1
	}
	return position, zeros
}
