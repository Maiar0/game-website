package chess

import (
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestConvertCoordinates(t *testing.T) {
	pos, err := ConvertCoordinates("e4")
	require.NoError(t, err)
	require.Equal(t, 3, pos.row)
	require.Equal(t, 4, pos.col)

	_, err = ConvertCoordinates("z9")
	require.Error(t, err)

	_, err = ConvertCoordinates("a10")
	require.Error(t, err)

	_, err = ConvertCoordinates("e")
	require.Error(t, err)
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
			require.Equal(t, tt.want, InBounds(tt.pos))
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

			require.Equal(t, tt.want, CheckPath(b, tt.from, tt.to))
		})
	}
}

func TestFindPieceInDirection(t *testing.T) {
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
		name      string
		from      Position
		drow      int
		dcol      int
		pieces    map[Position]rune
		wantPiece rune
		wantPos   Position
		wantFound bool
	}{
		{
			name:      "finds piece upward",
			from:      Position{row: 0, col: 4},
			drow:      1,
			dcol:      0,
			pieces:    map[Position]rune{Position{row: 3, col: 4}: 'r'},
			wantPiece: 'r',
			wantPos:   Position{row: 3, col: 4},
			wantFound: true,
		},
		{
			name: "finds first piece only",
			from: Position{row: 0, col: 4},
			drow: 1,
			dcol: 0,
			pieces: map[Position]rune{
				Position{row: 2, col: 4}: 'p',
				Position{row: 5, col: 4}: 'q',
			},
			wantPiece: 'p',
			wantPos:   Position{row: 2, col: 4},
			wantFound: true,
		},
		{
			name:      "finds piece horizontally",
			from:      Position{row: 4, col: 0},
			drow:      0,
			dcol:      1,
			pieces:    map[Position]rune{Position{row: 4, col: 6}: 'b'},
			wantPiece: 'b',
			wantPos:   Position{row: 4, col: 6},
			wantFound: true,
		},
		{
			name:      "finds piece diagonally",
			from:      Position{row: 1, col: 1},
			drow:      1,
			dcol:      1,
			pieces:    map[Position]rune{Position{row: 5, col: 5}: 'q'},
			wantPiece: 'q',
			wantPos:   Position{row: 5, col: 5},
			wantFound: true,
		},
		{
			name:      "does not check starting square",
			from:      Position{row: 2, col: 2},
			drow:      1,
			dcol:      0,
			pieces:    map[Position]rune{Position{row: 2, col: 2}: 'k'},
			wantPiece: '.',
			wantPos:   Position{},
			wantFound: false,
		},
		{
			name:      "returns false when no piece found",
			from:      Position{row: 3, col: 3},
			drow:      1,
			dcol:      0,
			pieces:    nil,
			wantPiece: '.',
			wantPos:   Position{},
			wantFound: false,
		},
		{
			name:      "returns false when immediately out of bounds",
			from:      Position{row: 7, col: 7},
			drow:      1,
			dcol:      0,
			pieces:    nil,
			wantPiece: '.',
			wantPos:   Position{},
			wantFound: false,
		},
		{
			name:      "finds piece going negative direction",
			from:      Position{row: 7, col: 7},
			drow:      -1,
			dcol:      -1,
			pieces:    map[Position]rune{Position{row: 4, col: 4}: 'n'},
			wantPiece: 'n',
			wantPos:   Position{row: 4, col: 4},
			wantFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := emptyBoard()

			for pos, piece := range tt.pieces {
				b[pos.row][pos.col] = piece
			}

			gotPiece, gotPos, gotFound := FindPieceInDirection(b, tt.from, tt.drow, tt.dcol)

			require.Equal(t, tt.wantPiece, gotPiece)
			require.Equal(t, tt.wantPos, gotPos)
			require.Equal(t, tt.wantFound, gotFound)
		})
	}
}
