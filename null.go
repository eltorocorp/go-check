package check

import (
	"context"
	"database/sql"
)

type nullDB struct{}

func (nullDB) Begin() (Tx, error) {
	return nullTx{}, nil
}

type nullTx struct{}

func (nullTx) Commit() error                                              { return nil }
func (nullTx) Rollback() error                                            { return nil }
func (nullTx) Exec(query string, args ...interface{}) (sql.Result, error) { panic("n/a") }
func (nullTx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	panic("n/a")
}
func (nullTx) Prepare(query string) (*sql.Stmt, error)                             { panic("n/a") }
func (nullTx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) { panic("n/a") }
func (nullTx) Query(query string, args ...interface{}) (*sql.Rows, error)          { panic("n/a") }
func (nullTx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	panic("n/a")
}
func (nullTx) QueryRow(query string, args ...interface{}) *sql.Row { panic("n/a") }
func (nullTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	panic("n/a")
}
func (nullTx) Stmt(stmt *sql.Stmt) *sql.Stmt                             { panic("n/a") }
func (nullTx) StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt { panic("n/a") }
