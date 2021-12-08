package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type PotentialSegment struct {
	top         []string
	topLeft     []string
	topRight    []string
	middle      []string
	bottomLeft  []string
	bottomRight []string
	bottom      []string
}

type Entry struct {
	allDigits         []string
	easyNumbers       []string
	display           []string
	appearances       int
	potentialSegments PotentialSegment
}

func normalizeSegment(segment string) string {
	segments := strings.Split(segment, "")
	sort.Strings(segments)
	return strings.Join(segments, "")
}

func compareSegments(a string, b string) bool {
	// if len(a) != len(b) {
	// 	return false
	// }
	// segmentsA := strings.Split(a, "")
	// segmentsB := strings.Split(b, "")
	// sort.Strings(segmentsA)
	// sort.Strings(segmentsB)
	// strA := strings.Join(segmentsA, "")
	// strB := strings.Join(segmentsB, "")

	//return normalizeSegment(a) == normalizeSegment(b)
	return a == b
	// for i, v := range segmentsA {
	// 	if v != segmentsB[i] {
	// 		return false
	// 	}
	// }
	// return true
}

func countMatches(numbers []string, display []string) int {
	count := 0
	for _, n := range numbers {
		for _, d := range display {
			if compareSegments(n, d) {
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
			allDigits:         strings.Split(linepart[0], " "),
			easyNumbers:       make([]string, 0),
			display:           strings.Split(linepart[1], " "),
			appearances:       0,
			potentialSegments: PotentialSegment{},
		}
		for i, v := range entry.allDigits {
			entry.allDigits[i] = normalizeSegment(v)
		}
		for i, v := range entry.display {
			entry.display[i] = normalizeSegment(v)
		}
		//1, 4, 7, or 8
		for _, v := range entry.allDigits {
			switch len(v) {
			case 3: //7
				entry.easyNumbers = append(entry.easyNumbers, v)
				// this is top, topRight, bottomRight
				entry.potentialSegments.top = append(entry.potentialSegments.top, v)
				entry.potentialSegments.topRight = append(entry.potentialSegments.topRight, v)
				entry.potentialSegments.bottomRight = append(entry.potentialSegments.bottomRight, v)
			case 2: //1
				entry.easyNumbers = append(entry.easyNumbers, v)
			case 4: //4
				entry.easyNumbers = append(entry.easyNumbers, v)
			case 7: //8
				entry.easyNumbers = append(entry.easyNumbers, v)
			}
		}
		entry.appearances += countMatches(entry.easyNumbers, entry.display)
		entries[x] = entry
	}
	return entries
}

func part1(filename string) int {
	//be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe\
	data := strings.Split(loadFile(filename), "\n")
	entries := buildEntries(data)
	total := 0
	for _, v := range entries {
		total += v.appearances
	}
	// fmt.Println(data)
	// fmt.Println(len(data))
	return total
}

func part2(filename string) int {
	return -1
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
