package main

import (
	"fmt"

	"github.com/maiar0/game-website/backend/internal/chess"
)

func main() {
	var board chess.Board
	tt := "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1"
	tt2 := "rnbqkb1r/pppppppp/5n2/8/8/5N2/PPPPPPPP/RNBQKB1R - 0 1"
	board.Fill("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	fmt.Println("Initial Chess Board:")
	chess.PrintBoard(board)

	board.Fill(tt)
	fmt.Println("Initial Chess Board:")
	chess.PrintBoard(board)

	board.Fill(tt2)
	fmt.Println("Initial Chess Board:")
	chess.PrintBoard(board)

	piece, _ := chess.GetPiece(board, "d8")
	fmt.Printf("Piece at d8: %c\n", piece)
}
