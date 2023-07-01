package main

import (
	"fmt"
	"time"
)

const (
	numRows            = 8
	numCols            = 8
	requiredClosedTour = false
	unvisited          = -1
)

type offset struct {
	dr, dc int
}

var (
	moveOffsets = []offset{
		{-2, -1},
		{-1, -2},
		{+2, -1},
		{+1, -2},
		{-2, +1},
		{-1, +2},
		{+2, +1},
		{+1, +2},
	}
	numCalls int64
)

func makeBoard(numRows, numCols int) [][]int {
	board := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		board[i] = make([]int, numCols)
		for j := 0; j < numCols; j++ {
			board[i][j] = unvisited
		}
	}

	return board
}

func dumpBoard(board [][]int) {
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%02d ", val)
		}
		fmt.Println()
	}
}

func moveAllowed(numRows, numCols, i, j int) bool {
	if i < 0 || i > numRows-1 {
		return false
	}

	if j < 0 || j > numCols-1 {
		return false
	}

	return true
}

// Try to extend a knight's tour starting at (startRow, startCol).
// Return true or false to indicate whether we have found a solution.
func findTour(board [][]int, numRows, numCols, curRow, curCol, numVisited int) bool {
	numCalls++

	if numVisited == numRows*numCols {
		if requiredClosedTour {
			for _, move := range moveOffsets {
				i := curRow + move.dr
				j := curCol + move.dc

				if !moveAllowed(numRows, numCols, i, j) {
					continue
				}

				if board[i][j] == 0 {
					return true
				}
			}

			return false
		}

		return true
	}

	for _, move := range moveOffsets {
		i := curRow + move.dr
		j := curCol + move.dc

		if !moveAllowed(numRows, numCols, i, j) {
			continue
		}

		if board[i][j] != -1 {
			continue
		}

		board[i][j] = numVisited
		if findTour(board, numRows, numCols, i, j, numVisited+1) {
			return true
		}

		board[i][j] = unvisited
	}

	return false
}

func main() {
	numCalls = 0

	// Create the blank board.
	board := makeBoard(numRows, numCols)

	// Try to find a tour.
	start := time.Now()
	board[0][0] = 0
	if findTour(board, numRows, numCols, 0, 0, 1) {
		fmt.Println("Success!")
	} else {
		fmt.Println("Could not find a tour.")
	}
	elapsed := time.Since(start)
	dumpBoard(board)
	fmt.Printf("%f seconds\n", elapsed.Seconds())
	fmt.Printf("%d calls\n", numCalls)
}
