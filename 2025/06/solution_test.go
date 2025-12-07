package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2025/pkg/tester"
)

func TestDay06(t *testing.T) {
	tester.TimeAndCheck(6, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solvePart1,
			Expected:    4277556,
		},
		{
			Description: "Part 1",
			File:        "input.txt",
			Solution:    solvePart1,
			Expected:    5784380717354,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    3263827,
		},
		{
			Description: "Part 2",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    7996218225744,
		},
	})
}
