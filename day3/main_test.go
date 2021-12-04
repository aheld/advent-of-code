package main

import (
	"strings"
	"testing"
)

const input = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

type TestCase struct {
	name      string
	readings  []string
	expected  int
	expected2 int
}

func Test_Part1(t *testing.T) {
	tests := []TestCase{
		{
			name:      "example",
			readings:  strings.Split(input, "\n"),
			expected:  198,
			expected2: 230,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := part1(tt.readings)
			if res != tt.expected {
				t.Errorf("part1() = %v, wanted %v", res, tt.expected)
			}
			res = part2(tt.readings)
			if res != tt.expected2 {
				t.Errorf("part2() = %v, wanted %v", res, tt.expected2)
			}
		})
	}
}
