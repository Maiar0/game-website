package main

import (
	"fmt"

	"github.com/maiar0/game-website/backend/internal/chess"
)

func main() {
	var board chess.Board
	board.Fill("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	fmt.Println("Initial Chess Board:")
	chess.PrintBoard(board)

	piece, _ := chess.GetPiece(board, "d8")
	fmt.Printf("Piece at d8: %c\n", piece)
}
