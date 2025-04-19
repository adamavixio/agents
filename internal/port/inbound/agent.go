package inbound

import (
	"context"

	"github.com/adamjohnston/agent/internal/domain"
)

type AgentCommandHandler interface {
	RegisterAgent(context.Context, domain.RegisterAgentCommand) error
	UnregisterAgent(context.Context, domain.UnregisterAgentCommand) error
}

type AgentEventHandler interface {
	AgentRegistered(context.Context, domain.AgentRegisteredEvent) error
	AgentUnregistered(context.Context, domain.AgentUnregisteredEvent) error
}
