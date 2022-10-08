package main

import (
	"fmt"
	"testing"
)

func Test_Step(t *testing.T) {
	tests := []struct {
		input   string
		step1   string
		step2   string
		flashes int
	}{{
		input: `11111
19991
19191
19991
11111`,
		step1: `34543
40004
50005
40004
34543`,
		step2: `45654
51115
61116
51115
45654`, flashes: 9}}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %v ", i), func(t *testing.T) {
			board := MakeBoard(tt.input)
			// board.print()
			board.step()
			expected := MakeBoard(tt.step1)
			if !board.equal(&expected) {
				t.Errorf("step1 = \n%v, wanted\n%v", board.toString(), tt.step1)
			}
			board = MakeBoard(tt.input)
			board.runSteps(2)
			expected = MakeBoard(tt.step2)
			if !board.equal(&expected) {
				t.Errorf("step1 = \n%v, wanted\n%v", board.toString(), tt.step2)
			}
			if board.flashes != tt.flashes {
				t.Errorf("flashes wrong, got %v wanted %v", board.flashes, tt.flashes)
			}
		})
	}
}

func (b *Board) toString() string {
	lines := b.grid
	output := ""
	for _, line := range lines {
		for _, c := range line {
			output = output + fmt.Sprintf("%d", c.value)
		}
		output = output + "\n"
	}
	return output
}

func Test_Part1(t *testing.T) {
	tests := []struct {
		name      string
		expected  int
		expected2 int
	}{
		{
			name:      "test_input.txt",
			expected:  1656,
			expected2: 195,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part1Passed := true
			res := part1(tt.name)
			if res != tt.expected {
				part1Passed = false
				t.Errorf("part1() = %v, wanted %v", res, tt.expected)
			}
			if !part1Passed {
				t.Skip()
			}
			res = part2(tt.name)
			if res != tt.expected2 {
				t.Errorf("part2() = %v, wanted %v", res, tt.expected2)
			}
		})
	}
}
