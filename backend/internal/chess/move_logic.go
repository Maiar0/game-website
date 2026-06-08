package chess

import (
	"fmt"
)

func CheckMove(piece rune, from Position, to Position) (bool, rune, error) {
	var dy = from.col - to.col
	var dx = from.row - to.row
	switch piece {
	case 'p':
		dy *= -1
		switch {
		case dy == 0 && dx == 1:
			return true, 'm', nil
		case dy == 0 && dx == 2 && from.row == 1:
			return true, 'd', nil
		case abs(dy) == 1 && dx == 1:
			return false, '0', fmt.Errorf("this is a capture: %s", from)
		}
	case 'P':
		//stuff
	case 'r':
		//stuff
	case 'R':
		//stuff
	case 'n':
		//stuff
	case 'N':
		//stuff
	case 'b':
		//stuff
	case 'B':
		//stuff
	case 'q':
		//stuff
	case 'Q':
		//stuff
	case 'k':
		//stuff
	case 'K':
		//stuff
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
