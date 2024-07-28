.PHONY: server client test cert build deploy
.DEFAULT_GOAL := build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./green-api-test ./cmd/main.go

docker:
	docker build -t green-api-test .
