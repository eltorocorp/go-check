package check

import "database/sql"

// SQLTxProvider provides access to sql transactions.
type SQLTxProvider struct {
	DB *sql.DB
}

// Begin starts a sql transaction.
func (s *SQLTxProvider) Begin() (Tx, error) {
	return s.DB.Begin()
}

// UseDB returns a SQLTxProvider that wraps db, enabling the database to be used
// by CheckTx.
func UseDB(db *sql.DB) *SQLTxProvider {
	return &SQLTxProvider{
		DB: db,
	}
}
