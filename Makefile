.PHONY: build
build:
	go build -o bin/app cmd/main.go

.PHONY: run
run:
	go run cmd/main.go

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: precommit
update:
	go mod tidy
	go mod vendor
	go fmt ./...
