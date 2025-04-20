package service

import (
	"context"
)

type Orchestrator interface {
	SubscribeRegisterAgent(context.Context) error
	SubscribeUnregisterAgent(context.Context) error
}
