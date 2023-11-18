include .env

## help: Print this message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^//'


## run: run the cmd/cli application
.PHONY: run
run:
	@go run ./cmd/cli -t

## tidy: download dependencies for application
.PHONY: tidy
tidy:
	@go mod tidy
