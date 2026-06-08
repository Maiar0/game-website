package chess

import (
	"fmt"
)

type Direction struct {
	DX int
	DY int
}

func ValidateMovePattern(piece rune, from Position, to Position) (bool, rune, error) {
	var dy = from.col - to.col
	var dx = from.row - to.row
	if dx == 0 && dy == 0 {
		return false, '0', fmt.Errorf("to position matchs from position: %v, %v", to, from)
	}
	if inBounds(to) {
		return false, '0', fmt.Errorf("to position out of bound: %v", to)
	}
	switch piece {
	case 'p':
		dy *= -1
		switch {
		case dy == 0 && dx == 1:
			return true, 'm', nil
		case dy == 0 && dx == 2 && from.row == 1:
			return true, 'd', nil
		case abs(dy) == 1 && dx == 1:
			return true, 'c', nil
		}
	case 'P':
		dy *= -1
		switch {
		case dy == 0 && dx == 1:
			return true, 'm', nil
		case dy == 0 && dx == 2 && from.row == 1:
			return true, 'd', nil
		case abs(dy) == 1 && dx == 1:
			return true, 'c', nil
		}
	case 'r', 'R':
		if (dx == 0 && dy != 0) || (dy == 0 && dx != 0) {
			return true, 'm', nil
		}
	case 'n', 'N':
		for _, dir := range knightDirections {
			if dx == dir.DX && dy == dir.DY {
				return true, 'm', nil
			}
		}
	case 'b', 'B':
		if abs(dx) == abs(dy) {
			return true, 'm', nil
		}
	case 'q', 'Q':
		if abs(dx) == abs(dy) {
			return true, 'm', nil
		}
		if (dx == 0 && dy != 0) || (dy == 0 && dx != 0) {
			return true, 'm', nil
		}
	case 'k', 'K':
		if abs(dx) <= 1 && abs(dy) <= 1 {
			return true, 'm', nil
		}
	default:
		return false, '0', fmt.Errorf("invalid char: %s", piece)

	}
	return false, '0', nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(n int) int {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}

var knightDirections = []Direction{
	{DX: 2, DY: 1},
	{DX: 2, DY: -1},
	{DX: -2, DY: 1},
	{DX: -2, DY: -1},
	{DX: 1, DY: 2},
	{DX: 1, DY: -2},
	{DX: -1, DY: 2},
	{DX: -1, DY: -2},
}

func inBounds(pos Position) bool {
	if pos.row > 7 || pos.row < 0 || pos.col > 7 || pos.col < 0 {
		return false
	}
	return true
}

func CheckPath(b Board, from Position, to Position) bool {
	dx := sign(from.row - to.row)
	dy := sign(from.col - to.col)
	cur := from
	for i := 0; i < 7; i++ {
		cur.row += dx
		cur.col += dy
		p, _ := GetPiece(b, cur)
		if p != '.' && cur != to {
			return false
		}
	}
	return true
}
