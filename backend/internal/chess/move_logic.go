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

	if !InBounds(to) {
		return false, fmt.Errorf("to position out of bounds: %v", to)
	}

	switch piece {
	case 'p':
		switch {
		case dy == 0 && dx == 1:
			return true, nil
		case dy == 0 && dx == 2 && from.row == 6:
			return true, nil
		case Abs(dy) == 1 && dx == 1:
			return true, nil
		}

	case 'P':
		dx *= -1
		switch {
		case dy == 0 && dx == 1:
			return true, nil
		case dy == 0 && dx == 2 && from.row == 1:
			return true, nil
		case Abs(dy) == 1 && dx == 1:
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
		if Abs(dx) == Abs(dy) {
			return true, nil
		}

	case 'q', 'Q':
		if Abs(dx) == Abs(dy) || (dx == 0 && dy != 0) || (dy == 0 && dx != 0) {
			return true, nil
		}

	case 'k', 'K':
		if Abs(dx) <= 1 && Abs(dy) <= 1 {
			return true, nil
		}
		if IsCastlePattern(piece, from, to) {
			return true, nil
		}

	default:
		return false, fmt.Errorf("invalid piece: %c", piece)
	}

	return false, nil
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

func IsSquareAttacked(b Board, pos Position, turn rune) bool {
	attackColor := '0'
	if turn == 'w' {
		attackColor = 'b'
	} else {
		attackColor = 'w'
	}
	//west
	p, _, found := FindPieceInDirection(b, pos, -1, 0)
	if found && PieceColor(p) == attackColor {
		switch p {
		case 'r', 'R', 'q', 'Q':
			return true
		}
	}
	//east
	p, _, found = FindPieceInDirection(b, pos, 1, 0)
	if found && PieceColor(p) == attackColor {
		switch p {
		case 'r', 'R', 'q', 'Q':
			return true
		}
	}
	//nw
	p, _, found = FindPieceInDirection(b, pos, -1, 1)
	if found && PieceColor(p) == attackColor {
		switch p {
		case 'b', 'B', 'q', 'Q':
			return true
		}
	}
	//ne
	p, _, found = FindPieceInDirection(b, pos, 1, 1)
	if found && PieceColor(p) == attackColor {
		switch p {
		case 'b', 'B', 'q', 'Q':
			return true
		}
	}
	//knights
	for _, dir := range knightDirections {
		cur := pos
		cur.row += dir.DX
		cur.col += dir.DY

		if !InBounds(cur) {
			continue
		}

		p, _ := GetPiece(b, cur)

		if PieceColor(p) == attackColor && (p == 'n' || p == 'N') {
			return true
		}
	}

	return false
}

func IsCastlePattern(piece rune, from Position, to Position) bool {
	switch piece {
	case 'K':
		return from == (Position{row: 0, col: 4}) &&
			(to == (Position{row: 0, col: 6}) || to == (Position{row: 0, col: 2}))

	case 'k':
		return from == (Position{row: 7, col: 4}) &&
			(to == (Position{row: 7, col: 6}) || to == (Position{row: 7, col: 2}))
	}

	return false
}

func IsEnPassant(b Board, from Position, to Position, flag string) bool {
	if flag == "-" {
		return false
	}
	p, _ := GetPiece(b, from)
	if p != 'p' && p != 'P' {
		return false
	}
	target, _ := GetPiece(b, Position{row: from.row, col: to.col})
	if PieceColor(p) == PieceColor(target) {
		return false
	}
	flagPos, _ := ConvertCoordinates(flag)
	if flagPos != to {
		return false
	}
	return true
}

func IsCapture(b Board, from Position, to Position) bool {
	fp, _ := GetPiece(b, from)
	tp, _ := GetPiece(b, to)
	if fp == '.' || tp == '.' {
		return false
	}
	if PieceColor(fp) == PieceColor(tp) {
		return false
	}
	switch fp {
	case 'r', 'R', 'b', 'B', 'q', 'Q':
		if !CheckPath(b, from, to) {
			return false
		}
	case 'p', 'P':
		if Abs(from.col-to.col) != 1 {
			return false
		}
	}
	return true
}

func IsPawnForwardMove(b Board, from Position, to Position) bool {
	fp, _ := GetPiece(b, from)
	tp, _ := GetPiece(b, to)

	if from.col != to.col {
		return false
	}

	if tp != '.' {
		return false
	}

	num := Abs(from.row - to.row)

	switch fp {
	case 'p':
		if num == 1 {
			return true
		}

		if num == 2 && from.row == 6 {
			skipped := Position{row: 5, col: from.col}
			sp, _ := GetPiece(b, skipped)
			return sp == '.'
		}

	case 'P':
		if num == 1 {
			return true
		}

		if num == 2 && from.row == 1 {
			skipped := Position{row: 2, col: from.col}
			sp, _ := GetPiece(b, skipped)
			return sp == '.'
		}
	}

	return false
}
