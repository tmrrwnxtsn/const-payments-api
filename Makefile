# строка подключения к БД
APP_DSN ?= postgres://127.0.0.1/const_payments_db?sslmode=disable&user=postgres&password=qwerty
# корневая папка сервиса (та, в которой лежат все исходники после загрузки)
WORKING_DIRECTORY := C:/Users/f0rge/Go/src/github.com/tmrrwnxtsn/const-payments-api

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
migrate-up: ## применение миграций к БД
	echo "Running all new database migrations..."
	@migrate -path $(WORKING_DIRECTORY)/migrations -database "$(APP_DSN)" up

.PHONY: migrate-down
migrate-down: ## откат миграций БД на 1 шаг
	echo "Reverting database to the last migration step..."
	@migrate -path $(WORKING_DIRECTORY)/migrations -database "$(APP_DSN)" down 1

.PHONY: migrate-reset
migrate-reset: ## перезапустить все миграции БД
	make migrate-down
	make migrate-up

.PHONY: testdata
testdata: ## заполнить БД тестовыми данными
	make migrate-reset
	echo "Filling database with test data..."
	psql -a -f ./testdata/testdata.sql "$(APP_DSN)"

.PHONY: db-start
db-start: ## запустит БД в Docker-контейнере
	docker run --rm --name postgres -v $(WORKING_DIRECTORY)/testdata:/testdata \
		-e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=const_payments_db -d -p 5432:5432 postgres

.PHONY: db-stop
db-stop: ## остановить БД, запущенную в Docker-контейнере
	docker stop postgres

.PHONY: migrate-up-docker
migrate-up-docker: ## прменение миграций к БД, запущенной в Docker-контейнере
	echo "Running all new Docker database migrations..."
	@docker run --rm -v $(WORKING_DIRECTORY)/migrations:/migrations --network host migrate/migrate:v4.15.2 -path /migrations/ -database "$(APP_DSN)" up

.PHONY: migrate-down-docker
migrate-down-docker: ## откат миграций БД, запущенной в Docker-контейнере
	echo "Reverting Docker database to the last migration step..."
	@docker run --rm -v $(WORKING_DIRECTORY)/migrations:/migrations --network host migrate/migrate:v4.15.2 -path /migrations/ -database "$(APP_DSN)" down 1

.PHONY: migrate-reset-docker
migrate-reset-docker: ## перезапустить все миграции БД, запущенной в Docker-контейнере
	make migrate-down-docker
	make migrate-up-docker

.PHONY: testdata-docker
testdata-docker: ## заполнить БД, запущенную в Docker, тестовыми данными
	make migrate-reset-docker
	echo "Filling Docker database with test data..."
	docker exec -it postgres psql -a -f /testdata/testdata.sql "$(APP_DSN)"

.PHONY: compose-up
compose-up: ## собирает образы API и БД при необходимости и запускает контейнеры (API сервер и Postgres БД)
	docker-compose up

.PHONY: test
test: ## запуск юнит-тестов
	go test -cover -covermode=count ./...

.PHONY: test-cover
test-cover: ## отобразить информацию о покрытии кода тестами
	go test -cover -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html

.PHONY: swag-init
swag-init: ## парсинг комментариев у методов и генерация Swagger-документации
	swag init -g cmd/server/main.go

.PHONY: swag-fmt
swag-fmt: ## форматирование комментариев swag
	swag fmt -g cmd/server/main.go

.DEFAULT_GOAL := run