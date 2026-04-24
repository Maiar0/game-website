package shared

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestCreateDB(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "1234.db")
	//check connection to db

	schema := `
		CREATE TABLE IF NOT EXISTS test_table (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
		);
	`

	db, err := CreateDB(dbPath, schema)
	require.NoError(t, err)
	require.NotNil(t, db)
	//check schema applied
	row := db.QueryRow(`
	SELECT name FROM sqlite_master 
	WHERE type='table' AND name='test_table'
	`)
	var name string
	err = row.Scan(&name)
	require.NoError(t, err)
	require.Equal(t, "test_table", name)

	defer db.Close()
}


func TestOpenSQLite(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "1234.db")
	schema := `
		CREATE TABLE IF NOT EXISTS test_table (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
		);
	`
	cdb, _ := CreateDB(dbPath, schema)

	cdb.Close()

	db, err := OpenSQLite(dbPath)

	require.NoError(t, err)
	require.NotNil(t, db)
}