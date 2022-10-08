package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func part1(depths []string) int {
	last_depth, _ := strconv.Atoi(depths[0])
	ascends := 0

	for i := 1; i < len(depths); i++ {
		var depth, _ = strconv.Atoi(depths[i])
		if depth > last_depth {
			ascends++
		}
		last_depth = depth
	}
	return ascends
}

func part2(depthStrings []string) int {
	depths := make([]int, len(depthStrings))
	for i, _ := range depthStrings {
		depths[i], _ = strconv.Atoi(depthStrings[i])
	}

	windows := make([]int, len(depths)-2)
	for i := 0; i < len(windows); i++ {
		windows[i] = depths[i] + depths[i+1] + depths[i+2]
	}

	ascends := 0
	last_depth := windows[0]
	for i := 1; i < len(windows); i++ {
		var depth = windows[i]
		if depth > last_depth {
			ascends++
		}
		last_depth = depth
	}
	return ascends
}

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input := strings.Split(string(fileBytes), "\n")

	fmt.Printf("\nPart1: %v", part1(input))
	fmt.Printf("\nPart2: %v", part2(input))
	fmt.Println("\nDone")

}
