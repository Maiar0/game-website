package chess

import (
	"fmt"
)

type Direction struct {
	DX int
	DY int
}

func ValidateMovePattern(piece rune, from Position, to Position) (bool, error) {
	dy := from.col - to.col
	dx := from.row - to.row

	if dx == 0 && dy == 0 {
		return false, fmt.Errorf("to position matches from position: %v, %v", to, from)
	}

	if !inBounds(to) {
		return false, fmt.Errorf("to position out of bounds: %v", to)
	}

	switch piece {
	case 'p':
		fmt.Printf("Pawn Move %v, %v", dy, dx)
		switch {
		case dy == 0 && dx == 1:
			return true, nil
		case dy == 0 && dx == 2 && from.row == 6:
			return true, nil
		case abs(dy) == 1 && dx == 1:
			return true, nil
		}

	case 'P':
		dx *= -1
		fmt.Printf("Pawn Move %v, %v", dy, dx)
		switch {
		case dy == 0 && dx == 1:
			return true, nil
		case dy == 0 && dx == 2 && from.row == 1:
			return true, nil
		case abs(dy) == 1 && dx == 1:
			return true, nil
		}

	case 'r', 'R':
		if (dx == 0 && dy != 0) || (dy == 0 && dx != 0) {
			return true, nil
		}

	case 'n', 'N':
		for _, dir := range knightDirections {
			if dx == dir.DX && dy == dir.DY {
				return true, nil
			}
		}

	case 'b', 'B':
		if abs(dx) == abs(dy) {
			return true, nil
		}

	case 'q', 'Q':
		if abs(dx) == abs(dy) || (dx == 0 && dy != 0) || (dy == 0 && dx != 0) {
			return true, nil
		}

	case 'k', 'K':
		if abs(dx) <= 1 && abs(dy) <= 1 {
			return true, nil
		}

	default:
		return false, fmt.Errorf("invalid piece: %c", piece)
	}

	return false, nil
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
