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
	agentID domain.AgentID,
) error {
	return a.store.Tx(ctx, func(t collection.Transaction[domain.AgentID]) error {
		if t.Has(agentID) {
			return domain.ErrorAlreadyExists
		}
		t.Put(agentID)
		return nil
	})
}

func (a agentStore) Has(
	ctx context.Context,
	agentID domain.AgentID,
) (bool, error) {
	return a.store.Has(ctx, agentID)
}

func (a agentStore) Del(
	ctx context.Context,
	agentID domain.AgentID,
) error {
	return a.store.Tx(ctx, func(t collection.Transaction[domain.AgentID]) error {
		if !t.Has(agentID) {
			return domain.ErrorNotFound
		}
		t.Del(agentID)
		return nil
	})
}
