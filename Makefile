MIGRATION_PATH := ./migrations
APP_VERSION := $(shell jq -r '.version' version.json)

# Database migration recipes
.PHONY: migrate-create
migrate-create:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-force
migrate-force:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) force $(filter-out $@,$(MAKECMDGOALS))

# image building and docker related recipes
.PHONY: build-image
build-image:
	@docker buildx build -t user-management:$(APP_VERSION) -f build/Dockerfile . --load

.PHONY: env-docker
env-docker:
	@echo "APP_VERSION=$(APP_VERSION)"
	@echo "APP_VERSION=$(APP_VERSION)" > ./build/.env

.PHONY: build-db
build-db:
	@cd build && docker-compose up -d

.PHONY: local-server
local-server:
    # ./bin is use by the air tool
	@mkdir -p ./bin
	@air

.PHONY: run-local
run-local: build-db local-server

.PHONY: run-docker
run-docker: env-docker
	@cd build && docker-compose --profile app up -d

.PHONY: clean
clean:
	@cd build && docker-compose down
	@rm -f ./build/.env