package domain

import "time"

type AgentID string

type Agent struct {
	ID AgentID `json:"id"`
}

type RegisterAgentCommand struct {
	ID AgentID
}

type UnregisterAgentCommand struct {
	ID AgentID
}

type AgentRegisteredEvent struct {
	ID        AgentID
	Timestamp time.Time
}

type AgentUnregisteredEvent struct {
	ID        AgentID
	Timestamp time.Time
}
