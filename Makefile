.PHONY: format tidy install-tools lint all

format:
	cd app && go fmt ./...

tidy:
	cd app && go mod tidy

test:
	cd app && go test ./...

gomock:
	cd app && go generate ./...

all: format tidy test
