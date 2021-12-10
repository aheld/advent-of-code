package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

func (b *Board) set(p Point, v int) {
	b.cells[p.y][p.x] = v
}

func (b *Board) top() int {
	return len(b.cells) - 1
}

func (b *Board) right() int {
	return len(b.cells[0]) - 1
}

func (b *Board) isLower(a Point, a2 Point) bool {
	return b.get(a) < b.get(a2)
}

func (b *Board) Print() {
	for _, v := range b.cells {
		fmt.Println(v)
	}
}

type Acc struct {
	lowestPoint    Point
	pointsToCheck  []Point
	pointsChecked  []Point
	pointsAccepted []Point
}

type Answers []Acc

func (a Answers) Len() int { return len(a) }

//sort high to low
func (a Answers) Less(i, j int) bool { return len(a[i].pointsAccepted) > len(a[j].pointsAccepted) }
func (a Answers) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Basin struct {
	lowestPoint Point
	lowPoints   []Point
}

func getBasin(board Board, acc Acc) Acc {
	if len(acc.pointsToCheck) == 0 || len(acc.pointsToCheck) > 20 {
		return acc
	}

	p := acc.pointsToCheck[0]
	acc.pointsToCheck = append(acc.pointsToCheck[:0], acc.pointsToCheck[1:]...)
	acc.pointsChecked = append(acc.pointsChecked, p)

	surrounding := getSurroundingPoints(board, p)
	for _, s := range surrounding {
		if board.isLower(p, s) {
			acc.pointsToCheck = append(acc.pointsToCheck, s)
			acc.pointsAccepted = append(acc.pointsAccepted, p)
		}
	}

	acc.pointsChecked = append(acc.pointsChecked, p)

	acc.pointsAccepted = dedupePoints(acc.pointsAccepted)
	acc.pointsChecked = dedupePoints(acc.pointsChecked)

	acc.pointsToCheck = filterPoints(acc.pointsToCheck, acc.pointsAccepted)
	acc.pointsToCheck = dedupePoints(acc.pointsToCheck)

	return getBasin(board, acc)
}

type Stack []Point

func (s Stack) Push(v Point) Stack {
	return append(s, v)
}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Pop() (Stack, Point, error) {
	l := len(s)
	if l == 0 {
		return s, Point{}, errors.New("stack is empty")
	}
	return s[:l-1], s[l-1], nil
}

func fillBasin(board Board, start Point) (Board, Basin) {
	basin := Basin{lowestPoint: start, lowPoints: []Point{start}}

	stack := Stack{start}
	for {
		if len(stack) <= 0 {
			break
		}
		var p Point
		var err error
		stack, p, err = stack.Pop()
		if err != nil {
			panic(err)
		}
		value := board.get(p)
		if value != -1 && value != 9 {
			basin.lowPoints = append(basin.lowPoints, p)
			if value != 9 {
				board.set(p, -1)
			}
			surrounding := getSurroundingPoints(board, p)
			for _, s := range surrounding {
				stack = stack.Push(s)
				// fmt.Print(s, board.get(s), " ")
			}
			// fmt.Print("\n")
		}
		// fmt.Println("Basin ", basin.lowPoints)
	}
	return board, basin
}

func LoadBoard(filename string) Board {
	//9pad
	board := loadBoard(filename)
	nineRow := make([]int, 0)
	for i := 0; i < len(board.cells[0]); i++ {
		nineRow = append(nineRow, 9)
	}
	board.cells = append([][]int{nineRow}, board.cells...)
	board.cells = append(board.cells, nineRow)
	for i := 0; i < len(board.cells); i++ {
		board.cells[i] = append([]int{9}, board.cells[i]...)
		board.cells[i] = append(board.cells[i], 9)
	}
	return board
}

func getAllPointsAtLevel(board Board, level int) []Point {
	res := make([]Point, 0)
	for x := 0; x <= board.right(); x++ {
		for y := 0; y <= board.top(); y++ {
			p := Point{x: x, y: y}
			if board.get(p) == level {
				res = append(res, p)
			}
		}
	}
	return res
}

func part2(filename string) int {
	board := LoadBoard(filename)
	basins := make([]Basin, 0)

	for level := 0; level <= board.top(); level++ {
		points := getAllPointsAtLevel(board, level)
		for _, p := range points {
			_, basin := fillBasin(board, p)
			basins = append(basins, basin)
		}
	}
	totals := make([]int, 0)
	for _, b := range basins {
		b.lowPoints = dedupePoints(b.lowPoints)
		totals = append(totals, len(b.lowPoints))
	}
	sort.Ints(totals)
	total := 1
	for i := len(totals); i > len(totals)-3; i-- {
		total *= totals[i-1]
	}
	return total
}

func dedupePoints(points []Point) []Point {
	check := make(map[Point]int)
	for _, v := range points {
		check[v] = 1
	}
	res := make([]Point, 0)
	for p := range check {
		res = append(res, p)
	}
	return res
}

func filterPoints(points []Point, removeList []Point) []Point {
	ret := make([]Point, 0)
	for _, p := range points {
		for _, s := range removeList {
			if s.x == p.x && s.y == p.y {
				continue
			} else {
				ret = append(ret, p)
			}
		}
	}
	// fmt.Println("Filtered ", points, " ", removeList, " ", ret)
	return ret
}
func part1(filename string) int {
	board := loadBoard(filename)
	// fmt.Println(board)
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

func loadBoard(filename string) Board {
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
	return Board{cells: grid}
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

func getSurroundingPoints(board Board, p Point) []Point {
	surrounding := make([]Point, 0)
	x := p.x
	y := p.y
	if p.y > 0 {
		surrounding = append(surrounding, Point{y: y - 1, x: x})
	}
	if p.x > 0 {
		surrounding = append(surrounding, Point{y: y, x: x - 1})
	}
	if p.y < board.top() {
		surrounding = append(surrounding, Point{y: y + 1, x: x})
	}
	if p.x < board.right() {
		surrounding = append(surrounding, Point{y: y, x: x + 1})
	}
	return surrounding
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
