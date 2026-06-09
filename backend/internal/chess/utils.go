package chess

import (
	"fmt"
)

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
	fmt.Printf("Converted %s to row: %d, col: %d\n", coord, row, col)
	return Position{row: row, col: col}, nil
}

func InBounds(pos Position) bool {
	if pos.row > 7 || pos.row < 0 || pos.col > 7 || pos.col < 0 {
		return false
	}
	return true
}

func CheckPath(b Board, from Position, to Position) bool {
	dx := Sign(to.row - from.row)
	dy := Sign(to.col - from.col)

	cur := from

	for cur != to {
		cur.row += dx
		cur.col += dy

		p, _ := GetPiece(b, cur)

		if p != '.' && cur != to {
			return false
		}
	}

	return true
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(n int) int {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}
