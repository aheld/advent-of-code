package main

import (
	"fmt"
	"testing"
)

func Test_Part1(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{
			name:     "small_test_input.txt",
			expected: 6,
		},
		{
			name:     "test_input.txt",
			expected: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getAllPaths(tt.name)
			if res != tt.expected {
				t.Errorf("getAllPaths(%v) = %v, wanted %v", tt.name, res, tt.expected)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			cells, _ := expandMap("test_input.txt")
			expected, _ := loadData("test_expected.txt")
			if len(cells) != len(expected) ||
				len(cells[0]) != len(expected[0]) ||
				cells[12][12].value != expected[12][12].value {
				t.Errorf("expansion failed")
				fmt.Println(cells)
			}
		})
	}
}
