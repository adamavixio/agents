package service

import (
	"context"

	"github.com/adamjohnston/agents/internal/domain"
)

type Worker interface {
	Register(context.Context, domain.AgentID) (domain.AgentRegisteredEvent, error)
	Unregister(context.Context, domain.AgentID) (domain.AgentUnregisteredEvent, error)
}
