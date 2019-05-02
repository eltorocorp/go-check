package check

import (
	"context"
	"database/sql"
)

// TxProvider is anything that can initialize a transaction. This is typically
// a *sql.DB wrapped in a SQLTxProvider (see function check.UseDB)
type TxProvider interface {
	Begin() (Tx, error)
}

// Tx is anything that can perform transactional operations. This is typically
// a *sql.Tx.
type Tx interface {
	Commit() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Rollback() error
	Stmt(stmt *sql.Stmt) *sql.Stmt
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
}
