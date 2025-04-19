package service

import (
	"context"
	"time"

	"github.com/adamjohnston/agent/internal/port/inbound"
	"github.com/adamjohnston/agent/internal/port/outbound"
)

type AgentCommandService struct {
	handler   inbound.AgentCommandHandler
	publisher outbound.AgentEventPublisher
}

func (s *AgentCommandService) RegisterAgent(ctx context.Context, cmd inbound.RegisterAgentCmd) error {
	if err := s.handler.RegisterAgent(ctx, cmd); err != nil {
		return err
	}

	return s.publisher.PublishAgentRegistered(ctx, outbound.AgentRegisteredEvt{
		ID:        cmd.ID,
		Timestamp: time.Now(),
	})
}

func (s *AgentCommandService) UnregisterAgent(ctx context.Context, cmd inbound.UnregisterAgentCmd) error {
	if err := s.handler.UnregisterAgent(ctx, cmd); err != nil {
		return err
	}

	return s.publisher.PublishAgentUnregistered(ctx, outbound.AgentUnregisteredEvt{
		ID:        cmd.ID,
		Timestamp: time.Now(),
	})
}
