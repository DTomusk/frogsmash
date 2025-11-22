package shared

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Tx interface {
	DBTX
	Commit() error
	Rollback() error
}

type TxStarter interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error)
}

type DBWithTxStarter interface {
	DBTX
	TxStarter
}

// postgresTx implements Tx
type postgresTx struct {
	tx *sql.Tx
}

func (t *postgresTx) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return t.tx.ExecContext(ctx, query, args...)
}

func (t *postgresTx) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return t.tx.QueryContext(ctx, query, args...)
}

func (t *postgresTx) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return t.tx.QueryRowContext(ctx, query, args...)
}

func (t *postgresTx) Commit() error {
	return t.tx.Commit()
}

func (t *postgresTx) Rollback() error {
	return t.tx.Rollback()
}

// postgresDB implements DBWithTxStarter
type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(db *sql.DB) *PostgresDB {
	return &PostgresDB{db: db}
}

func (p *PostgresDB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return p.db.ExecContext(ctx, query, args...)
}

func (p *PostgresDB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return p.db.QueryContext(ctx, query, args...)
}

func (p *PostgresDB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return p.db.QueryRowContext(ctx, query, args...)
}

func (p *PostgresDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	tx, err := p.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &postgresTx{tx: tx}, nil
}
