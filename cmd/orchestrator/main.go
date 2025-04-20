package main

import (
	"context"
	"log"

	"github.com/adamjohnston/agents/internal/adaptor/memory"
	"github.com/adamjohnston/agents/internal/adaptor/transport"
	"github.com/adamjohnston/agents/internal/app"
	"github.com/nats-io/nats.go"
)

func main() {
	conn, err := nats.Connect("nats://host.docker.internal:4222")
	if err != nil {
		log.Fatalf("Error connecting to orchestrator: %v", err)
	}
	defer conn.Drain()

	orchestrator := app.NewOrchestrator(
		memory.NewAgentStore(),
		transport.NewAgentPublisher(conn),
		transport.NewAgentSubscriber(conn),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := orchestrator.SubscribeRegisterAgent(ctx); err != nil {
		log.Fatalf("Error Registering Agent: %v", err)
	}

	if err := orchestrator.SubscribeUnregisterAgent(ctx); err != nil {
		log.Fatalf("Error Unregistering Agent: %v", err)
	}

	select {}
}
