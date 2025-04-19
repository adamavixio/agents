package outbound

import (
	"context"

	"github.com/adamjohnston/agent/internal/domain"
)

type AgentStore interface {
	Put(ctx context.Context, id domain.AgentID) error
	Has(ctx context.Context, id domain.AgentID) (bool, error)
	Del(ctx context.Context, id domain.AgentID) error
}

type AgentCommandPublisher interface {
	RegisterAgent(context.Context, domain.RegisterAgentCommand) error
	UnregisterAgent(context.Context, domain.UnregisterAgentCommand) error
}

type AgentEventPublisher interface {
	AgentRegistered(context.Context, domain.AgentRegisteredEvent) error
	AgentUnregistered(context.Context, domain.AgentUnregisteredEvent) error
}
