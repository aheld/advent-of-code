package main

import (
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

func (b *Board) top() int {
	return len(b.cells) - 1
}

func (b *Board) right() int {
	return len(b.cells[0]) - 1
}

func (b *Board) isLower(a Point, a2 Point) bool {
	return b.get(a) < b.get(a2)
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

// totalPoints = totalPoints * len(acc.pointsAccepted)
//var getBasin func(Acc) Acc

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

	// fmt.Println("OUTGOING")
	// fmt.Println("Points to Check", acc.pointsToCheck)
	// fmt.Println("Points Accepted", acc.pointsAccepted)
	// fmt.Println("Points  Checked", acc.pointsChecked)
	return getBasin(board, acc)
}

func part2(filename string) int {
	board := loadBoard(filename)

	type Basin struct {
		lowestPoint Point
		lowPoints   []Point
	}

	basins := make([]Basin, 0)

	for x := 0; x <= board.right(); x++ {
		for y := 0; y <= board.top(); y++ {
			p := Point{x: x, y: y}
			isLowPiont := isLowPiont(board, p)
			if isLowPiont {
				lowPoints := []Point{p}
				basin := Basin{lowestPoint: p, lowPoints: lowPoints}
				basins = append(basins, basin)
			}
		}
	}

	answers := []Acc{}
	total := 1
	for _, b := range basins {
		// fmt.Println("Basin ", b.lowestPoint.x, b.lowestPoint.y)
		pointsToCheck := []Point{b.lowestPoint}
		acc := Acc{
			pointsToCheck:  pointsToCheck,
			pointsAccepted: []Point{b.lowestPoint},
			pointsChecked:  []Point{},
			lowestPoint:    b.lowestPoint,
		}

		acc = getBasin(board, acc)
		// fmt.Println("Basin ", b.lowestPoint.x, b.lowestPoint.y)
		// fmt.Println("Points to Check", acc.pointsToCheck)
		// fmt.Println("Points Accepted", acc.pointsAccepted)
		// fmt.Println("Points  Checked", acc.pointsChecked)
		answers = append(answers, acc)
	}
	sort.Sort(Answers(answers))

	fmt.Println("Sorted")
	for i := 0; i < len(answers); i++ {
		fmt.Println(len(answers[i].pointsAccepted), " ", answers[i].lowestPoint)
		total *= len(answers[i].pointsAccepted)
	}
	return total
	//176715 - too low
	//134514

	/*
			62   {3 77}
		54   {13 95}
		51   {96 5}
		51   {47 95}
	*/
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
