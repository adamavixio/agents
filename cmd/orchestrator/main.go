package main

func main() {
	// nc, err := nats.Connect("nats://host.docker.internal:4222")
	// if err != nil {
	// 	log.Fatalf("failed to connect to NATS: %v", err)
	// }
	// defer nc.Drain()

	// store := adaptor.NewMemorySet[domain.Agent]()
	// subscriber := adaptor.NewSubscriber[domain.RegisterAgentRequest, domain.RegisterAgentResponse](nc, "agents.store.*")
	// orchestratorService := service.NewOrchestratorService(store, subscriber)

	// handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	// 	Level:     slog.LevelDebug,
	// 	AddSource: true,
	// })

	// orchestratorService.SubscribeWithMiddleware(func(request *domain.RegisterAgentRequest) (*domain.RegisterAgentResponse, error) {
	// 	if err := orchestratorService.Put(request.Agent); err != nil {
	// 		return nil, err
	// 	}
	// 	return &domain.RegisterAgentResponse{}, nil
	// }, adaptor.LoggerMiddleware[domain.RegisterAgentRequest, domain.RegisterAgentResponse](slog.New(handler)))

	// select {}
}
