package service

import (
	"context"

	"github.com/adamjohnston/agents/internal/domain"
)

type Orchestrator interface {
	RegisterAgent(context.Context, domain.RegisterAgentCommand) error
	UnregisterAgent(context.Context, domain.UnregisterAgentCommand) error
}
