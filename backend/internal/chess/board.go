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

func PrintBoard(b [8][8]rune) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j == 0 {
				fmt.Printf("%d ", i)
			}
			fmt.Printf("%c ", b[i][j])
		}
		fmt.Println()
	}
}

// takes in e4 outputs 3, 4
func ConvertCoordinates(coord string) (Position, error) {
	chars := []rune(coord)
	if len(chars) != 2 {
		return Position{}, fmt.Errorf("invalid coordinate length: %s", coord)
	}
	col := int(chars[0] - 'a')
	row := int(chars[1] - '1')
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return Position{}, fmt.Errorf("invalid coordinate: %s", coord)
	}
	//fmt.Printf("Converted %s to row: %d, col: %d\n", coord, row, col)
	return Position{row: row, col: col}, nil
}

func GetPiece(b Board, pos Position) (rune, error) {
	if pos.row < 0 || pos.row > 7 || pos.col < 0 || pos.col > 7 {
		return '0', fmt.Errorf("invalid position: row %d, col %d", pos.row, pos.col)
	}
	return b[pos.row][pos.col], nil
}

func (b *Board) MovePiece(from Position, to Position) error {
	pieceFrom, err := GetPiece(*b, from)
	if err != nil {
		return err
	}
	if pieceFrom == '.' {
		return fmt.Errorf("no piece at from position: row %d, col %d", from.row, from.col)
	}

	pieceTo, err := GetPiece(*b, to)
	if err != nil {
		return err
	}
	if pieceTo != '.' {
		return fmt.Errorf("cannot move to position: row %d, col %d, occupied by piece: %c", to.row, to.col, pieceTo)
	}

	b[to.row][to.col] = pieceFrom
	b[from.row][from.col] = '.'

	return nil
}

func (b *Board) CapturePiece(from Position, to Position) (rune, error) {
	pieceFrom, err := GetPiece(*b, from)
	if err != nil {
		return '0', err
	}
	if pieceFrom == '.' {
		return '0', fmt.Errorf("no piece at from position: row %d, col %d", from.row, from.col)
	}

	pieceTo, err := GetPiece(*b, to)
	if err != nil {
		return '0', err
	}
	if pieceTo == '.' {
		return '0', fmt.Errorf("no piece at to position: row %d, col %d", to.row, to.col)
	}

	b[to.row][to.col] = pieceFrom
	b[from.row][from.col] = '.'

	return pieceTo, nil
}
