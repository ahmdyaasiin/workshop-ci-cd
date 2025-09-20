include .env

.PHONY: compose-up compose-down migrate-up migrate-down migrate-fresh seed

DOCKER_COMPOSE_CMD = docker compose
EXEC = docker exec -it app-$(USERNAME) ./workshop-ci-cd

FORCE ?=
FORCE_FLAG := $(if $(filter 1 true yes on,$(FORCE)),-f,)

define docker-run
	@$(EXEC) $(1) 2>/dev/null || echo "Container 'app-$(USERNAME)' is not running"
endef

compose-up:
	@$(DOCKER_COMPOSE_CMD) up --detach --build

compose-down:
	@$(DOCKER_COMPOSE_CMD) down

migrate-up:
	$(call docker-run, --migrate up)

migrate-down:
	$(call docker-run, --migrate down $(FORCE_FLAG))

migrate-fresh:
	$(call docker-run, --migrate fresh $(FORCE_FLAG))

seed:
	$(call docker-run, --seed $(FORCE_FLAG))
