package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y  int
	value int
}

func (p *Point) Key() string {
	return fmt.Sprintf("%d_%d", p.x, p.y)
}

func getNeighbors(cells [][]Point, p Point) []Point {
	start := cells[0][0]
	end := cells[len(cells)-1][len(cells[0])-1]
	returnPoints := make([]Point, 0)
	deltas := [][]int{{0, -1}, {1, 0}, {-1, 0}, {0, 1}}
	for _, d := range deltas {
		x := p.x + d[0]
		y := p.y + d[1]
		if (x >= start.x && x <= end.x) && (y >= start.y && y <= end.y) {
			returnPoints = append(returnPoints, cells[x][y])
		}
	}
	return returnPoints
}

func minDistance(dist map[string]int, pathSet map[string]bool) string {
	min := math.MaxInt
	minPoint := ""
	for v := range dist {
		inPath, OK := pathSet[v]
		if OK {
			continue
		}
		if dist[v] < min && !inPath {
			min = dist[v]
			minPoint = v
		}
	}
	return minPoint
}

func getAllPaths(filename string) int {
	cells, visited := loadData(filename)
	totalPoints := len(cells) * len(cells[0])
	dist := make(map[string]int, totalPoints)
	allCells := make(map[string]Point, totalPoints)
	for _, row := range cells {
		for _, p := range row {
			dist[p.Key()] = math.MaxInt
			visited[p.Key()] = false
			allCells[p.Key()] = p
		}
	}
	shortestPath := make(map[string]bool)
	// path := make([]string, 0)
	start := cells[0][0]
	delete(visited, cells[0][0].Key())
	start.value = 0
	cells[0][0] = start
	end := cells[len(cells)-1][len(cells[0])-1]
	// fmt.Println(end)
	visited[start.Key()] = true
	dist[start.Key()] = 0
	//for i := 0; i < len(allCells); i++ {
	count := 0
	for {
		count++
		if count > totalPoints {
			break
		}
		closestKey := minDistance(dist, shortestPath)
		closest := allCells[closestKey]
		shortestPath[closest.Key()] = true
		visited[closest.Key()] = true
		neighbors := getNeighbors(cells, closest)
		for _, n := range neighbors {
			if !shortestPath[n.Key()] &&
				(dist[n.Key()] > (dist[closest.Key()] + n.value)) {
				// fmt.Printf("Updating dist for %v-> %v from %v to %v\n", closest.Key(), n.Key(), dist[n.Key()], dist[closest.Key()]+n.value)
				dist[n.Key()] = dist[closest.Key()] + n.value
			}

		}
	}
	// fmt.Println("Shortest ", shortestPath)
	//fmt.Println("DONE ", dist[end.Key()], dist)

	return dist[end.Key()]
}

func increaseCell(in int) int {
	if in < 9 {
		return in + 1
	}
	return 1
}

func expandMap(filename string) (Cells, Visited) {
	cells, _ := loadData(filename)
	for cnt := 1; cnt <= 5; cnt++ {
		width := len(cells)
		height := len(cells[0])
		fmt.Println("Width ", width, " Height ", height)
		cells = append(cells, make([][]Point, width)...)
		for i := 0; i < width; i++ {
			cells[i+width*cnt] = make([]Point, height)
		}
		for i, row := range cells {
			fmt.Println(i, row)
		}
		for y := 0; y < len(cells[0]); y++ {
			for x := 0; x < len(cells); x++ {
				fmt.Println("Writing ", x, y)
				p := cells[y][x]
				cells[x+width][y] = Point{x: p.x, y: p.y, value: increaseCell(p.value)}
			}
		}
	}
	return cells, make(Visited)
}

type Visited map[string]bool
type Cells [][]Point

func loadData(filename string) (Cells, Visited) {
	data := strings.Split(loadFile(filename), "\n")
	height := len(data)
	width := len(data[0])
	// fmt.Println("Height ", height, " Width ", width)
	cells := make([][]Point, width)
	for i := 0; i < width; i++ {
		cells[i] = make([]Point, height)
	}
	visited := make(Visited)
	for y, line := range data {
		for x, c := range strings.Split(line, "") {
			value, _ := strconv.Atoi(c)
			p := Point{x, y, value}
			cells[x][y] = p
			visited[p.Key()] = false
		}
	}
	return cells, visited
}

func loadFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(dat)
}

func part1(filename string) int {
	return getAllPaths(filename)
}

func part2(filename string) int {
	return -1
}

func main() {
	fmt.Println("Part1 ", part1("input.txt"))
	fmt.Println("Part2 ", part2("input.txt"))
}
