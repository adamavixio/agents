package service

import (
	"github.com/adamjohnston/agent/internal/port/inbound"
	"github.com/adamjohnston/agent/internal/port/outbound"
)

type Worker interface {
	inbound.AgentEventHandler
	outbound.AgentCommandPublisher
}
