package chess

import (
	"database/sql"
	"fmt"
	"path/filepath"

	_ "embed"

	"github.com/maiar0/game-website/backend/internal/shared"
	_ "modernc.org/sqlite"
)

//go:embed chess.sql
var schema string

const baseDir = "storage/games/chess"

var dbConnections = shared.NewDBManager()

func NewGame() (*sql.DB, string, error) {
	id, err := getNewId()
	if err != nil {
		return nil, "", fmt.Errorf("error creating new game: %w", err)
	}

	db, err := shared.CreateDB(baseDir+id+".db", schema)
	if err != nil {
		return nil, "", fmt.Errorf("error creating new game: %w", err)
	}
	
	return db, id, err
}

func GetDB(id string) (*sql.DB, error) {
	db, err := dbConnections.GetDBCon(baseDir + id + ".db")
	if err != nil {
		return nil, fmt.Errorf("error getting game: %w", err)
	}
	return db, err
}

func getNewId() (string, error) {
	for i := 0; i < 10; i++ {
		id, err := shared.RandomID(9)
		if !shared.FileExists(filepath.Join(baseDir, id+".db")) {
			return id, err
		}
	}
	return "", fmt.Errorf("failed to generate unique id after 10 attempts")
}
