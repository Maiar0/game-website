package chess

import (
	"fmt"
)

type GameState struct {
	GameID         string
	Placement      string
	Turn           rune
	Castling       string
	EnPassant      string
	HalfMove       int
	FullMove       int
	CapturedPieces []rune
}

// non captures
func Move(b *Board, from Position, to Position, gs *GameState) (bool, error) {
	capturedPiece := '.'

	fromP, err := GetPiece(*b, from)
	if err != nil {
		return false, err
	}
	if fromP == '.' { //Piece Exists to be moved.
		return false, fmt.Errorf("No Piece at From: %v ", from)
	}
	if PieceColor(fromP) != gs.Turn { //Correct persons turn
		return false, fmt.Errorf("Not players Piece: %v", from)
	}
	toP, err := GetPiece(*b, to)
	if err != nil {
		return false, err
	}
	if PieceColor(toP) == PieceColor(fromP) { //can not capture your own piece
		return false, fmt.Errorf("Can not capture your own Piece: %v", toP)
	}
	validPattern, err := ValidateMovePattern(fromP, from, to)
	if err != nil {
		return false, err
	}
	if !validPattern { //checks if move is valid for piece type
		return false, fmt.Errorf("Illegal Move")
	}
	if toP == '.' && !IsEnPassant(*b, from, to, gs.EnPassant) {
		switch fromP {
		case 'p':
			if IsPawnForwardMove(*b, from, to) {
				b.MovePiece(from, to)
				if Abs(from.row-to.row) == 2 {
					gs.EnPassant, err = PositionToCoordinate(Position{row: 5, col: to.col})
					if err != nil {
						return false, err
					}
				}
			} else {
				return false, fmt.Errorf("Not valid move")
			}
		case 'P':
			if IsPawnForwardMove(*b, from, to) {
				b.MovePiece(from, to)
				if Abs(from.row-to.row) == 2 {
					gs.EnPassant, err = PositionToCoordinate(Position{row: 2, col: to.col})
				}
				if err != nil {
					return false, err
				}
			} else {
				return false, fmt.Errorf("Not valid move")
			}
		case 'n', 'N':
			b.MovePiece(from, to)
		case 'r', 'R', 'b', 'B', 'q', 'Q':
			if CheckPath(*b, from, to) {
				b.MovePiece(from, to)
			} else {
				return false, fmt.Errorf("Path blocked cant move")
			}
		case 'k', 'K':
			b.MovePiece(from, to)
		}
	} else { //TODO:: THERE IS 100% a better way to do this.
		if (toP != '.' && IsCapture(*b, from, to)) || IsEnPassant(*b, from, to, gs.EnPassant) {

		} else {
			return false, fmt.Errorf("Is not move or capture. Illegal Move.")
		}
	}
	//captures
	if (toP != '.' && IsCapture(*b, from, to)) || IsEnPassant(*b, from, to, gs.EnPassant) {
		switch fromP {
		case 'p':
			if IsEnPassant(*b, from, to, gs.EnPassant) {
				capturedPiece, err = b.CapturePiece(Position{row: from.row, col: to.col})
				if err != nil {
					return false, err
				}
			} else {
				capturedPiece, _ = b.CapturePiece(to)
			}
			b.MovePiece(from, to)
		case 'P':
			if IsEnPassant(*b, from, to, gs.EnPassant) {
				capturedPiece, err = b.CapturePiece(Position{row: from.row, col: to.col})
				if err != nil {
					return false, err
				}
			} else {
				capturedPiece, _ = b.CapturePiece(to)
			}
			b.MovePiece(from, to)
		case 'n', 'N':
			capturedPiece, err = b.CapturePiece(to)
			if err != nil {
				return false, err
			}
			b.MovePiece(from, to)
		case 'r', 'R', 'b', 'B', 'q', 'Q':
			if CheckPath(*b, from, to) {
				capturedPiece, err = b.CapturePiece(to)
				if err != nil {
					return false, err
				}
				b.MovePiece(from, to)
			} else {
				return false, fmt.Errorf("Path blocked cant move")
			}
		case 'k', 'K':
			b.MovePiece(from, to)
		}
	}
	if capturedPiece != '.' {
		gs.CapturedPieces = append(gs.CapturedPieces, capturedPiece)
	}

	if (fromP == 'p' || fromP == 'P') && Abs(from.row-to.row) == 2 { //TODO:: We are setting this twice re vist
		ep := Position{
			row: (from.row + to.row) / 2,
			col: from.col,
		}

		gs.EnPassant, err = PositionToCoordinate(ep)
		if err != nil {
			return false, err
		}
	} else {
		gs.EnPassant = "-"
	}

	return true, nil
}
