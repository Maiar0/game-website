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
			PrintBoard(b)

			require.Equal(t, tt.want, CheckPath(b, tt.from, tt.to))
		})
	}
}
