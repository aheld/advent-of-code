package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	pointMap map[string]Point
}

func (b *Board) AddPoint(p Point) {
	key := fmt.Sprintf("%d,%d", p.x, p.y)
	b.pointMap[key] = p
}

func (b *Board) MovePoint(p Point, x int, y int) {
	oldKey := fmt.Sprintf("%d,%d", p.x, p.y)
	delete(b.pointMap, oldKey)
	key := fmt.Sprintf("%d,%d", x, y)
	b.pointMap[key] = p
}

type Point struct {
	x int
	y int
}

type Fold struct {
	along string
	value int
}

type Input struct {
	points []Point
	folds  []Fold
}

func extractPoint(input string) (Point, error) {
	chunks := strings.Split(input, ",")
	if len(chunks) != 2 {
		return Point{}, errors.New("not a point")
	}
	x, _ := strconv.Atoi(chunks[0])
	y, _ := strconv.Atoi(chunks[1])
	return Point{x, y}, nil
}

func part1(filename string) int {
	input := Input{}
	input.points = []Point{}
	input.folds = []Fold{}
	data := strings.Split(loadFile(filename), "\n")
	for _, line := range data {
		p, err := extractPoint(line)
		if err == nil {
			input.points = append(input.points, p)
			continue
		}
		chunks := strings.Split(line, " ")
		if len(chunks) == 0 {
			continue
		}
		if chunks[0] == "fold" {
			foldChunks := strings.Split(chunks[2], "=")
			value, err := strconv.Atoi(foldChunks[1])
			if err != nil {
				panic(err)
			}
			f := Fold{along: foldChunks[0], value: value}
			input.folds = append(input.folds, f)
			fmt.Println("Fold", f)
		}
	}
	fmt.Println("Points ", len(input.points))
	b := Board{}
	b.pointMap = make(map[string]Point)
	for _, p := range input.points {
		b.AddPoint(p)
	}
	foldOn := input.folds[0]
	if foldOn.along == "y" {
		for _, p := range input.points {
			if p.y > foldOn.value {
				newY := foldOn.value - (p.y - foldOn.value)
				b.MovePoint(p, p.x, newY)
			}
		}
	}
	if foldOn.along == "x" {
		for _, p := range input.points {
			if p.x > foldOn.value {
				newX := foldOn.value - (p.x - foldOn.value)
				b.MovePoint(p, newX, p.y)
			}
		}
	}
	return len(b.pointMap)
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
