package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2025/pkg/tester"
)

func TestDay01(t *testing.T) {
	tester.TimeAndCheck(1, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, true) },
			Expected:    3,
		},
		{
			Description: "Part 1",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, true) },
			Expected:    1141,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    6,
		},
		{
			Description: "Part 2 (sample2)",
			File:        "sample2.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    3,
		},
		{
			Description: "Part 2 (sample2)",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    6634,
		},
	})
}

func TestMove(t *testing.T) {
	testCases := map[string]struct {
		position, distance int
		exactZero          bool
		expectedPosition   int
		expectedZeros      int
	}{
		"5R5 exact":  {5, 5, true, 10, 0},
		"5R5":        {5, 5, false, 10, 0},
		"5L5 exact":  {5, -5, true, 0, 1},
		"5L5":        {5, -5, false, 0, 1},
		"1L2 exact":  {1, -2, true, 99, 0},
		"1L2":        {1, -2, false, 99, 1},
		"99R2 exact": {99, 2, true, 1, 0},
		"99R2":       {99, 2, false, 1, 1},
		"1L1 exact":  {1, -1, true, 0, 1},
		"1L1":        {1, -1, false, 0, 1},
		"99R1 exact": {99, 1, true, 0, 1},
		"99R1":       {99, 1, false, 0, 1},
		"99R300":     {99, 300, false, 99, 3},
		"99R301":     {99, 301, false, 0, 4},
		"99L398":     {99, -398, false, 1, 3},
		"99L399":     {99, -399, false, 0, 4},
		"82L30":      {82, -30, false, 52, 0},
		"0L100":      {0, -100, false, 0, 1},
		"0R100":      {0, 100, false, 0, 1},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			position, zeros := move(tc.position, tc.distance, tc.exactZero)
			if position != tc.expectedPosition {
				t.Errorf("position != expected: %d != %d", position, tc.expectedPosition)
			}
			if zeros != tc.expectedZeros {
				t.Errorf("zeros != expected: %d != %d", zeros, tc.expectedZeros)
			}
		})
	}
}
