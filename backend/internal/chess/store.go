package chess

import (
	"database/sql"
	"fmt"

	_ "embed"

	"github.com/maiar0/game-website/backend/internal/shared"
	_ "modernc.org/sqlite"
)

//go:embed chess.sql
var schema string

const baseDir = "storage/games/chess"

var dbConnections = shared.NewDBManager()

func NewGame(id string) (*sql.DB, error) {
	db, err := shared.CreateDB(baseDir+id+".db", schema)
	if err != nil {
		return nil, fmt.Errorf("error creating new game: %w", err)
	}
	return db, err
}

func GetDB(id string) (*sql.DB, error) {
	db, err := dbConnections.GetDBCon(baseDir + id + ".db")
	if err != nil {
		return nil, fmt.Errorf("error getting game: %w", err)
	}
	return db, err
}
