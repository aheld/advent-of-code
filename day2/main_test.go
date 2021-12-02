package main

import (
	"strings"
	"testing"
)

const input = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

type TestCase struct {
	name      string
	cmds      []string
	expected  int
	expected2 int
}

func Test_Part1(t *testing.T) {
	tests := []TestCase{
		{
			name:      "example",
			cmds:      strings.Split(input, "\n"),
			expected:  150,
			expected2: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := part1(tt.cmds)
			if res != tt.expected {
				t.Errorf("part1() = %v, wanted %v", res, tt.expected)
			}
			// res = part2(tt.depths)
			// if res != tt.expected2 {
			// 	t.Errorf("part2() = %v, wanted %v", res, tt.expected2)
			// }
		})
	}
}
