package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2025/pkg/tester"
)

func TestDay09(t *testing.T) {
	tester.TimeAndCheck(9, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    50,
		},
		{
			Description: "Part 1",
			File:        "input.txt",
			Solution:    solve,
			Expected:    4782896435,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    24,
		},
	})
}
