package outbound

import (
	"context"

	"github.com/adamjohnston/agents/internal/domain"
)

type AgentStore interface {
	Put(context.Context, domain.AgentID) error
	Has(context.Context, domain.AgentID) (bool, error)
	Del(context.Context, domain.AgentID) error
}

type AgentPublisher interface {
	PublishRegisterAgent(context.Context, domain.RegisterAgentCommand) error
	PublishUnregisterAgent(context.Context, domain.UnregisterAgentCommand) error
	PublishAgentRegistered(context.Context, domain.AgentRegisteredEvent) error
	PublishAgentUnregistered(context.Context, domain.AgentUnregisteredEvent) error
}
