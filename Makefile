.PHONY: format tidy install-tools lint all

format:
	cd app && go fmt ./...

tidy:
	cd app && go mod tidy

test:
	cd app && go test ./...

install-tools:
	go install some/tool@latest

lint:
	cd app && golangci-lint run

all: format tidy test
