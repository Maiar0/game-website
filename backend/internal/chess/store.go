package chess
 import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	_ "modernc.org/sqlite"
 )

 func CreateDB( dbPath string, schemaPath string)(*sql.DB, error){
	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		return nil, fmt.Errorf("create db directory : %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open sqlite db: %w", err)
	}

	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("read schema file: %w", err)
	}

	if _, err := db.Exec(string(schema)); err != nil {
		db.Close()
		return nil, fmt.Errorf("apply schema: %w", err)
	}

	return db, nil
 }