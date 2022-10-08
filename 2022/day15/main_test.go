package main

import (
	"fmt"
	"testing"
)

func Test_Part1(t *testing.T) {
	t.Skip()
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
	}
}
func Test_Expand(t *testing.T) {
	t.Run("Expansion Test", func(t *testing.T) {
		cells, _ := expandMap("test_input.txt")
		expected, _ := loadData("test_expected.txt")
		success := true
		for i, row := range cells {
			for x, p := range row {
				if p.value != expected[i][x].value {
					success = false
				}
			}
		}
		if !success {
			t.Errorf("expansion failed")
			for i, row := range cells {
				output := fmt.Sprintf("%v:[%v]\t", i, len(row))
				for x, p := range row {
					if p.value != expected[i][x].value {
						output += fmt.Sprintf("%v ", p.value)
					} else {
						output += "- "
					}
				}
				fmt.Println(output)
			}
		}
	})
}

func Test_Part2(t *testing.T) {
	t.Run("Part2 Test", func(t *testing.T) {
		res := part2("test_input.txt")
		expected := 315
		if res != expected {
			t.Errorf("part2() = %v, wanted %v", res, expected)
		}
	})
}
