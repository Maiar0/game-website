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
	var board Board
	board.Fill("rnbqkbnr/pppppppp/8/2p5/1P6/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1")
	from, _ := ConvertCoordinates("c5")
	to, _ := ConvertCoordinates("b4")
	piece, err := board.CapturePiece(from, to)
	require.NoError(t, err)
	require.Equal(t, 'P', piece)
}
