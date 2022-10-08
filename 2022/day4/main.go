package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(game Game) int {
	return game.getPuzzleSolution()
}

func part2(game Game) int {
	return game.getPuzzleSolution2()
}

type Board struct {
	cells       [][]int
	winner      bool
	winningDraw int
	drawsAtWin  []int
}

type Game struct {
	boards []Board
	draws  []int
	drawn  []int
}

func isNumberDrawn(drawn []int, number int) bool {
	for _, d := range drawn {
		if d == number {
			return true
		}
	}
	return false
}

func isRowWinner(drawn []int, row []int) bool {
	for _, c := range row {
		if !isNumberDrawn(drawn, c) {
			return false
		}
	}
	return true
}

func getColumn(cells [][]int, col int) []int {
	column := make([]int, 0)
	for i := 0; i < 5; i++ {
		for _, row := range cells {
			column = append(column, row[col])
		}
	}
	return column
}

func isColumnWinner(drawn []int, cells [][]int, col int) bool {
	//need a board obj
	for i := 0; i < 5; i++ {
		if isRowWinner(drawn, getColumn(cells, i)) {
			return true
		}
	}
	return false
}
func (g *Game) WinningBoard() (Board, error) {
	for _, b := range g.boards {
		if b.isWinner(g.drawn) {
			return b, nil
		}
	}
	return Board{}, errors.New("No Winner")
}

func (g *Game) FindWinner() Board {
	g.drawn = make([]int, 0)
	for _, d := range g.draws {
		g.drawn = append(g.drawn, d)
		for _, b := range g.boards {
			if b.isWinner(g.drawn) {
				return b
			}
		}
	}
	return Board{}
}

func getWinnersAndLosers(draws []int, boards []Board) ([]Board, []Board) {
	winner := make([]Board, 0)
	loser := make([]Board, 0)
	for _, b := range boards {
		if b.isWinner(draws) {
			winner = append(winner, b)
			b.winningDraw = draws[len(draws)-1]
			b.drawsAtWin = draws
			b.winner = true
		} else {
			loser = append(loser, b)
		}
	}
	return winner, loser
}

func (g *Game) FindLastWinner() Board {
	g.drawn = make([]int, 0)
	boards := make([]Board, len(g.boards))
	copy(boards, g.boards)
	allWinners := make([]Board, 0)
	var winners []Board
	for _, d := range g.draws {
		g.drawn = append(g.drawn, d)
		winners, boards = getWinnersAndLosers(g.drawn, boards)
		allWinners = append(allWinners, winners...)
		if len(boards) == 0 {
			fmt.Println("Done")
			break
		}
		fmt.Println("\n\nDrawn ", g.drawn)
		fmt.Println("Winners")
		for _, b := range allWinners {
			fmt.Print(b.cells[0][0], ", ")
		}
		fmt.Println("\nLosers")
		for _, b := range boards {
			fmt.Print(b.cells[0][0], ", ")
		}
		fmt.Println()
	}
	return allWinners[len(allWinners)-1]
}

func (b *Board) isWinner(drawn []int) bool {
	for i := 0; i < 5; i++ {
		if isRowWinner(drawn, b.cells[i]) {
			b.winner = true
			b.winningDraw = drawn[len(drawn)-1]
			b.drawsAtWin = drawn
			return true
		}
		if isColumnWinner(drawn, b.cells, i) {
			b.winner = true
			b.winningDraw = drawn[len(drawn)-1]
			b.drawsAtWin = drawn
			return true
		}
	}
	return false
}

func (b *Board) justGetAllCells() []int {
	allCells := make([]int, 0)
	for _, row := range b.cells {
		for _, cell := range row {
			allCells = append(allCells, cell)
		}
	}
	return allCells
}

func (g *Game) getPuzzleSolution2() int {
	board := g.FindLastWinner()
	unmarked := 0
	for _, c := range board.justGetAllCells() {
		if !isNumberDrawn(board.drawsAtWin, c) {
			unmarked = unmarked + c
		}
	}
	fmt.Println("unmarked", unmarked)
	fmt.Println("winning draw", board)
	return unmarked * board.winningDraw
}

func (g *Game) getPuzzleSolution() int {
	board := g.FindWinner()
	unmarked := 0
	for _, c := range board.justGetAllCells() {
		if !isNumberDrawn(g.drawn, c) {
			unmarked = unmarked + c
		}
	}
	return unmarked * g.drawn[len(g.drawn)-1]
}

func readInput(filename string) Game {
	game := Game{}
	input := loadFile(filename)
	inputBlocks := strings.Split(input, "\n\n")
	reNumber := regexp.MustCompile("(\\d+)")
	game.draws = make([]int, 0)
	for _, v := range reNumber.FindAllString(inputBlocks[0], -1) {
		d, _ := strconv.Atoi(v)
		game.draws = append(game.draws, d)
	}
	game.boards = make([]Board, 0)
	for _, v := range inputBlocks[1:] {
		board := Board{}
		rawInput := reNumber.FindAllString(v, -1)
		board.cells = make([][]int, 5)
		for i := 0; i < 5; i++ {
			board.cells[i] = make([]int, 5)
			for j := 0; j < 5; j++ {
				v, _ := strconv.Atoi(rawInput[(i*5)+j])
				board.cells[i][j] = v
			}
		}
		game.boards = append(game.boards, board)
	}
	return game
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
	game := readInput("input.txt")

	fmt.Printf("\nPart1: %v", part1(game))
	fmt.Printf("\nPart2: %v", part2(game))
	fmt.Println("\nDone")
}
