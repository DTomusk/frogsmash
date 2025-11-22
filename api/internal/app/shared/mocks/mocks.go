package mocks

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/shared"
)

type MockDBTX struct {
	ExecContextFunc  func(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContextFunc func(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowFunc     func(ctx context.Context, query string, args ...any) *sql.Row
}

func (m *MockDBTX) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return m.ExecContextFunc(ctx, query, args...)
}

func (m *MockDBTX) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return m.QueryContextFunc(ctx, query, args...)
}

func (m *MockDBTX) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return m.QueryRowFunc(ctx, query, args...)
}

type MockTxStarter struct {
	BeginTxFunc func(ctx context.Context, opts *sql.TxOptions) (shared.Tx, error)
}

func (m *MockTxStarter) BeginTx(ctx context.Context, opts *sql.TxOptions) (shared.Tx, error) {
	return m.BeginTxFunc(ctx, opts)
}
