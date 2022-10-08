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
			expected2:     2188189693529,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := part2(tt.name)
			if score != tt.expected2 {
				t.Errorf("wrong score, got %v, wanted %v", score, tt.expected2)
			}
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
			score = part1(tt.name)
			if score != tt.expected {
				t.Errorf("wrong score, got %v, wanted %v", score, tt.expected)
			}
			counts := doAlgo(tt.name, 1)
			if counts["C"] != 2 {
				t.Errorf("wrong score, got %v, wanted %v", counts["C"], counts)
			}
			counts = doAlgo(tt.name, 2)
			if counts["C"] != 4 {
				t.Errorf("wrong score, got %v, wanted 4  %v", counts["C"], counts)
			}
			counts = doAlgo(tt.name, 10)
			if counts["C"] != 298 && counts["B"] != 1749 {
				t.Errorf("wrong score, got %v, wanted 298 %v", counts["C"], counts)
			}

		})
	}
}
