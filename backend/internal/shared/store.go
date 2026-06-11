package shared

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// CreateDB creates a new SQLite database at the specified path and
// applies the provided schema.
//
// Parameters:
//   - dbPath: Full path where the database file should be created.
//     The database must not already exist.
//   - schema: SQL schema to execute immediately after creation
//     (CREATE TABLE statements, indexes, etc.).
//
// Returns:
//   - *sql.DB: An open database connection with the schema applied.
//   - error: An error if the database already exists, the directory
//     cannot be created, the database cannot be opened, or
//     the schema fails to execute.
func CreateDB(dbPath string, schema string) (*sql.DB, error) {
	//creates DB do not want to open an existing
	if _, err := os.Stat(dbPath); err == nil {
		return nil, fmt.Errorf("database already exists: %s", dbPath)
	}

	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		return nil, fmt.Errorf("create db directory: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open sqlite db: %w", err)
	}

	if _, err := db.Exec(schema); err != nil {
		db.Close()
		return nil, fmt.Errorf("apply schema: %w", err)
	}

	return db, nil
}

// OpenSQLite opens an existing SQLite database and applies the
// standard connection and PRAGMA settings used by the application.
//
// Parameters:
//   - path: Full path to an existing SQLite database file.
//     The database must already exist; this function will
//     not create a new database.
//
// Returns:
//   - *sql.DB: An open database connection configured with:
//   - WAL journaling
//   - 5 second busy timeout
//   - Foreign key enforcement
//   - NORMAL synchronous mode
//   - A small connection pool
//   - error: An error if the database does not exist, cannot be
//     opened, PRAGMA configuration fails, or the connection
//     cannot be verified with Ping().
func OpenSQLite(path string) (*sql.DB, error) {
	//ensures we only OPEN databases not create
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("database does not exist: %w", err)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	// Let database/sql reuse a small connection pool.
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(4)

	pragmas := []string{
		`PRAGMA journal_mode = WAL;`,
		`PRAGMA busy_timeout = 5000;`,
		`PRAGMA foreign_keys = ON;`,
		`PRAGMA synchronous = NORMAL;`,
	}

	for _, q := range pragmas {
		if _, err := db.Exec(q); err != nil {
			_ = db.Close()
			return nil, fmt.Errorf("apply pragma %q: %w", q, err)
		}
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	return db, nil
}
