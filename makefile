.PHONY: up down

GREEN := \033[0;32m
YELLOW := \033[0;33m
RED := \033[0;31m
NC := \033[0m 

up:
	docker compose up --build

down:
	docker compose down