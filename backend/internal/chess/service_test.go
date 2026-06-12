package chess

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMove_NoKingLogic(t *testing.T) {
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
		name           string
		from           Position
		to             Position
		turn           rune
		enPassant      string
		pieces         map[Position]rune
		wantOK         bool
		wantErr        bool
		wantFrom       rune
		wantTo         rune
		wantEnPassant  string
		wantCaptured   []rune
		checkExtra     map[Position]rune
	}{
		{"white pawn single move", Position{1, 4}, Position{2, 4}, 'w', "-", map[Position]rune{Position{1, 4}: 'P'}, true, false, '.', 'P', "-", nil, nil},
		{"black pawn single move", Position{6, 4}, Position{5, 4}, 'b', "-", map[Position]rune{Position{6, 4}: 'p'}, true, false, '.', 'p', "-", nil, nil},
		{"white pawn double move sets en passant", Position{1, 4}, Position{3, 4}, 'w', "-", map[Position]rune{Position{1, 4}: 'P'}, true, false, '.', 'P', "e3", nil, nil},
		{"black pawn double move sets en passant", Position{6, 3}, Position{4, 3}, 'b', "-", map[Position]rune{Position{6, 3}: 'p'}, true, false, '.', 'p', "d6", nil, nil},

		{"knight normal move", Position{0, 1}, Position{2, 2}, 'w', "-", map[Position]rune{Position{0, 1}: 'N'}, true, false, '.', 'N', "-", nil, nil},
		{"rook clear move", Position{0, 0}, Position{0, 5}, 'w', "-", map[Position]rune{Position{0, 0}: 'R'}, true, false, '.', 'R', "-", nil, nil},
		{"bishop clear move", Position{0, 2}, Position{3, 5}, 'w', "-", map[Position]rune{Position{0, 2}: 'B'}, true, false, '.', 'B', "-", nil, nil},
		{"queen clear move", Position{0, 3}, Position{4, 3}, 'w', "-", map[Position]rune{Position{0, 3}: 'Q'}, true, false, '.', 'Q', "-", nil, nil},

		{"white pawn captures black piece", Position{1, 4}, Position{2, 5}, 'w', "-", map[Position]rune{Position{1, 4}: 'P', Position{2, 5}: 'p'}, true, false, '.', 'P', "-", []rune{'p'}, nil},
		{"black pawn captures white piece", Position{6, 4}, Position{5, 3}, 'b', "-", map[Position]rune{Position{6, 4}: 'p', Position{5, 3}: 'P'}, true, false, '.', 'p', "-", []rune{'P'}, nil},
		{"knight captures", Position{0, 1}, Position{2, 2}, 'w', "-", map[Position]rune{Position{0, 1}: 'N', Position{2, 2}: 'p'}, true, false, '.', 'N', "-", []rune{'p'}, nil},
		{"rook captures clear path", Position{0, 0}, Position{0, 5}, 'w', "-", map[Position]rune{Position{0, 0}: 'R', Position{0, 5}: 'p'}, true, false, '.', 'R', "-", []rune{'p'}, nil},
		{"bishop captures clear path", Position{0, 2}, Position{3, 5}, 'w', "-", map[Position]rune{Position{0, 2}: 'B', Position{3, 5}: 'p'}, true, false, '.', 'B', "-", []rune{'p'}, nil},
		{"queen captures clear path", Position{0, 3}, Position{4, 3}, 'w', "-", map[Position]rune{Position{0, 3}: 'Q', Position{4, 3}: 'p'}, true, false, '.', 'Q', "-", []rune{'p'}, nil},

		{"white en passant", Position{4, 4}, Position{5, 3}, 'w', "d6", map[Position]rune{Position{4, 4}: 'P', Position{4, 3}: 'p'}, true, false, '.', 'P', "-", []rune{'p'}, map[Position]rune{Position{4, 3}: '.'}},
		{"black en passant", Position{3, 3}, Position{2, 4}, 'b', "e3", map[Position]rune{Position{3, 3}: 'p', Position{3, 4}: 'P'}, true, false, '.', 'p', "-", []rune{'P'}, map[Position]rune{Position{3, 4}: '.'}},

		{"no piece at from", Position{1, 4}, Position{2, 4}, 'w', "-", map[Position]rune{}, false, true, '.', '.', "-", nil, nil},
		{"wrong turn", Position{1, 4}, Position{2, 4}, 'b', "-", map[Position]rune{Position{1, 4}: 'P'}, false, true, 'P', '.', "-", nil, nil},
		{"cannot capture own piece", Position{0, 0}, Position{0, 5}, 'w', "-", map[Position]rune{Position{0, 0}: 'R', Position{0, 5}: 'N'}, false, true, 'R', 'N', "-", nil, nil},
		{"invalid move pattern", Position{0, 0}, Position{1, 1}, 'w', "-", map[Position]rune{Position{0, 0}: 'R'}, false, true, 'R', '.', "-", nil, nil},
		{"rook blocked path", Position{0, 0}, Position{0, 5}, 'w', "-", map[Position]rune{Position{0, 0}: 'R', Position{0, 3}: 'p'}, false, true, 'R', '.', "-", nil, map[Position]rune{Position{0, 3}: 'p'}},
		{"bishop blocked path", Position{0, 2}, Position{3, 5}, 'w', "-", map[Position]rune{Position{0, 2}: 'B', Position{1, 3}: 'p'}, false, true, 'B', '.', "-", nil, map[Position]rune{Position{1, 3}: 'p'}},
		{"pawn forward blocked", Position{1, 4}, Position{2, 4}, 'w', "-", map[Position]rune{Position{1, 4}: 'P', Position{2, 4}: 'p'}, false, true, 'P', 'p', "-", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := emptyBoard()
			gs := GameState{Turn: tt.turn, EnPassant: tt.enPassant}

			for pos, piece := range tt.pieces {
				b[pos.row][pos.col] = piece
			}

			ok, err := Move(&b, tt.from, tt.to, &gs)

			require.Equal(t, tt.wantOK, ok)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			gotFrom, _ := GetPiece(b, tt.from)
			gotTo, _ := GetPiece(b, tt.to)

			require.Equal(t, tt.wantFrom, gotFrom)
			require.Equal(t, tt.wantTo, gotTo)

			if tt.wantEnPassant != "" {
				require.Equal(t, tt.wantEnPassant, gs.EnPassant)
			}

			require.Equal(t, tt.wantCaptured, gs.CapturedPieces)

			for pos, want := range tt.checkExtra {
				got, _ := GetPiece(b, pos)
				require.Equal(t, want, got)
			}
		})
	}
}