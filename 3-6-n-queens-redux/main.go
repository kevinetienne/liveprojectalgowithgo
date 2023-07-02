package main

import (
	"fmt"
	"time"
)

const (
	empty = "."
	queen = "Q"
)

var moveOffsets = [][]int{
	{1, 0},
	{0, -1},
	{1, 1},
	{-1, 1},
}

func makeBoard(size int) [][]string {
	board := make([][]string, size)

	for i := 0; i < size; i++ {
		board[i] = make([]string, size)
		for j := 0; j < size; j++ {
			board[i][j] = empty
		}
	}

	return board
}

func dumpBoard(board [][]string) {
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%2s ", val)
		}
		fmt.Println()
	}
}

// Return true if this series of squares contains at most one queen.
func seriesIsLegal(board [][]string, size, r0, c0, dr, dc int) bool {
	queenFound := 0
	for {
		if board[r0][c0] == "Q" {
			queenFound++

			if queenFound == 2 {
				return false
			}
		}

		r0 += dr
		c0 += dc

		if r0 < 0 || r0 > size-1 {
			break
		}

		if c0 < 0 || c0 > size-1 {
			break
		}
	}

	return true
}

// Return true if the board is legal.
func boardIsLegal(board [][]string, size int) bool {
	for i, row := range board {
		for j := range row {
			for _, move := range moveOffsets {
				if ok := seriesIsLegal(board, size, i, j, move[0], move[1]); !ok {
					return false
				}
			}
		}
	}

	return true
}

// Return true if the board is legal and a solution.
func boardIsASolution(board [][]string, size int) bool {
	var numQueen int

	for _, row := range board {
		for _, val := range row {
			if val == queen {
				numQueen++
			}
		}
	}

	if numQueen != size {
		return false
	}

	return boardIsLegal(board, size)
}

// Try to place a queen in this column.
// Return true if we find a legal board.
func placeQueens4(board [][]string, size, c int) bool {
	if c == size {
		return boardIsLegal(board, size)
	}

	if c < size {
		if ok := boardIsLegal(board, size); !ok {
			return false
		}
	}

	for r := 0; r < size; r++ {
		board[r][c] = queen

		if ok := placeQueens4(board, size, c+1); ok {
			return true
		}

		board[r][c] = empty
	}

	return false
}

func main() {
	const numRows = 20

	start := time.Now()
	board := makeBoard(numRows)
	success := placeQueens4(board, numRows, 0)
	elapsed := time.Since(start)
	if success {
		fmt.Println("Success!")
		dumpBoard(board)
	} else {
		fmt.Println("No solution")
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}
