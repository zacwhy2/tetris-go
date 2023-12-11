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
		{piece: 'I', levels: []int{2, 4, 3, 4, 5, 2, 0, 2, 2, 3, 3, 3}, want: 2},
		{piece: 'O', levels: []int{4, 3, 2, 3, 5, 1, 0, 1, 2, 4, 3, 4}, want: 0},
	}

	for _, tt := range tests {
		if got := TetrisMove(tt.piece, tt.levels); got != tt.want {
			t.Errorf("TetrisMove(%c, %v) = %d, want %d", tt.piece, tt.levels, got, tt.want)
		}
	}
}
