package memory

import (
	"context"

	"github.com/adamjohnston/agents/internal/domain"
	"github.com/adamjohnston/agents/internal/port/outbound"
	"github.com/adamjohnston/agents/pkg/collection"
)

type agentStore struct {
	store collection.Store[domain.AgentID]
}

func NewAgentStore() outbound.AgentStore {
	return agentStore{
		store: collection.NewStore[domain.AgentID](),
	}
}

func (a agentStore) Put(
	ctx context.Context,
	id domain.AgentID,
) error {
	return a.store.Tx(ctx, func(t collection.Transaction[domain.AgentID]) error {
		if t.Has(id) {
			return domain.ErrorAlreadyExists
		}
		t.Put(id)
		return nil
	})
}

func (a agentStore) Has(
	ctx context.Context,
	id domain.AgentID,
) (bool, error) {
	return a.store.Has(ctx, id)
}

func (a agentStore) Del(
	ctx context.Context,
	id domain.AgentID,
) error {
	return a.store.Tx(ctx, func(t collection.Transaction[domain.AgentID]) error {
		if !t.Has(id) {
			return domain.ErrorNotFound
		}
		t.Del(id)
		return nil
	})
}
