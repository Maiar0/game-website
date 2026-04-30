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
	row, col, err := ConvertCoordinates("e4")
	require.NoError(t, err)
	require.Equal(t, 3, row)
	require.Equal(t, 4, col)

	_, _, err = ConvertCoordinates("z9")
	require.Error(t, err)

	_, _, err = ConvertCoordinates("a10")
	require.Error(t, err)

	_, _, err = ConvertCoordinates("e")
	require.Error(t, err)
}

func TestGetPiece(t *testing.T) {
	var board Board
	board.Fill("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	piece, err := GetPiece(board, "d8")
	require.NoError(t, err)
	require.Equal(t, 'q', piece)

	_, err = GetPiece(board, "z9")
	require.Error(t, err)
}
