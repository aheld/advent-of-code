package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var segments = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

type Entry struct {
	allDigits   []string
	easyNumbers []string
	display     []string
	appearances int
}

func part2(filename string) int {

	answer := make([]string, 0)
	for v := range segments {
		answer = append(answer, normalizeSegment(v))
	}
	sort.Strings(answer)

	data := strings.Split(loadFile(filename), "\n")
	entries := buildEntries(data)
	input := []rune("abcdefg")
	translations := getPerms(string(input))
	total := 0
	for _, entry := range entries {
		for i := 0; i < len(translations); i++ {
			translation := translations[i]
			newAlldigits := make([]string, 0)
			for _, origDigit := range entry.allDigits {
				newDigit := translateDigit(origDigit, input, translation)
				newAlldigits = append(newAlldigits, normalizeSegment(newDigit))
			}
			match := stringArrayEqual(newAlldigits, answer)
			if match {
				readout := make([]string, 4)
				for i, v := range entry.display {
					readout[i] = normalizeSegment(translateDigit(v, input, translation))
				}
				display := 0
				for _, v := range readout {
					display = display*10 + segments[v]
				}
				fmt.Println(display)
				total += display
			}
		}
	}
	return total
}

func translateDigit(origDigit string, input []rune, translation string) string {
	newDigit := origDigit
	for x, tx := range input {
		newDigit = strings.Replace(newDigit, string(tx), fmt.Sprint(x), -1)
	}
	for x, tx := range strings.Split(translation, "") {
		newDigit = strings.Replace(newDigit, fmt.Sprint(x), string(tx), -1)
	}
	return newDigit
}

func isElementExist(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func stringArrayEqual(a, b []string) bool {
	for _, v := range a {
		if !isElementExist(b, v) {
			return false
		}
	}
	return true
}

func insertAt(i int, char string, perm string) string {
	start := perm[0:i]
	end := perm[i:]
	return start + char + end
}

// https://gist.github.com/athap/0f2d2b1c84a8c03fadd9
func getPerms(str string) []string {
	// base case, for one char, all perms are [char]
	if len(str) == 1 {
		return []string{str}
	}

	current := str[0:1] // current char
	remStr := str[1:]   // remaining string

	perms := getPerms(remStr) // get perms for remaining string

	allPerms := make([]string, 0) // array to hold all perms of the string based on perms of substring

	// for every perm in the perms of substring
	for _, perm := range perms {
		// add current char at every possible position
		for i := 0; i <= len(perm); i++ {
			newPerm := insertAt(i, current, perm)
			allPerms = append(allPerms, newPerm)
		}
	}
	return allPerms
}

func normalizeSegment(segment string) string {
	segments := strings.Split(segment, "")
	sort.Strings(segments)
	return strings.Join(segments, "")
}

func countMatches(numbers []string, display []string) int {
	count := 0
	for _, n := range numbers {
		for _, d := range display {
			if n == d {
				count++
			}
		}
	}
	return count
}

func buildEntries(lines []string) []Entry {
	entries := make([]Entry, len(lines))
	for x, line := range lines {
		linepart := strings.Split(line, " | ")
		entry := Entry{
			allDigits:   strings.Split(linepart[0], " "),
			easyNumbers: make([]string, 0),
			display:     strings.Split(linepart[1], " "),
			appearances: 0,
		}
		for i, v := range entry.allDigits {
			entry.allDigits[i] = normalizeSegment(v)
		}
		for i, v := range entry.display {
			entry.display[i] = normalizeSegment(v)
		}
		for _, v := range entry.allDigits {
			switch len(v) {
			/*
				0: 6 *
				1: 2 *
				2: 5 *
				3: 5 *
				4: 4 *
				5: 5 *
				6: 6 *
				7: 3 *
				8: 7 *
				9: 6 *
			*/
			case 3: //7
				entry.easyNumbers = append(entry.easyNumbers, v)
			case 2: //1
				entry.easyNumbers = append(entry.easyNumbers, v)
			case 4: //4
				entry.easyNumbers = append(entry.easyNumbers, v)
			case 7: //8
				entry.easyNumbers = append(entry.easyNumbers, v)
			}
		}
		entries[x] = entry
	}
	return entries
}

func part1(filename string) int {
	//be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe\
	data := strings.Split(loadFile(filename), "\n")
	entries := buildEntries(data)
	total := 0
	for _, entry := range entries {
		entry.appearances += countMatches(entry.easyNumbers, entry.display)
		total += entry.appearances
	}
	// fmt.Println(data)
	// fmt.Println(len(data))
	return total
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
