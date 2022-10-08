package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	name              string
	inputFile         string
	expectedDraws     []int
	expectedLastBoard [][]int
}

func Test_Part1_details(t *testing.T) {
	testcase := TestCase{
		name:              "input test",
		inputFile:         "test_input.txt",
		expectedDraws:     []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
		expectedLastBoard: [][]int{{14, 21, 17, 24, 4}, {10, 16, 15, 9, 19}, {18, 8, 23, 26, 20}, {22, 11, 13, 6, 5}, {2, 0, 12, 3, 7}},
	}
	t.Run(testcase.name, func(t *testing.T) {
		res := readInput("test_input.txt")
		if !reflect.DeepEqual(res.draws, testcase.expectedDraws) {
			t.Errorf("part1() = %v, wanted %v", res.draws, testcase.expectedDraws)
		}
		lastBoard := res.boards[len(res.boards)-1].cells
		for i := 0; i < 5; i++ {
			if !reflect.DeepEqual(lastBoard[i], testcase.expectedLastBoard[i]) {
				t.Errorf("part1() = %v, wanted %v", lastBoard, testcase.expectedLastBoard)
			}
		}
		testBoard := Board{cells: testcase.expectedLastBoard}
		for i := 0; i < 5; i++ {
			if !testBoard.isWinner(testcase.expectedLastBoard[i]) {
				t.Errorf("isWinner for Row %v failed, wanted true got %v", i, testBoard.isWinner(testcase.expectedLastBoard[i]))
			}
		}
		for i := 0; i < 5; i++ {
			drawn := []int{70, 41, 9, 5, 1, 0, -1}
			if testBoard.isWinner(drawn) {
				t.Errorf("isWinner for Row %v failed, wanted false got %v", i, testBoard.isWinner(drawn))
			}
		}
		for i := 0; i < 5; i++ {
			drawn := make([]int, 0)
			for _, row := range testBoard.cells {
				drawn = append(drawn, row[i])
			}
			if !testBoard.isWinner(drawn) {
				fmt.Printf("Drawn = %v\n", drawn)
				for _, row := range testBoard.cells {
					fmt.Println(row)
				}
				t.Errorf("isWinner for Column %v failed, wanted true got %v", i, testBoard.isWinner(drawn))
			}
		}
	})
}
func Test_Part1_Board(t *testing.T) {
	testcase := TestCase{
		name:      "Board Winner test",
		inputFile: "test_input.txt",
	}
	game := readInput(testcase.inputFile)
	t.Run("No Winner", func(t *testing.T) {
		_, err := game.WinningBoard()
		if err == nil {
			t.Errorf("Expected no winner, got %v", err)
			return
		}
	})
	t.Run(testcase.name, func(t *testing.T) {
		game := readInput("test_input.txt")
		winner := game.FindWinner()
		winner14 := winner.cells[0][0]
		if !(winner14 == 14) {
			t.Errorf("Wrong Winning Board, expected the third, got %v", winner14)
		}
	})
	t.Run("Puzzle", func(t *testing.T) {
		game := readInput("test_input.txt")
		res := game.getPuzzleSolution()
		if !(res == 4512) {
			t.Errorf("Wrong, you get no stars , expected the 4512, got %v", res)
		}
	})
	t.Run("Puzzle - part 2", func(t *testing.T) {
		game := readInput("test_input.txt")
		res := game.getPuzzleSolution2()
		if !(res == 1924) {
			t.Errorf("Wrong, you get no stars , expected the 1924, got %v", res)
		}
	})

}
