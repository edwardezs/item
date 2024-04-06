.PHONY: build
build:
	go build -o bin/app cmd/main.go

.PHONY: deps
deps:

.PHONY: run
run:
	go run cmd/main.go

.PHONY: update
update:
	go mod tidy
	go fmt ./...

.PHONY: test
test:
	go test -v ./...
