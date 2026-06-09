package chess

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateMovePattern_ValidMoves(t *testing.T) {
	tests := []struct {
		name  string
		piece rune
		from  Position
		to    Position
	}{
		{"white pawn forward", 'P', Position{row: 1, col: 4}, Position{row: 2, col: 4}},
		{"white pawn double", 'P', Position{row: 1, col: 4}, Position{row: 3, col: 4}},
		{"white pawn capture shape", 'P', Position{row: 1, col: 4}, Position{row: 2, col: 5}},
		{"black pawn forward", 'p', Position{row: 6, col: 4}, Position{row: 5, col: 4}},
		{"black pawn double", 'p', Position{row: 6, col: 4}, Position{row: 4, col: 4}},
		{"black pawn capture shape", 'p', Position{row: 6, col: 4}, Position{row: 5, col: 5}},
		{"rook horizontal", 'R', Position{row: 0, col: 0}, Position{row: 0, col: 7}},
		{"rook vertical", 'r', Position{row: 0, col: 0}, Position{row: 7, col: 0}},
		{"knight move", 'N', Position{row: 3, col: 3}, Position{row: 1, col: 2}},
		{"bishop diagonal", 'B', Position{row: 0, col: 0}, Position{row: 3, col: 3}},
		{"queen diagonal", 'Q', Position{row: 0, col: 0}, Position{row: 4, col: 4}},
		{"queen straight", 'q', Position{row: 0, col: 0}, Position{row: 0, col: 5}},
		{"king one square", 'K', Position{row: 4, col: 4}, Position{row: 5, col: 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := ValidateMovePattern(tt.piece, tt.from, tt.to)

			require.NoError(t, err)
			require.True(t, ok)
		})
	}
}

func TestValidateMovePattern_InvalidMoves(t *testing.T) {
	tests := []struct {
		name  string
		piece rune
		from  Position
		to    Position
	}{
		{"rook diagonal", 'R', Position{row: 0, col: 0}, Position{row: 3, col: 3}},
		{"bishop straight", 'B', Position{row: 0, col: 0}, Position{row: 0, col: 3}},
		{"knight straight", 'N', Position{row: 0, col: 0}, Position{row: 0, col: 2}},
		{"king too far", 'K', Position{row: 0, col: 0}, Position{row: 2, col: 2}},
		{"queen invalid", 'Q', Position{row: 0, col: 0}, Position{row: 2, col: 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := ValidateMovePattern(tt.piece, tt.from, tt.to)

			require.NoError(t, err)
			require.False(t, ok)
		})
	}
}

func TestValidateMovePattern_SamePositionReturnsError(t *testing.T) {
	ok, err := ValidateMovePattern('R', Position{row: 0, col: 0}, Position{row: 0, col: 0})

	require.Error(t, err)
	require.False(t, ok)
}

func TestValidateMovePattern_InvalidPieceReturnsError(t *testing.T) {
	ok, err := ValidateMovePattern('x', Position{row: 0, col: 0}, Position{row: 0, col: 1})

	require.Error(t, err)
	require.False(t, ok)
}


