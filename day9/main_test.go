package main

import (
	"testing"
)

func Test_Part1(t *testing.T) {
	tests := []struct {
		name      string
		passedOne bool
		expected  int
		expected2 int
	}{
		{
			name:      "test_input.txt",
			expected:  15,
			expected2: 1134,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := part1(tt.name)
			if res != tt.expected {
				t.Errorf("part1() = %v, wanted %v", res, tt.expected)
			} else {
				tt.passedOne = true
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if !tt.passedOne {
				t.Skip("Skipping part 2")
			}
			res := part2(tt.name)
			if res != tt.expected2 {
				t.Errorf("part2() = %v, wanted %v", res, tt.expected2)
			}
		})
	}
}
