package tetris

type position struct{ row, column int }

// TetrisMove calculates the greatest number of horizontal lines that can be completed when the piece arrives at the bottom assuming it is dropped immediately after being rotated and moved horizontally from the top
func TetrisMove(piece rune, levels []int) int {
	switch piece {
	case 'I':
		return max(
			maxCompletedLines(levels, 4, getPositionsPieceI1),
			maxCompletedLines(levels, 1, getPositionsPieceI2),
		)
	case 'J':
		return -1 // TODO
	case 'L':
		return max(
			maxCompletedLines(levels, 2, getPositionsPieceL1),
			maxCompletedLines(levels, 3, getPositionsPieceL2),
			maxCompletedLines(levels, 2, getPositionsPieceL3),
			maxCompletedLines(levels, 3, getPositionsPieceL4),
		)
	case 'O':
		return maxCompletedLines(levels, 2, getPositionsPieceO)
	case 'S':
		return -1 // TODO
	case 'T':
		return max(
			maxCompletedLines(levels, 3, getPositionsPieceT1),
			maxCompletedLines(levels, 2, getPositionsPieceT2),
			maxCompletedLines(levels, 3, getPositionsPieceT3),
			maxCompletedLines(levels, 2, getPositionsPieceT4),
		)
	case 'Z':
		return -1 // TODO
	}
	return -1
}

/*
I1 looks like:

	XXXX
*/
func getPositionsPieceI1(levels []int, columnIndex int) []position {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	c := levels[columnIndex+2]
	d := levels[columnIndex+3]
	r := max(a, b, c, d)
	return []position{
		position{r, columnIndex},
		position{r, columnIndex + 1},
		position{r, columnIndex + 2},
		position{r, columnIndex + 3},
	}
}

/*
I2 looks like:

	X
	X
	X
	X
*/
func getPositionsPieceI2(levels []int, columnIndex int) []position {
	r := levels[columnIndex]
	return []position{
		position{r, columnIndex},
		position{r + 1, columnIndex},
		position{r + 2, columnIndex},
		position{r + 3, columnIndex},
	}
}

/*
L1 looks like:

	X
	X
	XX
*/
func getPositionsPieceL1(levels []int, columnIndex int) []position {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	r := max(a, b)
	return []position{
		position{r, columnIndex},
		position{r, columnIndex + 1},
		position{r + 1, columnIndex},
		position{r + 2, columnIndex},
	}
}

/*
L2 looks like:

	  X
	XXX
*/
func getPositionsPieceL2(levels []int, columnIndex int) []position {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	c := levels[columnIndex+2]
	r := max(a, b, c)
	return []position{
		position{r, columnIndex},
		position{r, columnIndex + 1},
		position{r, columnIndex + 2},
		position{r + 1, columnIndex + 2},
	}
}

/*
L3 looks like:

	XX
	 X
	 X
*/
func getPositionsPieceL3(levels []int, columnIndex int) []position {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	r := max(a+1, b+3) - 3
	return []position{
		position{r, columnIndex + 1},
		position{r + 1, columnIndex + 1},
		position{r + 2, columnIndex},
		position{r + 2, columnIndex + 1},
	}
}

/*
L4 looks like:

	XXX
	X
*/
func getPositionsPieceL4(levels []int, columnIndex int) []position {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	c := levels[columnIndex+2]
	r := max(a+2, b+1, c+1) - 2
	return []position{
		position{r, columnIndex + 2},
		position{r + 1, columnIndex},
		position{r + 1, columnIndex + 1},
		position{r + 1, columnIndex + 2},
	}
}

/*
O looks like:

	XX
	XX
*/
func getPositionsPieceO(levels []int, columnIndex int) []position {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	r := max(a, b)
	return []position{
		position{r, columnIndex},
		position{r, columnIndex + 1},
		position{r + 1, columnIndex},
		position{r + 1, columnIndex + 1},
	}
}

/*
T1 looks like:

	 X
	XXX
*/
func getPositionsPieceT1(levels []int, columnIndex int) []position {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	c := levels[columnIndex+2]
	r := max(a, b, c)
	return []position{
		position{r, columnIndex},
		position{r, columnIndex + 1},
		position{r, columnIndex + 2},
		position{r + 1, columnIndex + 1},
	}
}

/*
T2 looks like:

	X
	XX
	X
*/
func getPositionsPieceT2(levels []int, columnIndex int) []position {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	r := max(a+1, b) - 2
	return []position{
		position{r, columnIndex},
		position{r + 1, columnIndex},
		position{r + 1, columnIndex + 1},
		position{r + 2, columnIndex},
	}
}

/*
T3 looks like:

	 X
	XX
	 X
*/
func getPositionsPieceT3(levels []int, columnIndex int) []position {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	c := levels[columnIndex+2]
	r := max(a+1, b+2, c+1) - 2
	return []position{
		position{r, columnIndex + 1},
		position{r + 1, columnIndex},
		position{r + 1, columnIndex + 1},
		position{r + 1, columnIndex + 2},
	}
}

/*
T4 looks like:

	XXX
	 X
*/
func getPositionsPieceT4(levels []int, columnIndex int) []position {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	r := max(a, b+1) - 2
	return []position{
		position{r, columnIndex + 1},
		position{r + 1, columnIndex},
		position{r + 1, columnIndex + 1},
		position{r + 2, columnIndex + 1},
	}
}

func maxCompletedLines(
	levels []int,
	width int,
	fn func([]int, int) []position,
) int {
	max := 0
	for i := 0; i < len(levels)-width+1; i++ {
		piecePositions := positionsArrayToMap(fn(levels, i))
		completedLines := countCompletedLines(levels, piecePositions)
		if completedLines > max {
			max = completedLines
		}
	}
	return max
}

func positionsArrayToMap(ps []position) map[int]map[int]struct{} {
	positions := map[int]map[int]struct{}{}
	for _, p := range ps {
		if _, ok := positions[p.row]; !ok {
			positions[p.row] = map[int]struct{}{}
		}
		positions[p.row][p.column] = struct{}{}
	}
	return positions
}

func countCompletedLines(
	levels []int,
	piecePositions map[int]map[int]struct{},
) int {
	completedLines := 0
	for rowIndex := 0; ; rowIndex++ {
		filledCount := 0
		for columnIndex, level := range levels {
			if level > rowIndex {
				filledCount++
			} else if _, ok := piecePositions[rowIndex][columnIndex]; ok {
				filledCount++
			}
		}
		if filledCount == 0 {
			// stop once there is a whole row that is unfilled
			break
		}
		if filledCount == len(levels) {
			completedLines++
		}
	}
	return completedLines
}

func max(numbers ...int) int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
