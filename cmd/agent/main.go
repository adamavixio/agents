package main

import (
	"context"
	"log"
	"time"

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

	worker := app.NewWorker(
		transport.NewAgentPublisher(conn),
		transport.NewAgentSubscriber(conn),
	)

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if event, err := worker.Register(ctx, "agent_1"); err != nil {
			log.Println(err)
		} else {
			log.Printf("Registered: %v\n", event.AgentID)
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if event, err := worker.Register(ctx, "agent_2"); err != nil {
			log.Println(err)
		} else {
			log.Printf("Registered: %v\n", event.AgentID)
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if event, err := worker.Register(ctx, "agent_1"); err != nil {
			log.Println(err)
		} else {
			log.Printf("Registered: %v\n", event.AgentID)
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if event, err := worker.Register(ctx, "agent_2"); err != nil {
			log.Println(err)
		} else {
			log.Printf("Registered: %v\n", event.AgentID)
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if event, err := worker.Unregister(ctx, "agent_1"); err != nil {
			log.Println(err)
		} else {
			log.Printf("Unregistered: %v\n", event.AgentID)
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if event, err := worker.Unregister(ctx, "agent_2"); err != nil {
			log.Println(err)
		} else {
			log.Printf("Unregistered: %v\n", event.AgentID)
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if event, err := worker.Unregister(ctx, "agent_1"); err != nil {
			log.Println(err)
		} else {
			log.Printf("Unregistered: %v\n", event.AgentID)
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if event, err := worker.Unregister(ctx, "agent_2"); err != nil {
			log.Println(err)
		} else {
			log.Printf("Unregistered: %v\n", event.AgentID)
		}
	}
}
