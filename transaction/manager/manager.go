// Package manager implements a transaction.Manager interface.
package manager

import (
	"context"

	"go.uber.org/multierr"

	"github.com/avito-tech/go-transaction-manager/transaction"
	"github.com/avito-tech/go-transaction-manager/transaction/settings"
)

// Opt is a type to configure Manager.
// TODO is it necessary?
type Opt func(*Manager)

// Manager is an implementation of Manager based on storing Transaction in context.Context.
type Manager struct {
	factory transaction.TrFactory
	key     transaction.CtxKey
	log     logger
}

// New creates Manager.
func New(f transaction.TrFactory) *Manager {
	return NewWithOpts(f, settings.New())
}

// NewWithOpts creates Manager with Settings.
func NewWithOpts(f transaction.TrFactory, settings transaction.Settings) *Manager {
	// TODO implements other settings
	m := &Manager{
		factory: f,
		key:     settings.CtxKey(),
		log:     defaultLog,
	}

	return m
}

// Do processes a transaction inside a closure.
func (m *Manager) Do(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	ctx, closer, err := m.init(ctx)
	if err != nil {
		return err
	}
	// Pointer to error is required for recovery and subsequent Transaction.Rollback call.
	defer closer(ctx, &err) //nolint:errcheck // The error will be processed by the caller of Manager.Do.

	return fn(ctx)
}

type closer func(context.Context, *error) error

func (m *Manager) init(ctx context.Context) (context.Context, closer, error) {
	// TODO add propagation
	tr := transaction.TrFromCtx(ctx, m.key)

	if tr == nil {
		tr, err := m.factory()
		if err != nil {
			return nil, nil, multierr.Combine(transaction.ErrBegin, err)
		}

		return transaction.CtxWithTr(ctx, m.key, tr), newTxCommit(tr, m.log), nil
	}

	return ctx, newNilClose(), nil
}
