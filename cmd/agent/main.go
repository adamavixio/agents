package main

func main() {
	// nc, err := nats.Connect("nats://host.docker.internal:4222")
	// if err != nil {
	// 	log.Fatalf("Error connecting to orchestrator: %v", err)
	// }
	// defer nc.Drain()

	// agent := domain.Agent{ID: "test"}
	// publisher := adaptor.NewPublisher[domain.RegisterAgentRequest, domain.RegisterAgentResponse](nc, "agents.store.*")
	// agentService := service.NewAgentService(publisher)

	// response, err := agentService.Request(&domain.RegisterAgentRequest{Agent: agent})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(response)
}
