package main

import (
	"fmt"

	"github.com/RyanConnell/adventocode/2025/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/sample.txt")

	solutionPart1, err := solve(lines, false)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solve(lines, true)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
}

/// Part 1 & 2 \\\

func solve(lines []string, continuous bool) (int, error) {
	board := make(map[int]map[int]bool)
	for y, line := range lines {
		board[y] = make(map[int]bool)
		for x, char := range line {
			board[y][x] = char == '@'
		}
	}

	var total int
	for {
		var count int
		for y := range board {
			for x := range board[y] {
				if !board[y][x] {
					continue
				}
				if accessible(board, y, x) {
					count++
					if continuous {
						board[y][x] = false
					}
				}
			}
		}
		total += count
		if !continuous || count == 0 {
			break
		}
	}

	return total, nil
}

func accessible(board map[int]map[int]bool, x, y int) bool {
	var count int
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
				continue
			}
			if board[i][j] {
				count++
			}
		}
	}
	return count < 4
}
