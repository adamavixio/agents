package outbound

import (
	"context"

	"github.com/adamjohnston/agents/internal/domain"
)

type AgentStore interface {
	Put(ctx context.Context, id domain.AgentID) error
	Has(ctx context.Context, id domain.AgentID) (bool, error)
	Del(ctx context.Context, id domain.AgentID) error
}

type AgentPublisher interface {
	PublishRegisterAgent(context.Context, domain.RegisterAgentCommand) error
	PublishUnregisterAgent(context.Context, domain.UnregisterAgentCommand) error
	PublishAgentRegistered(context.Context, domain.AgentRegisteredEvent) error
	PublishAgentUnregistered(context.Context, domain.AgentUnregisteredEvent) error
}
