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
			expected:  17,
			expected2: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := part1(tt.name)
			if res != tt.expected {
				t.Errorf("part1() = %v, wanted %v", res, tt.expected)
			}
			res2 := part2(tt.name)
			if res2 != tt.expected2 {
				t.Errorf("part2() = %v, wanted %v", res2, tt.expected2)
			}
		})
	}
}
