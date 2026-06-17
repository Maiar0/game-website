package chess

import (
	"fmt"
	"strings"
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
	dr := to.row - from.row
	dc := to.col - from.col
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
				if Abs(dr) == 2 {
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
				if Abs(dr) == 2 {
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
			if IsCastlePattern(fromP, from, to) {
				if IsSquareAttacked(*b, from, gs.Turn) ||
					IsSquareAttacked(*b, Position{row: from.row + Sign(dr), col: from.col + Sign(dc)}, gs.Turn) ||
					IsSquareAttacked(*b, to, gs.Turn) {
					return false, fmt.Errorf("Illegal Castling path is underattack")
				}
				cp, cpos, _ := FindPieceInDirection(*b, from, Sign(dr), Sign(dc))
				if fromP == 'k' && cp == 'r' {
					if strings.ContainsRune(gs.Castling, 'k') && dc > 0 {
						b.MovePiece(from, to)
						b.MovePiece(cpos, Position{row: from.row, col: (to.col - 1)})
						break

					}
					if strings.ContainsRune(gs.Castling, 'q') && dc < 0 {
						b.MovePiece(from, to)
						b.MovePiece(cpos, Position{row: from.row, col: (to.col + 1)})
						break
					}
				}
				if fromP == 'K' && cp == 'R' {
					if strings.ContainsRune(gs.Castling, 'K') && dc > 0 {
						b.MovePiece(from, to)
						b.MovePiece(cpos, Position{row: from.row, col: (to.col - 1)})
						break
					}
					if strings.ContainsRune(gs.Castling, 'Q') && dc < 0 {
						b.MovePiece(from, to)
						b.MovePiece(cpos, Position{row: from.row, col: (to.col + 1)})
						break
					}
				}
				return false, fmt.Errorf("illegal castling")
			} else {
				b.MovePiece(from, to)
			}
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
			b.CapturePiece(to)
			b.MovePiece(from, to)
		}
	}
	//TODO:: nulled tell i can work on test
	if 1 == 0 {
		k, err := GetKing(*b, gs.Turn)
		if err != nil {
			return false, err
		}
		if IsSquareAttacked(*b, k, gs.Turn) {
			return false, fmt.Errorf("King is in check, king can not end move in check: %v", k)
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
