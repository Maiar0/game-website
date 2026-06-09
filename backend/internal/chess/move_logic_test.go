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

func TestInBounds(t *testing.T) {
	tests := []struct {
		name string
		pos  Position
		want bool
	}{
		{"top left corner", Position{row: 0, col: 0}, true},
		{"bottom right corner", Position{row: 7, col: 7}, true},
		{"middle square", Position{row: 3, col: 4}, true},

		{"row too low", Position{row: -1, col: 4}, false},
		{"row too high", Position{row: 8, col: 4}, false},
		{"col too low", Position{row: 4, col: -1}, false},
		{"col too high", Position{row: 4, col: 8}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, inBounds(tt.pos))
		})
	}
}

func TestCheckPath(t *testing.T) {
	emptyBoard := func() Board {
		var b Board
		for row := range b {
			for col := range b[row] {
				b[row][col] = '.'
			}
		}
		return b
	}

	tests := []struct {
		name    string
		from    Position
		to      Position
		blocker *Position
		want    bool
	}{
		{
			name: "clear vertical path upward",
			from: Position{row: 0, col: 4},
			to:   Position{row: 5, col: 4},
			want: true,
		},
		{
			name:    "blocked vertical path upward",
			from:    Position{row: 0, col: 4},
			to:      Position{row: 5, col: 4},
			blocker: &Position{row: 3, col: 4},
			want:    false,
		},
		{
			name: "clear horizontal path right",
			from: Position{row: 4, col: 0},
			to:   Position{row: 4, col: 6},
			want: true,
		},
		{
			name:    "blocked horizontal path right",
			from:    Position{row: 4, col: 0},
			to:      Position{row: 4, col: 6},
			blocker: &Position{row: 4, col: 3},
			want:    false,
		},
		{
			name: "clear diagonal path",
			from: Position{row: 1, col: 1},
			to:   Position{row: 5, col: 5},
			want: true,
		},
		{
			name:    "blocked diagonal path",
			from:    Position{row: 1, col: 1},
			to:      Position{row: 5, col: 5},
			blocker: &Position{row: 3, col: 3},
			want:    false,
		},
		{
			name:    "piece on destination does not block path",
			from:    Position{row: 1, col: 1},
			to:      Position{row: 5, col: 5},
			blocker: &Position{row: 5, col: 5},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := emptyBoard()

			if tt.blocker != nil {
				b[tt.blocker.row][tt.blocker.col] = 'p'
			}
			PrintBoard(b)

			require.Equal(t, tt.want, CheckPath(b, tt.from, tt.to))
		})
	}
}
