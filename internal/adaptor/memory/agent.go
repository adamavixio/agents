package memory

import (
	"context"

	"github.com/adamjohnston/agents/internal/domain"
	"github.com/adamjohnston/agents/pkg/collection"
)

type AgentStore struct {
	store collection.Store[domain.AgentID]
}

func (as AgentStore) Put(ctx context.Context, id domain.AgentID) error {
	return as.store.Tx(ctx, func(t collection.Transaction[domain.AgentID]) error {
		if t.Has(id) {
			return domain.ErrorAlreadyExists
		}
		t.Put(id)
		return nil
	})
}

func (as AgentStore) Has(ctx context.Context, id domain.AgentID) (bool, error) {
	return as.store.Has(ctx, id)
}

func (as AgentStore) Del(ctx context.Context, id domain.AgentID) error {
	return as.store.Tx(ctx, func(t collection.Transaction[domain.AgentID]) error {
		if t.Has(id) {
			return domain.ErrorAlreadyExists
		}
		t.Del(id)
		return nil
	})
}
