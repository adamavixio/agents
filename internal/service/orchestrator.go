package service

import (
	"context"
	"time"

	"github.com/adamjohnston/agents/internal/domain"
	"github.com/adamjohnston/agents/internal/port/inbound"
	"github.com/adamjohnston/agents/internal/port/outbound"
	"github.com/adamjohnston/agents/internal/port/service"
)

type Orchestrator struct {
	store     outbound.AgentStore
	handler   inbound.AgentCommandHandler
	publisher outbound.AgentEventPublisher
}

var _ service.Orchestrator = (*Orchestrator)(nil)

func (o *Orchestrator) RegisterAgent(ctx context.Context, cmd domain.RegisterAgentCommand) error {
	if err := o.handler.RegisterAgent(ctx, cmd); err != nil {
		return err
	}

	return o.publisher.AgentRegistered(ctx, domain.AgentRegisteredEvent{
		ID:        cmd.ID,
		Timestamp: time.Now(),
	})
}

func (o *Orchestrator) UnregisterAgent(ctx context.Context, cmd domain.UnregisterAgentCommand) error {
	if err := o.handler.UnregisterAgent(ctx, cmd); err != nil {
		return err
	}

	return o.publisher.AgentUnregistered(ctx, domain.AgentUnregisteredEvent{
		ID:        cmd.ID,
		Timestamp: time.Now(),
	})
}
