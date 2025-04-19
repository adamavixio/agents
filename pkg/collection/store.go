package collection

import (
	"context"
	"sync"
)

type Store[T comparable] interface {
	Put(context.Context, T) error
	Has(context.Context, T) (bool, error)
	Del(context.Context, T) error
	Tx(ctx context.Context, fn func(Transaction[T]) error) error
}

type store[T comparable] struct {
	mu  sync.RWMutex
	set Set[T]
}

func NewStore[T comparable]() Store[T] {
	return &store[T]{set: NewSet[T]()}
}

func (s *store[T]) Put(ctx context.Context, v T) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Put(v)
	return nil
}

func (s *store[T]) Has(ctx context.Context, v T) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.set.Has(v), nil
}

func (s *store[T]) Del(ctx context.Context, v T) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set.Del(v)
	return nil
}

func (s *store[T]) Tx(ctx context.Context, fn func(Transaction[T]) error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	tx := &txStore[T]{set: s.set}
	return fn(tx)
}

type Transaction[T comparable] interface {
	Put(T)
	Has(T) bool
	Del(T)
}

type txStore[T comparable] struct {
	set Set[T]
}

func (tx *txStore[T]) Put(v T)      { tx.set.Put(v) }
func (tx *txStore[T]) Has(v T) bool { return tx.set.Has(v) }
func (tx *txStore[T]) Del(v T)      { tx.set.Del(v) }
