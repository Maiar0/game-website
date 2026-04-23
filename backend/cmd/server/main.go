package main

import (
	"fmt"
	"github.com/maiar0/game-website/backend/internal/chess"
)

func main() {
	fmt.Println("hello")
	db, err := chess.NewGame("mynewgame234")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
