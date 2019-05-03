// Package mock_check is a generated GoMock package.
package mock_check

import (
	x "github.com/eltorocorp/go-check"
	context "context"
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTxProvider is a mock of TxProvider interface
type MockTxProvider struct {
	ctrl     *gomock.Controller
	recorder *MockTxProviderMockRecorder
}

// MockTxProviderMockRecorder is the mock recorder for MockTxProvider
type MockTxProviderMockRecorder struct {
	mock *MockTxProvider
}

// NewMockTxProvider creates a new mock instance
func NewMockTxProvider(ctrl *gomock.Controller) *MockTxProvider {
	mock := &MockTxProvider{ctrl: ctrl}
	mock.recorder = &MockTxProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxProvider) EXPECT() *MockTxProviderMockRecorder {
	return m.recorder
}

// Begin mocks base method
func (m *MockTxProvider) Begin() (x.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(x.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin
func (mr *MockTxProviderMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockTxProvider)(nil).Begin))
}

// MockTx is a mock of Tx interface
type MockTx struct {
	ctrl     *gomock.Controller
	recorder *MockTxMockRecorder
}

// MockTxMockRecorder is the mock recorder for MockTx
type MockTxMockRecorder struct {
	mock *MockTx
}

// NewMockTx creates a new mock instance
func NewMockTx(ctrl *gomock.Controller) *MockTx {
	mock := &MockTx{ctrl: ctrl}
	mock.recorder = &MockTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTx) EXPECT() *MockTxMockRecorder {
	return m.recorder
}

// Commit mocks base method
func (m *MockTx) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockTxMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTx)(nil).Commit))
}

// Exec mocks base method
func (m *MockTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec
func (mr *MockTxMockRecorder) Exec(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockTx)(nil).Exec), varargs...)
}

// ExecContext mocks base method
func (m *MockTx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext
func (mr *MockTxMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockTx)(nil).ExecContext), varargs...)
}

// Prepare mocks base method
func (m *MockTx) Prepare(query string) (*sql.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", query)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare
func (mr *MockTxMockRecorder) Prepare(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockTx)(nil).Prepare), query)
}

// PrepareContext mocks base method
func (m *MockTx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareContext", ctx, query)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareContext indicates an expected call of PrepareContext
func (mr *MockTxMockRecorder) PrepareContext(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareContext", reflect.TypeOf((*MockTx)(nil).PrepareContext), ctx, query)
}

// Query mocks base method
func (m *MockTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockTxMockRecorder) Query(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockTx)(nil).Query), varargs...)
}

// QueryContext mocks base method
func (m *MockTx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContext", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContext indicates an expected call of QueryContext
func (mr *MockTxMockRecorder) QueryContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockTx)(nil).QueryContext), varargs...)
}

// QueryRow mocks base method
func (m *MockTx) QueryRow(query string, args ...interface{}) *sql.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow
func (mr *MockTxMockRecorder) QueryRow(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockTx)(nil).QueryRow), varargs...)
}

// QueryRowContext mocks base method
func (m *MockTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext
func (mr *MockTxMockRecorder) QueryRowContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockTx)(nil).QueryRowContext), varargs...)
}

// Rollback mocks base method
func (m *MockTx) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback
func (mr *MockTxMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTx)(nil).Rollback))
}

// Stmt mocks base method
func (m *MockTx) Stmt(stmt *sql.Stmt) *sql.Stmt {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stmt", stmt)
	ret0, _ := ret[0].(*sql.Stmt)
	return ret0
}

// Stmt indicates an expected call of Stmt
func (mr *MockTxMockRecorder) Stmt(stmt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stmt", reflect.TypeOf((*MockTx)(nil).Stmt), stmt)
}

// StmtContext mocks base method
func (m *MockTx) StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StmtContext", ctx, stmt)
	ret0, _ := ret[0].(*sql.Stmt)
	return ret0
}

// StmtContext indicates an expected call of StmtContext
func (mr *MockTxMockRecorder) StmtContext(ctx, stmt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StmtContext", reflect.TypeOf((*MockTx)(nil).StmtContext), ctx, stmt)
}
