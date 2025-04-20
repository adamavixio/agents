package inbound

import (
	"context"

	"github.com/adamjohnston/agents/internal/domain"
)

type AgentHandler interface {
	RegisterAgent(context.Context, domain.RegisterAgentCommand) error
	UnregisterAgent(context.Context, domain.UnregisterAgentCommand) error
	AgentRegistered(context.Context, domain.AgentRegisteredEvent) error
	AgentUnregistered(context.Context, domain.AgentUnregisteredEvent) error
}

type AgentSubscriber interface {
	SubscribeRegisterAgent(context.Context, func(domain.RegisterAgentCommand) error) error
	SubscribeUnregisterAgent(context.Context, func(domain.UnregisterAgentCommand) error) error
	WaitAgentRegistered(context.Context, domain.AgentID) (domain.AgentRegisteredEvent, error)
	WaitAgentUnregistered(context.Context, domain.AgentID) (domain.AgentUnregisteredEvent, error)
}
