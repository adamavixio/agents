package domain

import "time"

type AgentID string

type Agent struct {
	ID AgentID `json:"id"`
}

type RegisterAgentCommand struct {
	AgentID AgentID
}

type UnregisterAgentCommand struct {
	AgentID AgentID
}

type AgentRegisteredEvent struct {
	AgentID   AgentID
	Timestamp time.Time
}

type AgentUnregisteredEvent struct {
	AgentID   AgentID
	Timestamp time.Time
}
