.PHONY: tidy
tidy: ## check go.mod
	go mod tidy

.PHONY: build
build:  ## build the API server binary
	go build -o server -a ./cmd/server

.PHONY: run
run: build  ## run the API server binary
	./server


.DEFAULT_GOAL := run