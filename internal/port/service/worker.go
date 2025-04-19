package service

import (
	"context"

	"github.com/adamjohnston/agents/internal/domain"
)

type Worker interface {
	Register(context.Context, domain.AgentID) error
	Unregister(context.Context, domain.AgentID) error
	// OnAgentRegistered(context.Context, domain.AgentRegisteredEvent) error
	// OnAgentUnregistered(context.Context, domain.AgentUnregisteredEvent) error
}
