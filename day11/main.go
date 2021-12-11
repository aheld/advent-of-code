package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	x, y    int
	value   int
	flashed bool
}

type Board struct {
	grid    [][]Cell
	flashes int
}

func (b *Board) Set(c Cell) {
	b.grid[c.y][c.x] = c
}

func (b *Board) Get(x, y int) Cell {
	return b.grid[y][x]
}

func (b *Board) print() {
	fmt.Print("\n Flashes: ", b.flashes, "\n")
	lines := b.grid
	for _, line := range lines {
		for _, c := range line {
			fmt.Printf("%d\t", c.value)
		}
		fmt.Print("\n")
	}
}

func (b *Board) equal(a *Board) bool {
	for _, line := range b.grid {
		for _, c := range line {
			if c.value != b.Get(c.x, c.y).value {
				fmt.Println("Not equal", c, b.Get(c.x, c.y))
				return false
			}
		}
	}
	return true
}

func processCells(b *Board, f func(Cell) Cell) {
	lines := b.grid
	for _, line := range lines {
		for _, c := range line {
			cell := f(c)
			b.Set(cell)
		}
	}
}

func (b *Board) getNeighbours(c Cell) []Cell {
	neighbourCells := make([]Cell, 0)
	delta := []int{-1, 0, 1}

	for _, dy := range delta {
		for _, dx := range delta {
			if dx == 0 && dy == 0 {
				continue
			}
			x := c.x + dx
			y := c.y + dy
			if x < 0 || y < 0 || x >= len(b.grid[0]) || y >= len(b.grid) {
				continue
			}
			neighbourCells = append(neighbourCells, b.grid[y][x])
		}
	}
	return neighbourCells
}

func (b *Board) cellCount() int {
	return len(b.grid) * len(b.grid[0])
}

func (b *Board) flashAndIncrement(c Cell) {
	neighbours := b.getNeighbours(c)
	for _, n := range neighbours {
		n.value = n.value + 1
		b.Set(n)
	}
}

func (b *Board) runSteps(steps int) {
	for i := 0; i < steps; i++ {
		b.step()
	}
}

func (b *Board) step() {
	incrementAll := func(c Cell) Cell {
		return Cell{c.x, c.y, c.value + 1, c.flashed}
	}
	needToRunFlashProcessing := true
	findCellsToFlash := func(c Cell) Cell {
		if c.value > 9 && !c.flashed {
			needToRunFlashProcessing = true
			b.flashAndIncrement(c)
			b.flashes++
			c.flashed = true
		}
		return c
	}
	processCells(b, incrementAll)

	resetFlash := func(c Cell) Cell {
		if c.value > 9 {
			c.flashed = false
			c.value = 0
		}
		return c
	}

	for i := 0; needToRunFlashProcessing; i++ {
		if i > 100 {
			panic("Infinite loop")
		}
		needToRunFlashProcessing = false
		processCells(b, findCellsToFlash)
		//		b.print()
	}

	processCells(b, resetFlash)
}

func MakeBoard(input string) Board {
	lines := strings.Split(input, "\n")
	b := Board{flashes: 0}
	b.grid = make([][]Cell, len(lines))
	for y, line := range lines {
		b.grid[y] = make([]Cell, len(line))
		for x, c := range line {
			value, _ := strconv.Atoi(string(c))
			b.grid[y][x] = Cell{x: x, y: y, flashed: false, value: value}
		}
	}
	b.print()
	return b
}

func part1(filename string) int {
	data := loadFile(filename)
	board := MakeBoard(data)
	board.runSteps(100)
	return board.flashes
}

func part2(filename string) int {
	data := loadFile(filename)
	board := MakeBoard(data)
	totalOctopi := board.cellCount()
	lastFlashed := 0
	steps := 2000
	for i := 1; i < steps; i++ {
		prevFlashes := board.flashes
		board.step()
		newFlashes := board.flashes
		lastFlashed = newFlashes - prevFlashes
		fmt.Printf("Step %d, Last Flashed %d of %v\n", i, lastFlashed, totalOctopi)
		if totalOctopi == lastFlashed {
			return i
		}
	}
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
