package main

import (
	"testing"

	"github.com/RyanConnell/adventofcode/2023/pkg/tester"
)

func TestDay01(t *testing.T) {
	tester.TimeAndCheck(1, []tester.TestCase[int]{
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    142,
		},
		{
			Description: "Part 2 (final)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    55701,
		},
	})
}
