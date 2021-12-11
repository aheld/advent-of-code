package main

import (
	"testing"
)

var corrupted_input = map[string]string{
	"{([(<{}[<>[]}>{[]{[(<()>": "}", //, - Expected ], but found } instead.
	"[[<[([]))<([[{}[[()]]]":   ")", // - Expected ], but found ) instead.
	"[{[{({}]{}}([{[{{{}}([]":  "]", // - Expected ), but found ] instead.
	"[<(<(<(<{}))><([]([]()":   ")", // - Expected >, but found ) instead.
	"<{([([[(<>()){}]>(<<{{":   ">", // - Expected ], but found > instead.`
}

var autocomplete_input = map[string]string{
	"[({(<(())[]>[[{[]{<()<>>": "}}]])})]",
	"[(()[<>])]({[<{<<[]>>(":   ")}>]})",
	"(((({<>}<{<{<>}{[]{[]{}":  "}}>}>))))",
	"{<[[]]>}<{[{[{[]{()[[[]":  "]]}}]}]}>",
	"<{([{{}}[<[[[<>{}]]]>[]]": "])}>",
}

func Test_Parser(t *testing.T) {
	t.Run("Corrupted input", func(t *testing.T) {
		for input, expected := range corrupted_input {
			illegals := findIllegals(input)
			if illegals[0] != expected {
				t.Errorf("parser(%s) = %v, wanted %v", input, illegals[0], expected)
			}
		}
	})
}
func Test_Autocomplete(t *testing.T) {
	t.Run("Autocomplete input", func(t *testing.T) {
		for input, expected := range autocomplete_input {
			completion := findAutocomplete(input)
			if completion != expected {
				t.Errorf("parser(%s) = %v, wanted %v", input, completion, expected)
			}
		}
	})
}
func Test_Part1(t *testing.T) {
	tests := []struct {
		name      string
		expected  int
		expected2 int
	}{
		{
			name:      "test_input.txt",
			expected:  26397,
			expected2: 288957,
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
