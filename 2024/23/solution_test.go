package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2024/pkg/tester"
)

func TestDay23(t *testing.T) {
	tester.TimeAndCheck(23, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    7,
		},
	})
}
