package main

import (
	"fmt"

	"github.com/maiar0/game-website/backend/internal/chess"
)

func main() {
	var board chess.Board
	tt := "rnbqkbnr/pppppppp/8/2n5/3P4/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1"

	board.Fill(tt)
	fmt.Println("tt:")
	chess.PrintBoard(board)

	pos, _ := chess.ConvertCoordinates("c5")
	piece, _ := chess.GetPiece(board, pos)
	fmt.Printf("Piece at: %c\n", piece)
}
