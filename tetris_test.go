package tetris

import (
	"testing"
)

func TestTetrisMove(t *testing.T) {
	tests := []struct {
		piece  rune
		levels []int
		want   int
	}{
		{piece: 'L', levels: []int{3, 4, 4, 5, 6, 2, 0, 6, 5, 3, 6, 6}, want: 3},
		{piece: 'I', levels: []int{2, 4, 3, 4, 5, 2, 0, 2, 2, 3, 3, 3}, want: 2},
		{piece: 'O', levels: []int{4, 3, 2, 3, 5, 1, 0, 1, 2, 4, 3, 4}, want: 0},

		// with already completed line
		{piece: 'L', levels: []int{4, 4, 4, 5, 6, 3, 1, 6, 5, 4, 6, 6}, want: 4},

		{piece: 'I', levels: []int{5, 5, 5, 5, 0, 0, 0, 0, 5, 5, 5, 5}, want: 1},
		{piece: 'I', levels: []int{5, 5, 5, 5, 5, 5, 0, 5, 5, 5, 5, 5}, want: 4},

		{piece: 'J', levels: []int{5, 5, 5, 5, 5, 5, 0, 0, 5, 5, 5, 5}, want: 1},
		{piece: 'J', levels: []int{5, 5, 5, 5, 5, 0, 0, 0, 5, 5, 5, 5}, want: 1},
		{piece: 'J', levels: []int{5, 5, 5, 5, 5, 5, 0, 2, 5, 5, 5, 5}, want: 3},
		{piece: 'J', levels: []int{5, 5, 5, 5, 5, 1, 1, 0, 5, 5, 5, 5}, want: 2},

		{piece: 'L', levels: []int{5, 5, 5, 5, 5, 5, 0, 0, 5, 5, 5, 5}, want: 1},
		{piece: 'L', levels: []int{5, 5, 5, 5, 5, 0, 1, 1, 5, 5, 5, 5}, want: 2},
		{piece: 'L', levels: []int{5, 5, 5, 5, 5, 5, 2, 0, 5, 5, 5, 5}, want: 3},
		{piece: 'L', levels: []int{5, 5, 5, 5, 5, 0, 0, 0, 5, 5, 5, 5}, want: 1},

		{piece: 'O', levels: []int{5, 5, 5, 5, 5, 5, 0, 0, 5, 5, 5, 5}, want: 2},

		{piece: 'S', levels: []int{5, 5, 5, 5, 5, 0, 0, 1, 5, 5, 5, 5}, want: 1},
		{piece: 'S', levels: []int{5, 5, 5, 5, 5, 1, 0, 5, 5, 5, 5, 5}, want: 2},

		{piece: 'T', levels: []int{5, 5, 5, 5, 5, 0, 0, 0, 5, 5, 5, 5}, want: 1},
		{piece: 'T', levels: []int{5, 5, 5, 5, 5, 5, 0, 1, 5, 5, 5, 5}, want: 2},
		{piece: 'T', levels: []int{5, 5, 5, 5, 5, 1, 0, 1, 5, 5, 5, 5}, want: 2},
		{piece: 'T', levels: []int{5, 5, 5, 5, 5, 1, 0, 5, 5, 5, 5, 5}, want: 2},

		{piece: 'Z', levels: []int{5, 5, 5, 5, 5, 1, 0, 0, 5, 5, 5, 5}, want: 1},
		{piece: 'Z', levels: []int{5, 5, 5, 5, 5, 0, 1, 5, 5, 5, 5, 5}, want: 2},
	}

	for _, tt := range tests {
		if got := TetrisMove(tt.piece, tt.levels); got != tt.want {
			t.Errorf("TetrisMove(%c, %v) = %d, want %d", tt.piece, tt.levels, got, tt.want)
		}
	}
}
