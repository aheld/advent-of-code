package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type BitCounts struct {
	ones int
	zero int
}

func part1(readings []string) int {
	bitLen := len(readings[0])
	bitCounts := make([]BitCounts, bitLen)
	for i := 0; i < bitLen; i++ {
		bitCounts[i] = BitCounts{ones: 0, zero: 0}
	}

	for i := 0; i < len(readings); i++ {
		reading := []rune(readings[i])
		for j := 0; j < len(reading); j++ {
			switch reading[j] {
			case '1':
				bitCounts[j].ones++
			case '0':
				bitCounts[j].zero++
			default:
				panic("Invalid command")
			}
		}
	}

	gammaArray := make([]string, bitLen)
	epsilonArray := make([]string, bitLen)
	for i := 0; i < len(bitCounts); i++ {
		if bitCounts[i].ones > bitCounts[i].zero {
			gammaArray[i] = "1"
			epsilonArray[i] = "0"
		} else {
			gammaArray[i] = "0"
			epsilonArray[i] = "1"
		}
	}
	gamma := strings.Join(gammaArray, "")
	epsilon := strings.Join(epsilonArray, "")
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epislonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(gammaInt * epislonInt)
}

func part2(cmds []string) int {
	return 0
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
