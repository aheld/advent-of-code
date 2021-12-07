package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(filename string) int {
	data := strings.Split(loadFile(filename), ",")
	inputArray := make([]int, len(data))
	min := inputArray[0]
	max := inputArray[0]
	for i, v := range data {
		inputArray[i], _ = strconv.Atoi(v)
		if inputArray[i] < min {
			min = inputArray[i]
		}
		if inputArray[i] > max {
			max = inputArray[i]
		}
	}
	minFuel := -1
	for i := min; i <= max; i++ {
		fuel := 0
		for _, v := range inputArray {
			delta := i - v
			fuel = fuel + int(math.Abs(float64(delta)))
		}
		// fmt.Println("test ", i, " fuel ", fuel)
		if minFuel == -1 {
			minFuel = fuel
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}
func part2(filename string) int {
	data := strings.Split(loadFile(filename), ",")
	inputArray := make([]int, len(data))
	min := inputArray[0]
	max := inputArray[0]
	for i, v := range data {
		inputArray[i], _ = strconv.Atoi(v)
		if inputArray[i] < min {
			min = inputArray[i]
		}
		if inputArray[i] > max {
			max = inputArray[i]
		}
	}
	minFuel := -1
	for i := min; i <= max; i++ {
		fuel := 0
		for _, v := range inputArray {
			delta := int(math.Abs(float64(i - v)))
			// triagle rule!
			fuel = fuel + (delta * (delta + 1) / 2)
		}
		// fmt.Println("test ", i, " fuel ", fuel)
		if minFuel == -1 {
			minFuel = fuel
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
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
