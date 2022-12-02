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
			name:      "x=20..30, y=-10..-5",
			expected:  45,
			expected2: 112,
		},
		{
			name:      "x=235..259, y=-118..-62",
			expected:  45,
			expected2: 112,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected == -1 {
				t.Skip()
			}
			res, count := part1(tt.name)
			if res != tt.expected {
				t.Errorf("part1() = %v, wanted %v", res, tt.expected)
			}
			if count != tt.expected2 {
				t.Errorf("part2() = %v, wanted %v", count, tt.expected2)
			}
		})
	}
}
