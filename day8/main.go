package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// type PotentialSegment struct {
// 	top         []string
// 	topLeft     []string
// 	topRight    []string
// 	middle      []string
// 	bottomLeft  []string
// 	bottomRight []string
// 	bottom      []string
// }
type Position string

const (
	Top         Position = "top"
	TopLeft     Position = "topleft"
	TopRight    Position = "topright"
	Middle      Position = "middle"
	BottomLeft  Position = "bottomleft"
	BottomRight Position = "bottomright"
	Bottom      Position = "bottom"
)

type Entry struct {
	allDigits   []string
	easyNumbers []string
	display     []string
	appearances int
	//	potentialSegments PotentialSegment
	potentialSegments map[Position]string
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
			allDigits:         strings.Split(linepart[0], " "),
			easyNumbers:       make([]string, 0),
			display:           strings.Split(linepart[1], " "),
			appearances:       0,
			potentialSegments: make(map[Position]string),
		}
		for i, v := range entry.allDigits {
			entry.allDigits[i] = normalizeSegment(v)
		}
		for i, v := range entry.display {
			entry.display[i] = normalizeSegment(v)
		}
		apppendTo := func(segments []Position, potentialSegments string) {
			for _, v := range segments {
				entry.potentialSegments[v] = entry.potentialSegments[v] + potentialSegments
			}
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
			case 6:
				//0
				apppendTo([]Position{Top, TopRight, BottomRight, TopLeft, BottomLeft, Bottom}, v)
				//6
				apppendTo([]Position{Top, Middle, BottomRight, TopLeft, BottomLeft, Bottom}, v)
				//9
				apppendTo([]Position{Top, Middle, TopRight, TopLeft, BottomRight, Bottom}, v)
			case 3: //7
				entry.easyNumbers = append(entry.easyNumbers, v)
				apppendTo([]Position{Top, TopRight, BottomRight}, v)
			case 2: //1
				entry.easyNumbers = append(entry.easyNumbers, v)
				apppendTo([]Position{TopRight, BottomRight}, v)
			case 4: //4
				entry.easyNumbers = append(entry.easyNumbers, v)
				apppendTo([]Position{TopLeft, TopRight, Middle, BottomRight}, v)
			case 5:
				//5
				apppendTo([]Position{Top, TopLeft, Middle, BottomRight, Bottom}, v)
				//2
				apppendTo([]Position{Top, TopRight, Middle, BottomLeft, Bottom}, v)
				//3
				apppendTo([]Position{Top, TopRight, Middle, BottomRight, Bottom}, v)
			case 7: //8
				entry.easyNumbers = append(entry.easyNumbers, v)
				apppendTo([]Position{Top, TopRight, TopLeft, Middle, BottomRight, BottomLeft, Bottom}, v)
			}

			type void struct{}
			var member void
			for k, segment := range entry.potentialSegments {
				set := make(map[string]void)
				entry.potentialSegments[k] = normalizeSegment(v)
				for _, v := range strings.Split(segment, "") {
					set[v] = member
				}
				combined := make([]string, 0)
				for s := range set {
					combined = append(combined, s)
				}
				sort.Strings(combined)
				entry.potentialSegments[k] = strings.Join(combined, "")
			}

		}
		fmt.Printf("%+v \n", entry.potentialSegments)

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

func part2(filename string) int {
	data := strings.Split(loadFile(filename), "\n")
	entries := buildEntries(data)
	// find top
	for _, entry := range entries {
		for c, v := range entry.potentialSegments {
			fmt.Println(c, v)
		}
	}
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
