package main

import (
	"strings"
	"testing"
)

const input = `199
200
208
210
200
207
240
269
260
263`

type TestCase struct {
	name      string
	depths    []string
	expected  int
	expected2 int
}

func Test_Part1(t *testing.T) {
	tests := []TestCase{
		{
			name:      "example",
			depths:    strings.Split(input, "\n"),
			expected:  7,
			expected2: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := part1(tt.depths)
			if res != tt.expected {
				t.Errorf("part1() = %v, wanted %v", res, tt.expected)
			}
			res = part2(tt.depths)
			if res != tt.expected2 {
				t.Errorf("part2() = %v, wanted %v", res, tt.expected2)
			}
		})
	}
}
