LOCAL_DSN ?= postgres://127.0.0.1/const_payments_db?sslmode=disable&user=postgres&password=qwerty

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: build
build:  ## сборка бинарника API сервера
	go build -o server ./cmd/server

.PHONY: run
run: build  ## запуск собранного бинарника API сервера
	./server

.PHONY: migrate-up
migrate-up:
	@echo "Running all new local database migrations..."
	@migrate -path ./migrations -database "$(LOCAL_DSN)" up

.PHONY: migrate-down
migrate-down:
	@echo "Reverting local database to the last migration step..."
	@migrate -path ./migrations -database "$(LOCAL_DSN)" down 1

.PHONY: test
test: ## запуск юнит-тестов
	@go test -cover -covermode=count ./...

.PHONY: test-cover
test-cover: ## отобразить информацию о покрытии кода тестами
	go test -cover -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html

.DEFAULT_GOAL := run