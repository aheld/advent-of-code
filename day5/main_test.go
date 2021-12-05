package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	name         string
	inputFile    string
	expected     int
	expectedCmds []Cmd
}

func Test_Part1_Parsing(t *testing.T) {
	cmd1 := Cmd{
		start: Point{x: 0, y: 0},
		end:   Point{x: 0, y: 5},
	}
	testcase := TestCase{
		name:         "input basic test",
		inputFile:    "input_basic.txt",
		expected:     0,
		expectedCmds: []Cmd{Cmd{}, Cmd{}},
	}
	t.Run(testcase.name, func(t *testing.T) {
		res := parseCmds(testcase.inputFile, false)
		// fmt.Println(res)
		assert.Equal(t, 2, len(res), "there should be 2 commands")
		assert.Equal(t, res[0], cmd1)
	})
}
func Test_getLine(t *testing.T) {
	tests := []struct {
		name     string
		cmd      Cmd
		expected []Point
	}{{
		name: "make a simple Horizontal line",
		cmd: Cmd{
			start: Point{x: 0, y: 0},
			end:   Point{x: 3, y: 0},
		},
		expected: []Point{{x: 0, y: 0}, {x: 1, y: 0}, {x: 2, y: 0}, {x: 3, y: 0}},
	},
		{name: "make a backwards Horizontal line",
			cmd: Cmd{
				start: Point{x: 10, y: 0},
				end:   Point{x: 8, y: 0},
			},
			expected: []Point{{x: 10, y: 0}, {x: 9, y: 0}, {x: 8, y: 0}},
		},
		{name: "make a diagonal line",
			cmd: Cmd{
				start: Point{x: 1, y: 1},
				end:   Point{x: 3, y: 3},
			},
			expected: []Point{{x: 1, y: 1}, {x: 2, y: 2}, {x: 3, y: 3}},
		},
		{name: "make a diagonal line top right to bottom left ",
			cmd: Cmd{
				start: Point{x: 3, y: 3},
				end:   Point{x: 1, y: 1},
			},
			expected: []Point{{x: 3, y: 3}, {x: 2, y: 2}, {x: 1, y: 1}},
		},
		{name: "make a diagonal line top left to bottom right ",
			cmd: Cmd{
				start: Point{x: 1, y: 3},
				end:   Point{x: 3, y: 1},
			},
			expected: []Point{{x: 1, y: 3}, {x: 2, y: 2}, {x: 3, y: 1}},
		},
		{name: "make a diagonal line bottom right to top left ",
			cmd: Cmd{
				start: Point{x: 3, y: 1},
				end:   Point{x: 1, y: 3},
			},
			expected: []Point{{x: 3, y: 1}, {x: 2, y: 2}, {x: 1, y: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.cmd.getLine()
			// fmt.Println(res)
			assert.Equal(t, tt.expected, res, tt.name+" failed, there should be a line")
		})
	}
}

func Test_Part1_CellCounts(t *testing.T) {
	type TestCase struct {
		name          string
		inputFile     string
		expectedPoint Point
		vents         int
	}
	tests := []TestCase{{
		name:          "input basic Counts",
		inputFile:     "input_basic.txt",
		expectedPoint: Point{x: 0, y: 5},
		vents:         2,
	},
		{
			name:          "input testfile Counts",
			inputFile:     "test_input.txt",
			expectedPoint: Point{x: 0, y: 9},
			vents:         2,
		},
		{
			name:          "input testfile Counts",
			inputFile:     "test_input.txt",
			expectedPoint: Point{x: 4, y: 9},
			vents:         1,
		}}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			cmds := parseCmds(testcase.inputFile, false)
			counts := getCountForCell(cmds, testcase.expectedPoint)
			assert.Equal(t, testcase.vents, counts, "there should be 2 vents at point")
		})
	}
}

func Test_Part1(t *testing.T) {
	type TestCase struct {
		name      string
		inputFile string
		expected  int
	}
	tests := []TestCase{{
		name:      "input testfile Counts",
		inputFile: "test_input.txt",
		expected:  5,
	}}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			cmds := parseCmds(testcase.inputFile, false)
			counts := part1(cmds)
			assert.Equal(t, testcase.expected, counts, "there should be 5 vents")
			// fmt.Println(774-340, 300-734)
		})
	}
}

func Test_Part2(t *testing.T) {
	type TestCase struct {
		name      string
		inputFile string
		expected  int
	}
	tests := []TestCase{{
		name:      "input testfile Counts",
		inputFile: "test_input.txt",
		expected:  12,
	}}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			cmds := parseCmds(testcase.inputFile, true)
			counts := part1(cmds)
			assert.Equal(t, testcase.expected, counts, "there should be 12 vents")
		})
	}
}
