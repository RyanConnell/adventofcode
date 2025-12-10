package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2025/pkg/tester"
)

func TestDay10(t *testing.T) {
	tester.TimeAndCheck(10, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    7,
		},
		{
			Description: "Part 1",
			File:        "input.txt",
			Solution:    solve,
			Expected:    401,
		},
	})
}
