package tetris

func TetrisMove(piece rune, levels []int) int {
	switch piece {
	case 'I':
		return max([]int{
			maxCompletedLines(levels, 4, getPositionsPieceI1),
			maxCompletedLines(levels, 1, getPositionsPieceI2),
		})
	case 'J':
		return -1 // TODO
	case 'L':
		return max([]int{
			maxCompletedLines(levels, 2, getPositionsPieceL1),
			maxCompletedLines(levels, 3, getPositionsPieceL2),
			maxCompletedLines(levels, 2, getPositionsPieceL3),
			maxCompletedLines(levels, 3, getPositionsPieceL4),
		})
	case 'O':
		return maxCompletedLines(levels, 2, getPositionsPieceO)
	case 'S':
		return -1 // TODO
	case 'T':
		return -1 // TODO
	case 'Z':
		return -1 // TODO
	}
	return -1
}

func getPositionsPieceI1(levels []int, columnIndex int) map[int]map[int]struct{} {
	r := max(levels[columnIndex : columnIndex+4])
	positions := map[int]map[int]struct{}{}
	addPosition(positions, r, columnIndex)
	addPosition(positions, r, columnIndex+1)
	addPosition(positions, r, columnIndex+2)
	addPosition(positions, r, columnIndex+3)
	return positions
}

func getPositionsPieceI2(levels []int, columnIndex int) map[int]map[int]struct{} {
	r := levels[columnIndex]
	positions := map[int]map[int]struct{}{}
	addPosition(positions, r, columnIndex)
	addPosition(positions, r+1, columnIndex)
	addPosition(positions, r+2, columnIndex)
	addPosition(positions, r+3, columnIndex)
	return positions
}

func getPositionsPieceL1(levels []int, columnIndex int) map[int]map[int]struct{} {
	r := max(levels[columnIndex : columnIndex+2])
	positions := map[int]map[int]struct{}{}
	addPosition(positions, r, columnIndex)
	addPosition(positions, r, columnIndex+1)
	addPosition(positions, r+1, columnIndex)
	addPosition(positions, r+2, columnIndex)
	return positions
}

func getPositionsPieceL2(levels []int, columnIndex int) map[int]map[int]struct{} {
	r := max(levels[columnIndex : columnIndex+3])
	positions := map[int]map[int]struct{}{}
	addPosition(positions, r, columnIndex)
	addPosition(positions, r, columnIndex+1)
	addPosition(positions, r, columnIndex+2)
	addPosition(positions, r+1, columnIndex+2)
	return positions
}

func getPositionsPieceL3(levels []int, columnIndex int) map[int]map[int]struct{} {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	r := max([]int{a + 1, b + 3}) - 3
	positions := map[int]map[int]struct{}{}
	addPosition(positions, r, columnIndex+1)
	addPosition(positions, r+1, columnIndex+1)
	addPosition(positions, r+2, columnIndex)
	addPosition(positions, r+2, columnIndex+1)
	return positions
}

func getPositionsPieceL4(levels []int, columnIndex int) map[int]map[int]struct{} {
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	c := levels[columnIndex+2]
	r := max([]int{a + 2, b + 1, c + 1}) - 2
	positions := map[int]map[int]struct{}{}
	addPosition(positions, r, columnIndex+2)
	addPosition(positions, r+1, columnIndex)
	addPosition(positions, r+1, columnIndex+1)
	addPosition(positions, r+1, columnIndex+2)
	return positions
}

func getPositionsPieceO(levels []int, columnIndex int) map[int]map[int]struct{} {
	r := max(levels[columnIndex : columnIndex+2])
	positions := map[int]map[int]struct{}{}
	addPosition(positions, r, columnIndex)
	addPosition(positions, r, columnIndex+1)
	addPosition(positions, r+1, columnIndex)
	addPosition(positions, r+1, columnIndex+1)
	return positions
}

func addPosition(p map[int]map[int]struct{}, r int, c int) {
	if _, ok := p[r]; !ok {
		p[r] = map[int]struct{}{}
	}
	p[r][c] = struct{}{}
}

func maxCompletedLines(levels []int, width int, fn func([]int, int) map[int]map[int]struct{}) int {
	max := 0
	for i := 0; i < len(levels)-width+1; i++ {
		piecePositions := fn(levels, i)
		completedLines := countCompletedLines(levels, piecePositions)
		if completedLines > max {
			max = completedLines
		}
	}
	return max
}

func countCompletedLines(levels []int, piecePositions map[int]map[int]struct{}) int {
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

		if filledCount == len(levels) {
			completedLines++
		}

		if filledCount == 0 {
			// stop once there is a whole row that is unfilled
			break
		}
	}

	return completedLines
}

func max(numbers []int) int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
