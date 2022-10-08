package main

import (
	"testing"
)

func Test_Explode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "[[[[[9,8],1],2],3],4]",
			expected: "[[[[0,9],2],3],4]",
		},
		{
			input:    "[7,[6,[5,[4,[3,2]]]]]",
			expected: "[7,[6,[5,[7,0]]]]",
		},
		{
			input:    "[[6,[5,[4,[3,2]]]],1]",
			expected: "[[6,[5,[7,0]]],3]",
		}}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			sn, _ := Parse(tt.input)
			res := sn.Explode()
			if res != true || sn.String() != tt.expected {
				t.Errorf("didn't explode correctly %v, \ngot %v, wanted %v", tt.input, sn.String(), tt.expected)
			}
		})
	}
}
func Test_Add_And_Reduce(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			expected: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			input:    "[7,[6,[5,[4,[3,2]]]]]",
			expected: "[7,[6,[5,[7,0]]]]",
		},
		{
			input:    "[[6,[5,[4,[3,2]]]],1]",
			expected: "[[6,[5,[7,0]]],3]",
		}}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			sn, _ := Parse(tt.input)
			sn.reduce()
			//fmt.Printf("%v", sn)
			if sn.String() != tt.expected {
				t.Errorf("didn't Add correctly %v, \ngot %v, wanted %v", tt.input, sn.String(), tt.expected)
			}
		})
	}
}
func TestAddAll(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected string
	}{
		{
			input: `[1,1]
[2,2]
[3,3]
[4,4]`,
			expected: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			input: `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]`,
			expected: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			input: `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]`,
			expected: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			input: `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`,
			expected: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
	} {
		t.Run(tt.expected, func(t *testing.T) {
			sn := AddAll(tt.input)
			if tt.expected != sn.String() {
				t.Errorf("AddAll failed for %v\n got %v \n wanted %v", tt.input, sn.String(), tt.expected)
			}
		})
	}
}

func Test_Part1(t *testing.T) {
	tests := []struct {
		name      string
		expected  int
		expected2 int
	}{
		{
			name:      "test_input.txt",
			expected:  4140,
			expected2: 3993,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := part1(tt.name)
			if res != tt.expected {
				t.Errorf("part1() = %v, wanted %v", res, tt.expected)
			}
			res = part2(tt.name)
			if res != tt.expected2 {
				t.Errorf("part2() = %v, wanted %v", res, tt.expected2)
			}
		})
	}
}
