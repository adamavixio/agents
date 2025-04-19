package event

import (
	"github.com/adamjohnston/agent/internal/domain"
	"github.com/adamjohnston/agent/internal/port/transport"
	"github.com/adamjohnston/agent/pkg/message"
	"github.com/nats-io/nats.go"
)

func NewAgentIdSubscriber(conn *nats.Conn, subj string) transport.AgentIDSubscriber {
	return message.NewSubscriber[domain.AgentID, struct{}](conn, subj)
}
