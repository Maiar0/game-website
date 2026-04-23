package main

import (
	"fmt"
	"github.com/maiar0/game-website/backend/internal/chess"
)

func main() {
	fmt.Println("hello")
	db, err := chess.CreateDB("internal/chess/gamesdb/1234.db", "internal/chess/chess.sql")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
