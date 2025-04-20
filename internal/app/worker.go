package app

import (
	"context"

	"github.com/adamjohnston/agents/internal/domain"
	"github.com/adamjohnston/agents/internal/port/inbound"
	"github.com/adamjohnston/agents/internal/port/outbound"
	"github.com/adamjohnston/agents/internal/port/service"
)

type worker struct {
	publisher  outbound.AgentPublisher
	subscriber inbound.AgentSubscriber
}

func NewWorker(
	publisher outbound.AgentPublisher,
	subscriber inbound.AgentSubscriber,
) service.Worker {
	return &worker{
		publisher:  publisher,
		subscriber: subscriber,
	}
}

func (w *worker) Register(
	ctx context.Context,
	agentID domain.AgentID,
) (domain.AgentRegisteredEvent, error) {
	cmd := domain.RegisterAgentCommand{
		AgentID: agentID,
	}
	if err := w.publisher.PublishRegisterAgent(ctx, cmd); err != nil {
		return domain.AgentRegisteredEvent{}, err
	}
	return w.subscriber.WaitAgentRegistered(ctx, agentID)
}

func (w *worker) Unregister(
	ctx context.Context,
	agentID domain.AgentID,
) (domain.AgentUnregisteredEvent, error) {
	cmd := domain.UnregisterAgentCommand{
		AgentID: agentID,
	}
	if err := w.publisher.PublishUnregisterAgent(ctx, cmd); err != nil {
		return domain.AgentUnregisteredEvent{}, err
	}
	return w.subscriber.WaitAgentUnregistered(ctx, agentID)
}
