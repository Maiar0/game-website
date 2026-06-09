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
		{"white kingside castle", 'K', Position{row: 0, col: 4}, Position{row: 0, col: 6}},
		{"white queenside castle", 'K', Position{row: 0, col: 4}, Position{row: 0, col: 2}},
		{"black kingside castle", 'k', Position{row: 7, col: 4}, Position{row: 7, col: 6}},
		{"black queenside castle", 'k', Position{row: 7, col: 4}, Position{row: 7, col: 2}},
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
		{"king two squares vertical", 'K', Position{row: 4, col: 4}, Position{row: 6, col: 4}},
		{"king two squares diagonal", 'K', Position{row: 4, col: 4}, Position{row: 6, col: 6}},
		{"king two squares horizontal not castle", 'K', Position{row: 4, col: 4}, Position{row: 4, col: 6}},
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

func TestIsSquareAttacked(t *testing.T) {
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
		name        string
		pos         Position
		attackColor rune
		pieces      map[Position]rune
		want        bool
	}{
		{
			name:        "attacked by black rook vertical negative direction",
			pos:         Position{row: 4, col: 4},
			attackColor: 'b',
			pieces: map[Position]rune{
				Position{row: 1, col: 4}: 'r',
			},
			want: true,
		},
		{
			name:        "attacked by black queen vertical positive direction",
			pos:         Position{row: 4, col: 4},
			attackColor: 'b',
			pieces: map[Position]rune{
				Position{row: 7, col: 4}: 'q',
			},
			want: true,
		},
		{
			name:        "attacked by white bishop diagonal negative positive",
			pos:         Position{row: 4, col: 4},
			attackColor: 'w',
			pieces: map[Position]rune{
				Position{row: 2, col: 6}: 'B',
			},
			want: true,
		},
		{
			name:        "attacked by white queen diagonal positive positive",
			pos:         Position{row: 4, col: 4},
			attackColor: 'w',
			pieces: map[Position]rune{
				Position{row: 6, col: 6}: 'Q',
			},
			want: true,
		},
		{
			name:        "attacked by black knight",
			pos:         Position{row: 4, col: 4},
			attackColor: 'b',
			pieces: map[Position]rune{
				Position{row: 6, col: 5}: 'n',
			},
			want: true,
		},
		{
			name:        "same color requested ignores opposite color rook",
			pos:         Position{row: 4, col: 4},
			attackColor: 'b',
			pieces: map[Position]rune{
				Position{row: 1, col: 4}: 'R',
			},
			want: false,
		},
		{
			name:        "blocked rook does not attack",
			pos:         Position{row: 4, col: 4},
			attackColor: 'b',
			pieces: map[Position]rune{
				Position{row: 3, col: 4}: 'p',
				Position{row: 1, col: 4}: 'r',
			},
			want: false,
		},
		{
			name:        "first piece in direction not attacking piece",
			pos:         Position{row: 4, col: 4},
			attackColor: 'b',
			pieces: map[Position]rune{
				Position{row: 1, col: 4}: 'b',
			},
			want: false,
		},
		{
			name:        "empty board not attacked",
			pos:         Position{row: 4, col: 4},
			attackColor: 'b',
			pieces:      nil,
			want:        false,
		},
		{
			name:        "knight outside checked bounds safely",
			pos:         Position{row: 0, col: 0},
			attackColor: 'b',
			pieces: map[Position]rune{
				Position{row: 2, col: 1}: 'n',
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := emptyBoard()

			for pos, piece := range tt.pieces {
				b[pos.row][pos.col] = piece
			}

			got := IsSquareAttacked(b, tt.pos, tt.attackColor)

			require.Equal(t, tt.want, got)
		})
	}
}
