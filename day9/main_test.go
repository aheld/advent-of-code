package main

import (
	"testing"
)

func Test_Part2_fill(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		start    Point
	}{
		{
			name:     "test_input.txt",
			expected: 15,
			start:    Point{3, 3},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Skip()
			board := LoadBoard(tt.name)
			board, basin := fillBasin(board, tt.start)
			res := len(basin.lowPoints)
			if res != tt.expected {
				t.Errorf("Fill = %v, wanted %v", res, tt.expected)
			}
		})
	}
}

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
			t.Skip()
			res := part1(tt.name)
			if res != tt.expected {
				t.Errorf("part1() = %v, wanted %v", res, tt.expected)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			res := part2(tt.name)
			if res != tt.expected2 {
				t.Errorf("part2() = %v, wanted %v", res, tt.expected2)
			}
		})
	}
}
