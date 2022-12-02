package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func part1(filename string) int {
	return breedFish(filename, 80)
}

// after 200 days, this uses too much memory ;)
//207
//Alloc = 5339 MiB	TotalAlloc = 30338 MiB	Sys = 9583 MiB	NumGC = 59
func breedFish(filename string, days int) int {
	data := strings.Split(loadFile(filename), ",")
	fish := make([]int, len(data))
	for i, v := range data {
		newfish, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		fish[i] = newfish
	}
	for i := 0; i < days; i++ {
		fmt.Println(i)
		newfish := make([]int, 0)
		for i, v := range fish {
			switch v {
			case 0:
				fish[i] = 6
				newfish = append(newfish, 8)
			default:
				fish[i]--
			}
		}
		fish = append(fish, newfish...)
	}
	return len(fish)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

//Alloc = 5 MiB       TotalAlloc = 26 MiB     Sys = 21 MiB    NumGC = 7
func part2(filename string) int {
	data := strings.Split(loadFile(filename), ",")
	fish := make(map[int]int)
	fishByDay := make([]map[int]int, 257)
	for i := 0; i < 9; i++ {
		fish[i] = 0
	}
	for _, v := range data {
		newfish, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		fish[newfish]++
	}
	fishByDay[0] = fish
	// fmt.Println(fishByDay)
	for i := 0; i < 256; i++ {
		// fmt.Println("\n\nDay: ", i)
		printFishByAge(fishByDay[i])
		PrintMemUsage()
		fishByDay[i+1] = make(map[int]int)
		for age := 1; age < 9; age++ {
			// fmt.Println("Age: ", age, "Fish: ", fishByDay[i][age])
			fishByDay[i+1][age-1] = fishByDay[i][age]
		}
		fishByDay[i+1][8] = fishByDay[i][0]
		fishByDay[i+1][6] = fishByDay[i+1][6] + fishByDay[i][0]
	}
	fmt.Print("\n")
	totalFish := 0
	for _, v := range fishByDay[256] {
		totalFish += v
	}
	return totalFish
}

func printFishByAge(fish map[int]int) {
	fmt.Print("\n")
	for i := 0; i < len(fish); i++ {
		fmt.Print("\t", fish[i])
	}
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
