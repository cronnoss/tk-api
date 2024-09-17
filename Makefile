BIN_TICKET := "./bin/ticket"
DOCKER_IMG1="ticket:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

Red='\033[0;31m'
Green='\033[0;32m'
Color_Off='\033[0m'

help:
	@echo ${Red}"Please select a subcommand"${Color_Off}
	@echo ${Green}"make run-postgres"${Color_Off}" to run postgres"
	@echo ${Green}"make create-db"${Color_Off}" to create db"
	@echo ${Green}"make migrate-up"${Color_Off}" to migrate up"
	@echo
	@echo ${Green}"make build"${Color_Off}" to build applications"
	@echo ${Green}"make run"${Color_Off}" to run ticket"
	@echo
	@echo ${Green}"make generate"${Color_Off}" to generate"
	@echo
	@echo ${Red}"Or use docker-compose:"
	@echo ${Green}"make dc"${Color_Off}" to run docker-compose"
	@echo ${Green}"make down"${Color_Off}" to stop docker-compose"
	@echo ${Green}"make destroy"${Color_Off}" to stop docker-compose and remove volumes"
	@echo
	@echo ${Green}"make test"${Color_Off}" to run unit tests"

build: build-ticket

build-ticket:
	go build -v -o $(BIN_TICKET) -ldflags "$(LDFLAGS)" ./cmd/ticket

run: build-ticket
	$(BIN_TICKET) -config ./configs/ticket_config.toml

build-img-ticket:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG1) \
		-f build/ticket/Dockerfile .

build-img: build-img-ticket

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN_TICKET) version

test: ## Runs tests
	${info Running tests...}
	go test -v -race ./... -cover -coverprofile cover.out
	go tool cover -func cover.out | grep total

bench: ## Runs benchmarks
	${info Running benchmarks...}
	go test -bench=. -benchmem ./... -run=^#

vulcheck: ## Runs vulnerability check
	${info Running vulnerability check...}
	govulncheck ./...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.61.0

lint: install-lint-deps ## Runs linters
	@echo "-- linter running"
	golangci-lint run ./...

lint_d:
	docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:v1.61.0 golangci-lint run --timeout 5m ./...

gogen: ## generate code
	${info generate code...}
	go generate ./...

run-postgres:
	docker run -d --rm --name pg -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretkey -e PGDATA=/var/lib/postgresql/data/pgdata -v psqldata:/var/lib/postgresql/data -p 5432:5432 postgres:latest

create-db:
	docker exec -it pg createdb --username=root --owner=root ticket

drop-db:
	docker exec -it pg dropdb ticket

migrate-up:
	goose -dir migrations postgres "host=localhost user=root password=secretkey dbname=ticket sslmode=disable" up

migrate-down:
	goose -dir migrations postgres "host=localhost user=root password=secretkey dbname=ticket sslmode=disable" down

migrate-status:
	goose -dir migrations postgres "host=localhost user=root password=secretkey dbname=ticket sslmode=disable" status

migrate-reset:
	goose -dir migrations postgres "host=localhost user=root password=secretkey dbname=ticket sslmode=disable" reset

dc: destroy build-img
	docker-compose -f ./deployments/docker-compose.yaml up -d

down:
	docker-compose -f ./deployments/docker-compose.yaml down

destroy:
	docker-compose -f ./deployments/docker-compose.yaml down -v

dev_up: destroy  ## Runs local environment
	docker run -d --rm --name pg -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretkey -e PGDATA=/var/lib/postgresql/data/pgdata -v psqldata:/var/lib/postgresql/data -p 5432:5432 postgres:latest
	sleep 5
	docker exec -it pg createdb --username=root --owner=root ticket && \
	goose -dir migrations postgres "host=localhost user=root password=secretkey dbname=ticket sslmode=disable" up

logs:
	docker logs -f deployments-ticket-1

.PHONY: build build-ticket
.PHONY: run
.PHONY: build-img run-img version test lint run-postgres create-db drop-db
.PHONY: migrate-up migrate-down generate install-lint-deps help