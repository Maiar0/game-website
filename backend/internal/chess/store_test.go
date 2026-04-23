package chess

import (
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)
func TestNewGame(t *testing.T) {
	id := "MytestID"
	//check connection to db
	db, err := NewGame(id)
	require.NoError(t, err)
	require.NotNil(t, db)
	//check schema applied
	row := db.QueryRow(`
	SELECT name FROM sqlite_master 
	WHERE type='table' AND name='game_state'
	`)
	var name string
	err = row.Scan(&name)
	require.NoError(t, err)
	require.Equal(t, "game_state", name)

	defer db.Close()
}
