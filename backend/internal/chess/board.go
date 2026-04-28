package chess

import (
	"fmt"
	"strings"
)

type Board [8][8]rune

// fen string init rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
func (b *Board) Fill(fen string) {
	pieces := strings.Split(fen, " ")[0]
	rows := strings.Split(pieces, "/")
	for i, row := range rows {
		rowN := 7 - i //adjust to ensure orientation of board is correct
		for o, char := range row {
			if char >= '1' && char <= '8' {
				for e := 0; e < int(char-'0'); e++ {
					b[rowN][o] = '.'
					o++
				}
			} else {
				b[rowN][o] = char
				o++
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
func ConvertCoordinates(coord string) (int, int, error) {
	chars := []rune(coord)
	if len(chars) != 2 {
		return 0, 0, fmt.Errorf("invalid coordinate length: %s", coord)
	}
	col := int(chars[0] - 'a')
	row := int(chars[1] - '1')
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return 0, 0, fmt.Errorf("invalid coordinate: %s", coord)
	}
	fmt.Printf("Converted %s to row: %d, col: %d\n", coord, row, col)
	return row, col, nil
}

func GetPiece(b Board, coord string) (rune, error) {
	row, col, err := ConvertCoordinates(coord)
	if err != nil {
		return '.', err
	}
	return b[row][col], nil
}
