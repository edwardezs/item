.PHONY: build
build:
	go build -o bin/app cmd/main.go

.PHONY: run
run:
	go run cmd/main.go

.PHONY: up
up:
	docker compose up -d --build

.PHONY: down
down:
	docker compose down

.PHONY: update
update:
	go mod tidy
	go mod vendor
	go fmt ./...
	go vet ./...

.PHONY: test
test:
	go test -v ./...
