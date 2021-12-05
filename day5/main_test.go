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
		res := parseCmds(testcase.inputFile)
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
		}}
	for _, tt := range tests {
		t.Run("make a line", func(t *testing.T) {
			res := tt.cmd.getLine()
			// fmt.Println(res)
			assert.Equal(t, tt.expected, res, "there should be a line")
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
			cmds := parseCmds(testcase.inputFile)
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
			cmds := parseCmds(testcase.inputFile)
			counts := part1(cmds)
			assert.Equal(t, testcase.expected, counts, "there should be 5 vents")
		})
	}
}
