package main

import (
	"fmt"
	"tetris/tetris"
)

func main() {
	piece := 'I'
	levels := []int{2, 4, 3, 4, 5, 2, 0, 2, 2, 3, 3, 3}
	completedLines := tetris.TetrisMove(piece, levels)
	fmt.Printf("TetrisMove(%c, %v) = %d", piece, levels, completedLines)
}
