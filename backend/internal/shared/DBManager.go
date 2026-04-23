package shared
import (
	"database/sql"
	"sync"

	_ "modernc.org/sqlite"
)
//in progress
type DBManager struct {
	mu sync.Mutex
	dbs map[string]*sql.DB
}
func NewDBManager() *DBManager