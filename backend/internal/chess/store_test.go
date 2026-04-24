package chess

import (
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestNewGame(t *testing.T) {
	//check connection to db
	db, _, err := NewGame()
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

func TestGetDB(t *testing.T) {
	nbd, id, _ := NewGame()
	nbd.Close()
	db, err := GetDB(id)
	require.NoError(t, err)
	require.NotNil(t, db)

}
