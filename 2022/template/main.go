package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func part1(filename string) int {
	data := strings.Split(loadFile(filename), "\n")
	fmt.Println(data)
	fmt.Println(len(data))
	return -1
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
