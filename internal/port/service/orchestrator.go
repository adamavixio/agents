package service

import (
	"github.com/adamjohnston/agent/internal/port/inbound"
	"github.com/adamjohnston/agent/internal/port/outbound"
)

type Orchestrator interface {
	outbound.AgentStore
	inbound.AgentCommandHandler
	outbound.AgentEventPublisher
}
