package shared
import (
	"database/sql"
	"sync"

	_ "modernc.org/sqlite"
)

type DBManager struct {
	mu sync.Mutex
	dbs map[string]*sql.DB
}
func NewDBManager() *DBManager {
	return &DBManager{
		dbs: make(map[string]*sql.DB),
	}
}

func ( m *DBManager )GetDBCon(path string) (*sql.DB, error){
	m.mu.Lock()
	defer m.mu.Unlock()
	//rudamentary fix for idle connection handles.
	m.sanityCheckLocked()
	
	if db, ok := m.dbs[path]; ok {
		return db, nil
	}
	 db, err := OpenSQLite(path)
	 if  err !=nil {
		return nil, err
	 }

	 m.dbs[path] = db
	 return db, nil

}

func (m *DBManager) Close( path string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	db, ok := m.dbs[path]

	if !ok {
		return nil
	}

	delete(m.dbs, path)
	return db.Close()
}

const maxOpenDBs = 500

func (m *DBManager) sanityCheckLocked() {
	if len(m.dbs) <= maxOpenDBs {
		return
	}

	for key, db := range m.dbs {
		_ = db.Close()
		delete(m.dbs, key)
	}
}