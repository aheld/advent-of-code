package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type BitCounts struct {
	ones        int
	zero        int
	mostCommon  rune
	leastCommon rune
}

func inputToRunesArray(readingString []string) [][]rune {
	readings := make([][]rune, 0)
	for i := 0; i < len(readingString); i++ {
		reading := []rune(readingString[i])
		readings = append(readings, reading)
	}
	return readings
}

func getBitCounts(readingString []string) []BitCounts {
	return getBitCountsFromRunes(inputToRunesArray(readingString))
}

func getBitCountsFromRunes(readings [][]rune) []BitCounts {
	bitLen := len(readings[0])
	bitCounts := make([]BitCounts, bitLen)
	for i := 0; i < bitLen; i++ {
		bitCounts[i] = BitCounts{ones: 0, zero: 0}
	}

	for i := 0; i < len(readings); i++ {
		reading := readings[i]
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

	for i := 0; i < len(bitCounts); i++ {
		if bitCounts[i].ones >= bitCounts[i].zero {
			bitCounts[i].mostCommon = '1'
			bitCounts[i].leastCommon = '0'
		} else {
			bitCounts[i].mostCommon = '0'
			bitCounts[i].leastCommon = '1'
		}
	}
	return bitCounts
}

func strArrayToInt(in []string) int {
	strArr := strings.Join(in, "")
	res, _ := strconv.ParseInt(strArr, 2, 64)
	return int(res)
}

func runeArrayToInt(in []rune) int {
	str := string(in)
	res, _ := strconv.ParseInt(str, 2, 64)
	return int(res)
}

func filterByPos(in [][]rune, position int, value rune) (ret [][]rune) {
	for _, r := range in {
		if r[position] == value {
			ret = append(ret, r)
		}
	}
	return ret
}

func getValue(initialReadings [][]rune, bitMatch func(BitCounts) rune) int {
	validReadings := make([][]rune, len(initialReadings))
	copy(validReadings, initialReadings)
	for bitPos := 0; bitPos < len(validReadings[0]); bitPos++ {
		bitCounts := getBitCountsFromRunes(validReadings)
		validReadings = filterByPos(validReadings, bitPos, bitMatch(bitCounts[bitPos]))
		if len(validReadings) == 1 {
			break
		}
	}
	return runeArrayToInt(validReadings[0])
}

func part1(readings []string) int {
	bitCounts := getBitCounts(readings)
	bitLen := len(bitCounts)
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

	return int(strArrayToInt(gammaArray) * strArrayToInt(epsilonArray))
}

func part2(readings []string) int {
	initialReadings := inputToRunesArray(readings)
	ogBitFunc := func(bc BitCounts) rune { return bc.mostCommon }
	oxygenGeneratorRating := getValue(initialReadings, ogBitFunc)

	coBitFunc := func(bc BitCounts) rune { return bc.leastCommon }
	co2ScrubberRating := getValue(initialReadings, coBitFunc)

	return co2ScrubberRating * oxygenGeneratorRating
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
