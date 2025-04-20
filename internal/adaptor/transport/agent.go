package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/adamjohnston/agents/internal/domain"
	"github.com/adamjohnston/agents/internal/port/inbound"
	"github.com/adamjohnston/agents/internal/port/outbound"
	"github.com/nats-io/nats.go"
)

type agentPublisher struct {
	conn *nats.Conn
}

func NewAgentPublisher(
	conn *nats.Conn,
) outbound.AgentPublisher {
	return agentPublisher{
		conn: conn,
	}
}

func (a agentPublisher) PublishRegisterAgent(ctx context.Context, cmd domain.RegisterAgentCommand) error {
	data, err := json.Marshal(&cmd)
	if err != nil {
		return err
	}
	subj := fmt.Sprintf("command.agent.register.%v", cmd.AgentID)
	return a.conn.Publish(subj, data)
}

func (a agentPublisher) PublishUnregisterAgent(ctx context.Context, cmd domain.UnregisterAgentCommand) error {
	data, err := json.Marshal(&cmd)
	if err != nil {
		return err
	}
	subj := fmt.Sprintf("command.agent.unregister.%v", cmd.AgentID)
	return a.conn.Publish(subj, data)
}

func (a agentPublisher) PublishAgentRegistered(ctx context.Context, evt domain.AgentRegisteredEvent) error {
	data, err := json.Marshal(&evt)
	if err != nil {
		return err
	}
	subject := fmt.Sprintf("event.agent.registered.%v", evt.AgentID)
	return a.conn.Publish(subject, data)
}

func (a agentPublisher) PublishAgentUnregistered(ctx context.Context, evt domain.AgentUnregisteredEvent) error {
	data, err := json.Marshal(&evt)
	if err != nil {
		return err
	}
	subj := fmt.Sprintf("event.agent.unregistered.%v", evt.AgentID)
	return a.conn.Publish(subj, data)
}

type agentSubscriber struct {
	conn *nats.Conn
}

func NewAgentSubscriber(
	conn *nats.Conn,
) inbound.AgentSubscriber {
	return agentSubscriber{
		conn: conn,
	}
}

func (a agentSubscriber) SubscribeRegisterAgent(ctx context.Context, handler func(domain.RegisterAgentCommand) error) error {
	sub, err := a.conn.Subscribe("command.agent.register.*", func(msg *nats.Msg) {
		var cmd domain.RegisterAgentCommand
		if err := json.Unmarshal(msg.Data, &cmd); err != nil {
			log.Print(err)
			return
		}

		if err := handler(cmd); err != nil {
			log.Print(err)
			return
		}
	})

	if err != nil {
		sub.Unsubscribe()
		return err
	}

	go func() {
		defer sub.Unsubscribe()
		<-ctx.Done()
	}()

	return nil
}

func (a agentSubscriber) SubscribeUnregisterAgent(ctx context.Context, handler func(domain.UnregisterAgentCommand) error) error {
	sub, err := a.conn.Subscribe("command.agent.unregister.*", func(msg *nats.Msg) {
		var cmd domain.UnregisterAgentCommand
		if err := json.Unmarshal(msg.Data, &cmd); err != nil {
			log.Print(err)
			return
		}

		if err := handler(cmd); err != nil {
			log.Print(err)
			return
		}
	})

	if err != nil {
		sub.Unsubscribe()
		return err
	}

	go func() {
		defer sub.Unsubscribe()
		<-ctx.Done()
	}()

	return nil
}

func (a agentSubscriber) WaitAgentRegistered(ctx context.Context, AgentID domain.AgentID) (domain.AgentRegisteredEvent, error) {
	subj := fmt.Sprintf("event.agent.registered.%v", AgentID)

	sub, err := a.conn.SubscribeSync(subj)
	if err != nil {
		return domain.AgentRegisteredEvent{}, err
	}
	defer sub.Unsubscribe()

	msg, err := sub.NextMsgWithContext(ctx)
	if err != nil {
		return domain.AgentRegisteredEvent{}, err
	}

	var evt domain.AgentRegisteredEvent
	if err := json.Unmarshal(msg.Data, &evt); err != nil {
		return domain.AgentRegisteredEvent{}, err
	}

	return evt, nil
}

func (a agentSubscriber) WaitAgentUnregistered(ctx context.Context, AgentID domain.AgentID) (domain.AgentUnregisteredEvent, error) {
	subj := fmt.Sprintf("event.agent.unregistered.%v", AgentID)

	sub, err := a.conn.SubscribeSync(subj)
	if err != nil {
		return domain.AgentUnregisteredEvent{}, err
	}
	defer sub.Unsubscribe()

	msg, err := sub.NextMsgWithContext(ctx)
	if err != nil {
		return domain.AgentUnregisteredEvent{}, err
	}

	var evt domain.AgentUnregisteredEvent
	if err := json.Unmarshal(msg.Data, &evt); err != nil {
		return domain.AgentUnregisteredEvent{}, err
	}

	return evt, nil
}
