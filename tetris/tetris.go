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
		return max(
			maxCompletedLines(levels, 2, getPositionsPieceJ1),
			maxCompletedLines(levels, 3, getPositionsPieceJ2),
			maxCompletedLines(levels, 2, getPositionsPieceJ3),
			maxCompletedLines(levels, 3, getPositionsPieceJ4),
		)
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
		return max(
			maxCompletedLines(levels, 3, getPositionsPieceS1),
			maxCompletedLines(levels, 2, getPositionsPieceS2),
		)
	case 'T':
		return max(
			maxCompletedLines(levels, 3, getPositionsPieceT1),
			maxCompletedLines(levels, 2, getPositionsPieceT2),
			maxCompletedLines(levels, 3, getPositionsPieceT3),
			maxCompletedLines(levels, 2, getPositionsPieceT4),
		)
	case 'Z':
		return max(
			maxCompletedLines(levels, 3, getPositionsPieceZ1),
			maxCompletedLines(levels, 2, getPositionsPieceZ2),
		)
	}
	return -1
}

/*
I1 looks like:

	XXXX
*/
func getPositionsPieceI1(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	c := levels[column+2]
	d := levels[column+3]
	row := max(a+1, b+1, c+1, d+1) - 1
	return []position{
		{row + 0, column + 0},
		{row + 0, column + 1},
		{row + 0, column + 2},
		{row + 0, column + 3},
	}
}

/*
I2 looks like:

	X
	X
	X
	X
*/
func getPositionsPieceI2(levels []int, column int) []position {
	a := levels[column+0]
	row := max(a+1) - 1
	return []position{
		{row + 0, column + 0},
		{row + 1, column + 0},
		{row + 2, column + 0},
		{row + 3, column + 0},
	}
}

/*
J1 looks like:

	 X
	 X
	XX
*/
func getPositionsPieceJ1(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	row := max(a+1, b+1) - 1
	return []position{
		{row + 0, column + 0},
		{row + 0, column + 1},
		{row + 1, column + 1},
		{row + 2, column + 1},
	}
}

/*
J2 looks like:

	X
	XXX
*/
func getPositionsPieceJ2(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	c := levels[column+2]
	row := max(a+1, b+1, c+1) - 1
	return []position{
		{row + 0, column + 0},
		{row + 0, column + 1},
		{row + 0, column + 2},
		{row + 1, column + 0},
	}
}

/*
J3 looks like:

	XX
	X
	X
*/
func getPositionsPieceJ3(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	row := max(a+3, b+1) - 3
	return []position{
		{row + 0, column + 0},
		{row + 1, column + 0},
		{row + 2, column + 0},
		{row + 2, column + 1},
	}
}

/*
J4 looks like:

	XXX
	  X
*/
func getPositionsPieceJ4(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	c := levels[column+2]
	row := max(a+1, b+1, c+2) - 2
	return []position{
		{row + 0, column + 2},
		{row + 1, column + 0},
		{row + 1, column + 1},
		{row + 1, column + 2},
	}
}

/*
L1 looks like:

	X
	X
	XX
*/
func getPositionsPieceL1(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	row := max(a+1, b+1) - 1
	return []position{
		{row + 0, column + 0},
		{row + 0, column + 1},
		{row + 1, column + 0},
		{row + 2, column + 0},
	}
}

/*
L2 looks like:

	XXX
	X
*/
func getPositionsPieceL2(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	c := levels[column+2]
	row := max(a+2, b+1, c+1) - 2
	return []position{
		{row + 0, column + 0},
		{row + 1, column + 0},
		{row + 1, column + 1},
		{row + 1, column + 2},
	}
}

/*
L3 looks like:

	XX
	 X
	 X
*/
func getPositionsPieceL3(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	row := max(a+1, b+3) - 3
	return []position{
		{row + 0, column + 1},
		{row + 1, column + 1},
		{row + 2, column + 0},
		{row + 2, column + 1},
	}
}

/*
L4 looks like:

	  X
	XXX
*/
func getPositionsPieceL4(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	c := levels[column+2]
	row := max(a+1, b+1, c+1) - 1
	return []position{
		{row + 0, column + 0},
		{row + 0, column + 1},
		{row + 0, column + 2},
		{row + 1, column + 2},
	}
}

/*
O looks like:

	XX
	XX
*/
func getPositionsPieceO(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	row := max(a+1, b+1) - 1
	return []position{
		{row + 0, column + 0},
		{row + 0, column + 1},
		{row + 1, column + 0},
		{row + 1, column + 1},
	}
}

/*
S1 looks like:

	 XX
	XX
*/
func getPositionsPieceS1(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	c := levels[column+2]
	row := max(a+2, b+2, c+1) - 2
	return []position{
		{row + 0, column + 0},
		{row + 0, column + 1},
		{row + 1, column + 1},
		{row + 1, column + 2},
	}
}

/*
S2 looks like:

	X
	XX
	 X
*/
func getPositionsPieceS2(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	row := max(a+1, b+2) - 2
	return []position{
		{row + 0, column + 1},
		{row + 1, column + 0},
		{row + 1, column + 1},
		{row + 2, column + 0},
	}
}

/*
T1 looks like:

	 X
	XXX
*/
func getPositionsPieceT1(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	c := levels[column+2]
	row := max(a+1, b+1, c+1) - 1
	return []position{
		{row + 0, column + 0},
		{row + 0, column + 1},
		{row + 0, column + 2},
		{row + 1, column + 1},
	}
}

/*
T2 looks like:

	X
	XX
	X
*/
func getPositionsPieceT2(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	row := max(a+2, b+1) - 2
	return []position{
		{row + 0, column + 0},
		{row + 1, column + 0},
		{row + 1, column + 1},
		{row + 2, column + 0},
	}
}

/*
T3 looks like:

	XXX
	 X
*/
func getPositionsPieceT3(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	c := levels[column+2]
	row := max(a+1, b+2, c+1) - 2
	return []position{
		{row + 0, column + 1},
		{row + 1, column + 0},
		{row + 1, column + 1},
		{row + 1, column + 2},
	}
}

/*
T4 looks like:

	 X
	XX
	 X
*/
func getPositionsPieceT4(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	row := max(a+1, b+2) - 2
	return []position{
		{row + 0, column + 1},
		{row + 1, column + 0},
		{row + 1, column + 1},
		{row + 2, column + 1},
	}
}

/*
Z1 looks like:

	XX
	 XX
*/
func getPositionsPieceZ1(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	c := levels[column+2]
	row := max(a+1, b+2, c+2) - 2
	return []position{
		{row + 0, column + 1},
		{row + 0, column + 2},
		{row + 1, column + 0},
		{row + 1, column + 1},
	}
}

/*
Z2 looks like:

	 X
	XX
	X
*/
func getPositionsPieceZ2(levels []int, column int) []position {
	a := levels[column+0]
	b := levels[column+1]
	row := max(a+2, b+1) - 2
	return []position{
		{row + 0, column + 0},
		{row + 1, column + 0},
		{row + 1, column + 1},
		{row + 2, column + 1},
	}
}

func maxCompletedLines(
	levels []int,
	width int,
	getPositions func([]int, int) []position,
) int {
	max := 0
	for i := 0; i < len(levels)-width+1; i++ {
		piecePositions := positionsArrayToMap(getPositions(levels, i))
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
