package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var delimiters = []string{"<>", "{}", "[]", "()"}

func findIllegals(input string) []string {
	startLen := len(input)
	endLen := 0
	for startLen != endLen {
		startLen = len(input)
		for _, delim := range delimiters {
			input = strings.ReplaceAll(input, delim, "")
		}
		endLen = len(input)
	}
	firstBad := make([]string, 0)
	for _, s := range input {
		counts_open := make(map[string]int)
		counts_closed := make(map[string]int)
		counts := make(map[string]int)
		for _, d := range delimiters {
			if s == rune(d[0]) {
				counts_open[d]++
				counts[d]++
			} else if s == rune(d[1]) {
				counts_closed[d]++
				counts[d]--
				if counts[d] < 0 {
					firstBad = append(firstBad, string(s))
				}
				counts_closed[d]++
			}
		}
	}

	// fmt.Println(input, "   ", firstBad[0])
	return firstBad
}

func findAutocomplete(input string) string {
	startLen := len(input)
	endLen := 0
	for startLen != endLen {
		startLen = len(input)
		for _, delim := range delimiters {
			input = strings.ReplaceAll(input, delim, "")
		}
		endLen = len(input)
	}
	rns := []rune(input)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	dstring := strings.Join(delimiters, "")
	for i := 0; i < len(rns); i++ {
		delimStart := strings.Index(dstring, string(rns[i]))
		rns[i] = rune(dstring[delimStart+1])
	}
	input = string(rns)
	return input
}

func part1(filename string) int {
	data := strings.Split(loadFile(filename), "\n")
	scoring := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	total := 0
	for _, line := range data {
		res := findIllegals(line)
		if len(res) > 0 {
			total += scoring[res[0]]
		}
	}
	return total
}

func getAcScore(input string) int {
	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	score := 0
	for _, c := range input {
		score *= 5
		score += points[string(c)]
	}
	return score
}

func part2(filename string) int {
	data := strings.Split(loadFile(filename), "\n")
	input := make([]string, 0)
	for _, line := range data {
		res := findIllegals(line)
		if len(res) == 0 {
			input = append(input, line)
		}
	}
	scores := []int{}
	for _, line := range input {
		ac := findAutocomplete(line)
		score := getAcScore(ac)
		scores = append(scores, score)
	}
	sort.Ints(scores)
	mid := int(len(scores) / 2)
	return scores[mid]
}

func loadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(data)
}

func main() {
	fmt.Println("Part1 ", part1("input.txt"))
	fmt.Println("Part2 ", part2("input.txt"))
}
