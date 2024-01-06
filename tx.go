package sql

import (
	"database/sql/driver"

	"go.opentelemetry.io/otel/trace"
)

// conn defines a tracing wrapper for driver.Tx.
type tx struct {
	tx     driver.Tx
	tracer *tracer
	span   trace.Span
}

// Commit implements driver.Tx Commit.
func (t *tx) Commit() error {
	if t.span != nil {
		defer t.span.End()
	}
	return t.tx.Commit()
}

// Rollback implements driver.Tx Rollback.
func (t *tx) Rollback() error {
	if t.span != nil {
		defer t.span.End()
	}
	return t.tx.Rollback()
}
