package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Cmd struct {
	direction string
	distance  int
}

func part1(cmds []string) int {
	commands := make([]Cmd, len(cmds))
	for i := 0; i < len(cmds); i++ {
		parsedCmd := strings.Split(cmds[i], " ")
		distance, _ := strconv.Atoi(parsedCmd[1])
		cmd := Cmd{parsedCmd[0], distance}
		commands[i] = cmd
	}
	totalDistance := 0
	totalDepth := 0
	for i := 0; i < len(commands); i++ {
		switch commands[i].direction {
		case "forward":
			totalDistance = totalDistance + commands[i].distance
		case "down":
			totalDepth = totalDepth + commands[i].distance
		case "up":
			totalDepth = totalDepth - commands[i].distance
		}
	}
	return totalDistance * totalDepth
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
	// fmt.Printf("\nPart2: %v", part2(input))
	fmt.Println("\nDone")

}
