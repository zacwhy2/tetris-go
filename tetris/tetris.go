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
		return maxCompletedLines(levels, 2, getPositionsPieceL3)
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

func getPositionsPieceL3(levels []int, columnIndex int) map[int]map[int]struct{} {
	positions := map[int]map[int]struct{}{}
	a := levels[columnIndex]
	b := levels[columnIndex+1]
	max := max([]int{a + 1, b + 3}) - 3
	positions[max] = map[int]struct{}{}
	positions[max][columnIndex+1] = struct{}{}
	positions[max+1] = map[int]struct{}{}
	positions[max+1][columnIndex+1] = struct{}{}
	positions[max+2] = map[int]struct{}{}
	positions[max+2][columnIndex] = struct{}{}
	positions[max+2][columnIndex+1] = struct{}{}
	return positions
}

func getPositionsPieceI1(levels []int, columnIndex int) map[int]map[int]struct{} {
	max := max(levels[columnIndex : columnIndex+4])
	positions := map[int]map[int]struct{}{}
	positions[max] = map[int]struct{}{}
	for i := 0; i < 4; i++ {
		positions[max][i+columnIndex] = struct{}{}
	}
	return positions
}

func getPositionsPieceI2(levels []int, columnIndex int) map[int]map[int]struct{} {
	positions := map[int]map[int]struct{}{}
	level := levels[columnIndex]
	for i := 0; i < 4; i++ {
		positions[level+i] = map[int]struct{}{}
		positions[level+i][columnIndex] = struct{}{}
	}
	return positions
}

func getPositionsPieceO(levels []int, columnIndex int) map[int]map[int]struct{} {
	max := max(levels[columnIndex : columnIndex+2])
	positions := map[int]map[int]struct{}{}
	positions[max] = map[int]struct{}{}
	positions[max][columnIndex] = struct{}{}
	positions[max][columnIndex+1] = struct{}{}
	positions[max+1] = map[int]struct{}{}
	positions[max+1][columnIndex] = struct{}{}
	positions[max+1][columnIndex+1] = struct{}{}
	return positions
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
