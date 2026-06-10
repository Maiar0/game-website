package chess

import(
	"fmt"
)

func ValidateMove(b Board, from Position, to Position, turnColor rune) (bool, error) {
	fromP, err := GetPiece(b, from)
	if err != nil {
		return false, err
	}
	if fromP == '.' {
		return false, fmt.Errorf("No Piece at From: %v ", from)
	}
	if PieceColor(fromP) == turnColor {
		return false, fmt.Errorf("Not players Piece: %v", from)
	}
	toP, err := GetPiece(b, to)
	if err != nil {
		return false, err
	}
	if PieceColor(toP) == turnColor {
		return false, fmt.Errorf("Can not capture your own Piece: %v", toP)
	}
	validPattern, err := ValidateMovePattern(fromP, from, to)
	if err != nil {
		return false, err
	}
	if !validPattern {
		return false, fmt.Errorf("Illegal Move")
	}

	return true, nil
}