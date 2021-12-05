package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(cmds []Cmd) int {
	total := 0
	board := buildBoard(cmds)
	for _, val := range board {
		if val > 1 {
			total++
		}
	}
	return total
}

func part2() int {
	return -1
}

type Point struct {
	x int
	y int
}

type Cmd struct {
	start Point
	end   Point
}

func (c *Cmd) getLine() []Point {
	if c.isHorizontal() {
		if c.end.x >= c.start.x {
			f := func(start Point) Point {
				return Point{x: start.x + 1, y: start.y}
			}
			length := c.end.x - c.start.x
			return paintLine(c.start, length, f)
		}
		if c.end.x < c.start.x {
			f := func(start Point) Point {
				return Point{x: start.x - 1, y: start.y}
			}
			length := c.start.x - c.end.x
			return paintLine(c.start, length, f)
		}
	}
	if c.isVertical() {
		if c.end.y >= c.start.y {
			f := func(start Point) Point {
				return Point{x: start.x, y: start.y + 1}
			}
			length := c.end.y - c.start.y
			return paintLine(c.start, length, f)
		}
		if c.end.y < c.start.y {
			f := func(start Point) Point {
				return Point{x: start.x, y: start.y - 1}
			}
			length := c.start.y - c.end.y
			return paintLine(c.start, length, f)
		}
	}
	return []Point{}
}

func paintLine(start Point, length int, next func(Point) Point) []Point {
	line := make([]Point, 0)
	line = append(line, start)
	for i := 0; i < length; i++ {
		p := next(line[len(line)-1])
		line = append(line, p)
	}
	return line
}

func buildBoard(cmds []Cmd) map[Point]int {
	board := make(map[Point]int)
	for _, c := range cmds {
		for _, p := range c.getLine() {
			_, exists := board[p]
			if exists {
				board[p] = board[p] + 1
			} else {
				board[p] = 1
			}
		}
	}
	return board
}

func getCountForCell(cmds []Cmd, pos Point) int {
	board := buildBoard(cmds)
	return board[pos]
}

func MakeCmd(input string) Cmd {
	reCmd := regexp.MustCompile(`^(\d*),(\d*) -> (\d*),(\d*)$`)
	match := reCmd.FindAllStringSubmatch(input, -1)
	cmd := Cmd{}
	cmd.start.x, _ = strconv.Atoi(match[0][1])
	cmd.start.y, _ = strconv.Atoi(match[0][2])
	cmd.end.x, _ = strconv.Atoi(match[0][3])
	cmd.end.y, _ = strconv.Atoi(match[0][4])
	return cmd
}

func (c *Cmd) isValid() bool {
	return (c.start.x == c.end.x || c.start.y == c.end.y)
}

func (c *Cmd) isHorizontal() bool {
	return c.start.y == c.end.y
}

func (c *Cmd) isVertical() bool {
	return c.start.x == c.end.x
}

func parseCmds(filename string) []Cmd {
	input := loadFile(filename)

	commands := make([]Cmd, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		cmd := MakeCmd(line)
		if cmd.isValid() {
			commands = append(commands, cmd)
		}
	}
	return commands
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
	cmds := parseCmds("input.txt")
	fmt.Printf("\nPart1: %v", part1(cmds))
	fmt.Printf("\nPart2: %v", part2())
	fmt.Println("\nDone")
}
