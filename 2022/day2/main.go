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

func parseCmd(cmds []string) []Cmd {
	commands := make([]Cmd, len(cmds))
	for i := 0; i < len(cmds); i++ {
		parsedCmd := strings.Split(cmds[i], " ")
		distance, _ := strconv.Atoi(parsedCmd[1])
		cmd := Cmd{parsedCmd[0], distance}
		commands[i] = cmd
	}
	return commands
}
func part1(cmds []string) int {
	commands := parseCmd(cmds)
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

func part2(cmds []string) int {
	commands := parseCmd(cmds)
	aim := 0
	totalDistance := 0
	totalDepth := 0
	for i := 0; i < len(commands); i++ {
		switch commands[i].direction {
		case "forward":
			totalDistance = totalDistance + commands[i].distance
			totalDepth = totalDepth + (commands[i].distance * aim)
		case "down":
			aim = aim + commands[i].distance
		case "up":
			aim = aim - commands[i].distance
		}
	}
	return totalDistance * totalDepth
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
