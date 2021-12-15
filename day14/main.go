package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Insertions struct {
	pair   string
	insert string
}

func step(current string, insertions []Insertions) string {
	type InsertsToMake struct {
		pair   string
		insert string
		index  int
	}
	insertionMap := make([]InsertsToMake, 0)
	for i := 0; i < len(current)-1; i++ {
		for _, ins := range insertions {
			if current[i:i+1] == ins.pair[0:1] && current[i+1:i+2] == ins.pair[1:2] {
				insertionMap = append(insertionMap, InsertsToMake{index: i + 1, insert: ins.insert, pair: ins.pair})
			}
		}
	}
	output := current
	for i := len(insertionMap) - 1; i >= 0; i-- {
		ins := insertionMap[i]
		output = output[:ins.index] + ins.insert + output[ins.index:]
	}
	return output
}

func runSteps(start string, insertions []Insertions, count int) string {
	for i := 0; i < count; i++ {
		start = step(start, insertions)
		// fmt.Println(i, len(start))
	}
	return start
}

func countLetters(s string) map[string]int {
	counts := make(map[string]int)
	for _, c := range s {
		counts[string(c)]++
	}
	return counts
}

func parseInput(filename string) (string, []Insertions) {
	insertions := make([]Insertions, 0)
	data := strings.Split(loadFile(filename), "\n")
	start := data[0]
	for _, line := range data[2:] {
		chunks := strings.Split(line, " -> ")
		insertion := Insertions{chunks[0], chunks[1]}
		insertions = append(insertions, insertion)
	}
	return start, insertions
}

func part1(filename string) int {
	return run(filename, 10)
}

func run(filename string, steps int) int {
	start, insertions := parseInput(filename)
	end := runSteps(start, insertions, steps)
	scores := countLetters(end)
	type Score struct {
		letter string
		score  int
	}
	max := Score{letter: "", score: 0}
	for k, v := range scores {
		if v > max.score {
			max.letter = k
			max.score = v
		}
	}
	min := Score{letter: "", score: max.score}
	for k, v := range scores {
		if v < min.score {
			min.letter = k
			min.score = v
		}
	}
	return max.score - min.score
}

func part2(filename string) int {
	counts := doAlgo(filename, 40)
	fmt.Println(counts)
	max := 0
	for _, v := range counts {
		if v > max {
			max = v
		}
	}
	min := max
	for _, v := range counts {
		if v < min {
			min = v
		}
	}
	return max - min
}

// I needed this. I will commit my shame
// https://www.reddit.com/r/adventofcode/comments/rfzq6f/comment/hoktmbx/
func doAlgo(filename string, count int) map[string]int {
	start, insertions := parseInput(filename)
	counts := make(map[string]int)
	for _, c := range start {
		counts[string(c)]++
	}
	pairs := make(map[string]int)
	for i := 0; i < len(start)-1; i++ {
		pairs[start[i:i+2]]++
	}
	for i := 0; i < count; i++ {
		newPairs := make(map[string]int)
		for pair, pcount := range pairs {
			for _, ins := range insertions {
				if pair == ins.pair {
					left := pair[0:1]
					right := pair[1:2]
					newPairs[left+ins.insert] += pcount
					newPairs[ins.insert+right] += pcount
					counts[ins.insert] = counts[ins.insert] + pcount
				}
			}
		}
		pairs = newPairs
	}
	return counts
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
