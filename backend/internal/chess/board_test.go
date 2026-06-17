package chess

import (
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestFill(t *testing.T) {
	var board Board
	board.Fill("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	expected := Board{
		{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'},
		{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'},
		{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'},
	}
	require.Equal(t, expected, board)
}

func TestGetPiece(t *testing.T) {
	var board Board
	board.Fill("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	pos, err := ConvertCoordinates("d8")
	require.NoError(t, err)
	piece, err := GetPiece(board, pos)
	require.NoError(t, err)
	require.Equal(t, 'q', piece)

	_, err = GetPiece(board, Position{row: 8, col: 3})
	require.Error(t, err)
	_, err = GetPiece(board, Position{row: 3, col: 8})
	require.Error(t, err)
}

func TestMovePiece(t *testing.T) {
	var board Board
	board.Fill("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	from, _ := ConvertCoordinates("e2")
	to, _ := ConvertCoordinates("e4")
	err := board.MovePiece(from, to)
	require.NoError(t, err)
	piece, _ := GetPiece(board, to)
	require.Equal(t, 'P', piece)
	piece, _ = GetPiece(board, from)
	require.Equal(t, '.', piece)
}
func TestCapturePiece(t *testing.T) {
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
		pos       Position
		pieces    map[Position]rune
		wantPiece rune
		wantErr   bool
	}{
		{
			name: "captures white pawn",
			pos:  Position{row: 3, col: 1}, // b4
			pieces: map[Position]rune{
				Position{row: 3, col: 1}: 'P',
			},
			wantPiece: 'P',
			wantErr:   false,
		},
		{
			name: "captures black pawn",
			pos:  Position{row: 4, col: 2}, // c5
			pieces: map[Position]rune{
				Position{row: 4, col: 2}: 'p',
			},
			wantPiece: 'p',
			wantErr:   false,
		},
		{
			name: "captures queen",
			pos:  Position{row: 0, col: 3}, // d1
			pieces: map[Position]rune{
				Position{row: 0, col: 3}: 'Q',
			},
			wantPiece: 'Q',
			wantErr:   false,
		},
		{
			name:      "empty square returns error",
			pos:       Position{row: 3, col: 3},
			pieces:    map[Position]rune{},
			wantPiece: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := emptyBoard()

			for pos, piece := range tt.pieces {
				b[pos.row][pos.col] = piece
			}

			gotPiece, err := b.CapturePiece(tt.pos)

			if tt.wantErr {
				require.Error(t, err)
				require.Equal(t, tt.wantPiece, gotPiece)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.wantPiece, gotPiece)

			gotAfter, err := GetPiece(b, tt.pos)
			require.NoError(t, err)
			require.Equal(t, '.', gotAfter)
		})
	}
}

func TestGetKing(t *testing.T) {
	tests := []struct {
		name    string
		fen     string
		color   rune
		wantPos Position
		wantErr bool
	}{
		{
			name:    "find white king",
			fen:     "4k3/8/8/8/8/8/8/4K3 w - - 0 1",
			color:   'w',
			wantPos: Position{row: 0, col: 4},
			wantErr: false,
		},
		{
			name:    "find black king",
			fen:     "4k3/8/8/8/8/8/8/4K3 w - - 0 1",
			color:   'b',
			wantPos: Position{row: 7, col: 4},
			wantErr: false,
		},
		{
			name:    "white king missing",
			fen:     "4k3/8/8/8/8/8/8/8 w - - 0 1",
			color:   'w',
			wantPos: Position{row: -1, col: -1},
			wantErr: true,
		},
		{
			name:    "black king missing",
			fen:     "8/8/8/8/8/8/8/4K3 w - - 0 1",
			color:   'b',
			wantPos: Position{row: -1, col: -1},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b Board
			b.Fill(tt.fen)

			got, err := b.GeKing(tt.color)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.wantPos, got)
		})
	}
}