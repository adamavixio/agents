package service

import (
	"context"

	"github.com/adamjohnston/agents/internal/domain"
	"github.com/adamjohnston/agents/internal/port/inbound"
	"github.com/adamjohnston/agents/internal/port/outbound"
	"github.com/adamjohnston/agents/internal/port/service"
)

type Worker struct {
	handler   inbound.AgentEventHandler
	publisher outbound.AgentCommandPublisher
}

var _ service.Worker = (*Worker)(nil)

func (w *Worker) Register(ctx context.Context, id domain.AgentID) error {
	return w.publisher.RegisterAgent(ctx, domain.RegisterAgentCommand{
		ID: id,
	})
}

func (w *Worker) Unregister(ctx context.Context, id domain.AgentID) error {
	return w.publisher.UnregisterAgent(ctx, domain.UnregisterAgentCommand{
		ID: id,
	})
}
