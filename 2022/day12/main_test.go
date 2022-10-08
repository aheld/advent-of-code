package main

import (
	"testing"
)

func Test_Part1(t *testing.T) {
	tests := []struct {
		name      string
		expected  int
		expected2 int
	}{
		{
			name:      "test_input.txt",
			expected:  10,
			expected2: 36,
		},
		{
			name:      "test2_input.txt",
			expected:  19,
			expected2: 103,
		},
		{
			name:      "test3_input.txt",
			expected:  226,
			expected2: 3509,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := part1(tt.name)
			if res != tt.expected {
				t.Errorf("%v = %v, wanted %v", tt.name, res, tt.expected)
			}
		})
		t.Run(tt.name+" Part 2", func(t *testing.T) {
			res := part2(tt.name)
			if res != tt.expected2 {
				t.Errorf("part 2 %v = %v, wanted %v", tt.name, res, tt.expected2)
			}
		})
	}
}
