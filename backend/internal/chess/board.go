package chess

import (
	"fmt"
	"strings"
)

type Board [8][8]rune

type Position struct {
	row int
	col int
}

// fen string init rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
func (b *Board) Fill(fen string) {
	pieces := strings.Split(fen, " ")[0]
	rows := strings.Split(pieces, "/")
	for i, row := range rows {
		rowN := 7 - i //adjust to ensure orientation of board is correct
		col := 0
		for _, char := range row {
			if char >= '1' && char <= '8' {
				for e := 0; e < int(char-'0'); e++ {
					b[rowN][col] = '.'
					col++
				}
			} else {
				b[rowN][col] = char
				col++
			}
		}
	}
}

func GetPiece(b Board, pos Position) (rune, error) {
	if !InBounds(pos) {
		return '0', fmt.Errorf("Position Out of Bounds: %v", pos)
	}
	return b[pos.row][pos.col], nil
}

func (b *Board) GeKing(color rune) (Position, error) {
	for i, row := range b {
		for o, p := range row {
			if (p == 'k' || p == 'K') && PieceColor(p) == color {
				return Position{row: i, col: o}, nil
			}
		}
	}
	return Position{row: -1, col: -1}, fmt.Errorf("GetKing: King not found")
}

func (b *Board) MovePiece(from Position, to Position) error {
	pieceFrom, _ := GetPiece(*b, from)
	if pieceFrom == '.' {
		return fmt.Errorf("no piece at from position: row %d, col %d", from.row, from.col)
	}

	pieceTo, _ := GetPiece(*b, to)
	if pieceTo != '.' {
		return fmt.Errorf("cannot move to position: row %d, col %d, occupied by piece: %c", to.row, to.col, pieceTo)
	}

	b[to.row][to.col] = pieceFrom
	b[from.row][from.col] = '.'

	return nil
}

// we will not actually move a piece here we will onyl capture the piece remove specified pos piece from board and return it.
func (b *Board) CapturePiece(pos Position) (rune, error) {
	pp, _ := GetPiece(*b, pos)
	if pp == '.' {
		return 0, fmt.Errorf("no piece at to position: row %d, col %d", pos.row, pos.col)
	}

	b[pos.row][pos.col] = '.'

	return pp, nil
}
