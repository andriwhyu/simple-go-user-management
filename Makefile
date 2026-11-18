MIGRATION_PATH := ./migrations
APP_VERSION := $(shell jq -r '.version' version.json)

# read DB related variables from .env
ifneq (,$(wildcard .env))
include .env
endif

# set default if .env empty
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= admin
DB_PASSWORD ?= adminpassword
DB_NAME ?= user_management
DB_SSLMODE ?= disable

DB_ADDR ?= postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

# multiplatform images
BUILDER ?=
BUILDER_FLAG :=
BUILDER_DRIVER ?= docker-container
PLATFORM := linux/amd64,linux/arm64
REGISTRY ?= ghcr.io/andriwhyu
IMAGE_NAME ?= user-management

ifneq ($(BUILDER),)
	BUILDER_FLAG = --builder $(BUILDER)
endif

# Database migration recipes
.PHONY: migrate-create
migrate-create:
	@test -n "$(MIGRATION_NAME)" || { echo "MIGRATION_NAME is required, e.g. make migrate-create MIGRATION_NAME=create_users_table"; exit 1; }
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(MIGRATION_NAME)

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database="$(DB_ADDR)" up

# N_MIGRATION: number of migration that need to be rollback. N_MIGRATION=1 undo last migration
.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database="$(DB_ADDR)" down $(N_MIGRATION)

# V_MIGRATION: Set the schema migration to version V_MIGRATION without running any migration
.PHONY: migrate-force
migrate-force:
	@test -n "$(V_MIGRATION)" || { echo "V_MIGRATION is required, e.g. make migrate-force V_MIGRATION=1"; exit 1; }
	@migrate -path=$(MIGRATION_PATH) -database="$(DB_ADDR)" force $(V_MIGRATION)

# image building and docker related recipes
.PHONY: init-builder
init-builder:
	@[ -z "$(BUILDER)" ] || \
	docker buildx inspect $(BUILDER) > /dev/null 2>&1 || \
	docker buildx create --name $(BUILDER) --driver $(BUILDER_DRIVER) --bootstrap --use

.PHONY: build-image
build-image: init-builder
	@docker buildx build $(BUILDER_FLAG) -t $(IMAGE_NAME):$(APP_VERSION) -f build/Dockerfile . --load

.PHONY: publish-multiarch
publish-multiarch: init-builder
	@docker buildx build $(BUILDER_FLAG) --platform $(PLATFORM) \
		-t $(REGISTRY)/$(IMAGE_NAME):$(APP_VERSION) \
		--provenance=true --sbom=true \
		-f build/Dockerfile . --push

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