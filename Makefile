LOCAL_DSN ?= postgres://127.0.0.1/const_payments_db?sslmode=disable&user=postgres&password=qwerty

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: build
build:  ## сборка бинарника API сервера
	go build -o server -a ./cmd/server

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


.DEFAULT_GOAL := run