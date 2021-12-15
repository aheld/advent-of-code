package main

import (
	"testing"
)

func Test_Part1(t *testing.T) {
	tests := []struct {
		name          string
		expected      int
		expectedLen   int
		expectedStep1 string
		expectedStep2 string
		expected2     int
	}{
		{
			name:          "test_input.txt",
			expected:      1588,
			expectedLen:   3073,
			expectedStep1: "NCNBCHB",
			expectedStep2: "NBCCNBBBCBHCB",
			expected2:     0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, insertions := parseInput(tt.name)
			res := step(start, insertions)
			if res != tt.expectedStep1 {
				t.Errorf("part1() = %v, wanted %v", res, tt.expectedStep1)
			}
			res = step(res, insertions)
			if res != tt.expectedStep2 {
				t.Errorf("part1() = %v, wanted %v", res, tt.expectedStep2)
			}
			res = runSteps(start, insertions, 10)
			if len(res) != tt.expectedLen {
				t.Errorf("part1() = %v, wanted %v", len(res), tt.expectedLen)
			}
			letterCounts := countLetters(res)
			if letterCounts["C"] != 298 && letterCounts["B"] != 174911 {
				t.Errorf("lettercounts off, %v", letterCounts)
			}
			score := part1(tt.name)
			if score != tt.expected {
				t.Errorf("wrong score, got %v, wanted %v", score, tt.expected)
			}
		})
		// t.Run(tt.name, func(t *testing.T) {
		// 	if tt.expected2 == 0 {
		// 		t.SkipNow()
		// 	}
		// 	res := part2(tt.name)
		// 	if res != tt.expected2 {
		// 		t.Errorf("part2() = %v, wanted %v", res, tt.expected2)
		// 	}
		// })
	}
}
