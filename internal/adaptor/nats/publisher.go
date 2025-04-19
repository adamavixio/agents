package event

import (
	"github.com/adamjohnston/agent/internal/port/inbound"
)

type AgentPublisher struct {
	service inbound.AgentService
}
