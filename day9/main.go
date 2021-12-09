package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Board struct {
	cells [][]int
}

func (b *Board) get(p Point) int {
	return b.cells[p.y][p.x]
}

func (b *Board) top() int {
	return len(b.cells) - 1
}

func (b *Board) right() int {
	return len(b.cells[0]) - 1
}

func part1(filename string) int {
	data := strings.Split(loadFile(filename), "\n")
	grid := make([][]int, len(data))
	for i := 0; i < len(data); i++ {
		cells := strings.Split(data[i], "")
		grid[i] = make([]int, len(cells))
		for x, cell := range cells {
			v, _ := strconv.Atoi(cell)
			grid[i][x] = v
		}
	}
	board := Board{cells: grid}
	fmt.Println(board)
	risk := 0
	for x := 0; x <= board.right(); x++ {
		for y := 0; y <= board.top(); y++ {
			p := Point{x: x, y: y}
			isLowPiont := isLowPiont(board, p)
			if isLowPiont {
				risk = risk + 1 + board.get(p)
			}
		}
	}
	return risk
}

func isLowPiont(board Board, p Point) bool {
	surrounding := getSurrounding(board, p)
	value := board.get(p)
	for _, v := range surrounding {
		if v > value {
			continue
		}
		return false
	}
	return true
}

func getSurrounding(board Board, p Point) []int {
	surrounding := make([]int, 0)
	x := p.x
	y := p.y
	if p.y > 0 {
		surrounding = append(surrounding, board.get(Point{y: y - 1, x: x}))
	}
	if p.x > 0 {
		surrounding = append(surrounding, board.get(Point{y: y, x: x - 1}))
	}
	if p.y < board.top() {
		surrounding = append(surrounding, board.get(Point{y: y + 1, x: x}))
	}
	if p.x < board.right() {
		surrounding = append(surrounding, board.get(Point{y: y, x: x + 1}))
	}
	return surrounding
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
