.PHONY: format tidy all

format:
	cd app && go fmt ./...

tidy:
	cd app && go mod tidy

all: format tidy
