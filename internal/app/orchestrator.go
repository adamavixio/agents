package app

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/adamjohnston/agents/internal/domain"
	"github.com/adamjohnston/agents/internal/port/inbound"
	"github.com/adamjohnston/agents/internal/port/outbound"
	"github.com/adamjohnston/agents/internal/port/service"
)

type orchestrator struct {
	store      outbound.AgentStore
	publisher  outbound.AgentPublisher
	subscriber inbound.AgentSubscriber
}

func NewOrchestrator(
	store outbound.AgentStore,
	publisher outbound.AgentPublisher,
	subscriber inbound.AgentSubscriber,
) service.Orchestrator {
	return orchestrator{
		store:      store,
		publisher:  publisher,
		subscriber: subscriber,
	}
}

func (o orchestrator) SubscribeRegisterAgent(
	ctx context.Context,
) error {
	return o.subscriber.SubscribeRegisterAgent(ctx, func(cmd domain.RegisterAgentCommand) error {
		if err := o.store.Put(ctx, cmd.AgentID); err != nil {
			if errors.Is(err, domain.ErrorAlreadyExists) {
				log.Printf("Register Agent: Agent with ID '%v' already exists", cmd.AgentID)
				return nil
			}
			return err
		}
		return o.publisher.PublishAgentRegistered(ctx, domain.AgentRegisteredEvent{
			AgentID:   cmd.AgentID,
			Timestamp: time.Now(),
		})
	})
}

func (o orchestrator) SubscribeUnregisterAgent(
	ctx context.Context,
) error {
	return o.subscriber.SubscribeUnregisterAgent(ctx, func(cmd domain.UnregisterAgentCommand) error {
		if err := o.store.Del(ctx, cmd.AgentID); err != nil {
			if errors.Is(err, domain.ErrorNotFound) {
				log.Printf("Unregister Agent: Agent with ID '%v' does not exists", cmd.AgentID)
				return nil
			}
			return err
		}
		return o.publisher.PublishAgentUnregistered(ctx, domain.AgentUnregisteredEvent{
			AgentID:   cmd.AgentID,
			Timestamp: time.Now(),
		})
	})
}
