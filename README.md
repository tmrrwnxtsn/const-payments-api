# Эмулятор платёжного сервиса

[![Go Report Card](https://goreportcard.com/badge/github.com/tmrrwnxtsn/const-payments-api)](https://goreportcard.com/report/github.com/tmrrwnxtsn/const-payments-api)
[![codebeat badge](https://codebeat.co/badges/71e2f5e5-9e6b-405d-baf9-7cc8b5037330)](https://codebeat.co/projects/github-com-tmrrwnxtsn-const-payments-api-master)
[![Go Reference](https://pkg.go.dev/badge/github.com/tmrrwnxtsn/const-payments-api.svg)](https://pkg.go.dev/github.com/tmrrwnxtsn/const-payments-api)

Тестовое задание на стажировку в [Константу](https://const.tech/).

## Содержание

- [Задание](#Задание)
- [Подготовка](#Подготовка)
- [Запуск](#Запуск)
- [Эндпойнты](#Эндпойнты)
- [Тесты](#Тесты)
- [Структура](#Структура)
- [Зависимости](#Зависимости)

## Задание

Сервис должен принимать запросы через REST API, сохранять/изменять состояния платежей в базе данных. Код должен быть
написан на Go. В качестве базы данных, пожалуйста, используй любую реляционную.

Мы будем работать с двумя сущностями: пользователь и транзакция. Для простоты допустим, что база пользователей хранится
вне нашего сервиса. О пользователе мы можем знать только его ID (число) и email. Транзакции хранятся в нашем сервисе.
Вот что нам надо хранить о каждой транзакции

- ID транзакции (число)
- ID пользователя
- email пользователя
- сумма
- валюта
- дата и время создания
- дата и время последнего изменения
- статус

Статус может принимать одно из следующих значений: НОВЫЙ, УСПЕХ, НЕУСПЕХ, ОШИБКА. Цикл жизни платежа выглядит следующим
образом: пользователь создает платеж, он создается в статусе НОВЫЙ. После платежная система должна уведомить нас о том,
прошел ли платеж на ее стороне (п. 2 ниже), после чего мы меняем статус в нашей базе. Статусы УСПЕХ и НЕУСПЕХ являются
терминальными - если платеж находится в них, его статус должно быть невозможно поменять. Переход в статусы УСПЕХ и
НЕУСПЕХ должен осуществляться только после получения запроса из п.2 ОШИБКА - это статус, когда в момент создания платежа
что-то пошло не так. Будет хорошо, если сделаешь, чтобы случайное количество платежей при создании переходили в этот
статус.

API должно поддерживать следующие действия:

- Создание платежа (на вход принимает id пользователя, email пользователя, сумму и валюту платежа);
- Изменение статуса платежа платежной системой (хорошо, если будет к этому запросу будет применяться авторизация);
- Проверка статуса платежа по ID;
- Получение списка всех платежей пользователя по его ID;
- Получение списка всех платежей пользователя по его e-mail;
- Отмена платежа по его ID. API должно вернуть ошибку, если отмена невозможна (например потому что платеж в том статусе,
  в котором отменить нельзя).

## Подготовка

Сервис состоит из двух компонентов: API сервер и база данных PostgreSQL, поэтому для его успешной работы необходимо
установить следующее ПО:

- [Go](https://golang.org/doc/install) >=1.17;
- [PostgreSQL](https://www.postgresql.org/docs/current/tutorial-start.html) >=13 при запуске API сервера с
  использованием собственного сервера БД или [Docker](https://www.docker.com/get-started) >=20.10.14, если собственный
  сервер БД отсутствует.

## Запуск

После установки необходимого ПО необходимо скачать исходный код сервиса и перейти в директорию с исходным кодом:

```shell
git clone https://github.com/tmrrwnxtsn/const-payments-api.git
cd const-payments-api
```

Есть несколько вариантов запуска системы, все из которых происходят
благодаря [Makefile](https://github.com/tmrrwnxtsn/const-payments-api/blob/master/Makefile). **Перед запуском
рекомендуется обратить внимание на переменные, указанные в этом файле, и изменить их, если потребуется**. Также следует
обратить внимание на конфигурационный
файл [configs/local.yml](https://github.com/tmrrwnxtsn/const-payments-api/blob/master/configs/local.yml) и внести
необходимые правки. Если возникнет необходимость задать соответствующие переменные среды окружения, то они должны
добавляться с префиксом *APP_*, например, *APP_BIND_ADDR* (
см. [docker-compose.yml](https://github.com/tmrrwnxtsn/const-payments-api/blob/master/docker-compose.yml)).

### 1. [Docker Compose](https://docs.docker.com/compose/gettingstarted/)

Пожалуй, самый "безболезненный" и простой способ. Оба компонента системы (API сервер и БД) разворачиваются в отдельных
Docker-контейнерах. Настройки компонентов указываются в
[docker-compose.yml](https://github.com/tmrrwnxtsn/const-payments-api/blob/master/docker-compose.yml).

```shell
# запуск компонентов в отдельных Docker-контейнерах (без тестовых данных)
make compose-up
```

### 2. Docker

Имеется возможность разворачивать в Docker-контейнерах как API сервер, так и БД.

```shell
# запуск БД в Docker-контейнере
make db-start

# применение миграций к БД
make migrate-up-docker

# загрузка тестовых данных в БД 
make testdata-docker

# сборка образа API сервера
make build-docker

# запуск API сервера в Docker-контейнере на основе собранного образа
make run-docker
```

### 3. Local

Также есть возможность запустить API сервер локально (у себя на хосте), используя различные подключения к БД (
локальная/внешняя/Docker).

```shell
# применение миграций к БД
make migrate-up

# загрузка тестовых данных в БД 
make testdata

# компиляция и запуск бинарника с API сервером на хосте 
make run

# компиляция
make build

# запуск бинарника API сервера
make run
```

*Рекомендуется проверить настройки сервиса в соответствующих конфигурационных файлах перед запуском любым из
вышеописанных способов.*

## Эндпойнты

После успешного запуска сервиса одним из представленных способов, RESTful API сервер будет доступен по
адресу `http://localhost:8080` (если были использованы настройки по умолчанию, указанные
в [configs/local.yml](https://github.com/tmrrwnxtsn/const-payments-api/blob/master/configs/local.yml)). Сервер
поддерживает следующие эндпойнты:

* `POST /api/transactions/`: создание платежа (транзакции)
* `GET /api/transactions/`: получение списка всех платежей (транзакций) пользователя по его ID или e-mail.
* `GET /api/transactions/:id/status/`: проверка статуса платежа (транзакции) по ID
* `PATCH /api/transactions/:id/status/`: изменение статуса платежа (транзакции) системой
* `DELETE /api/transactions/:id`: отмена платежа (транзакции) по ID
* `GET /swagger/index.html`: Swagger-документация

## Тесты

Перед запуском тестов на хосте необходимо создать тестовую БД, применить к ней миграции и указать соответствующий URL
в [internal/store/sqlstore/store_test.go](https://github.com/tmrrwnxtsn/const-payments-api/blob/master/internal/store/sqlstore/store_test.go)
, либо задать в переменных среды (*APP_DSN_TEST*):

```go
// internal/store/sqlstore/store_test.go

package sqlstore_test

import (
	"github.com/tmrrwnxtsn/const-payments-api/internal/config"
	"os"
	"testing"
)

var dsn string

func TestMain(m *testing.M) {
	dsn = os.Getenv(config.EnvVariablesPrefix + "DSN_TEST")
	if dsn == "" {
		dsn = "postgres://127.0.0.1/const_payments_db_test?sslmode=disable&user=postgres&password=qwerty"
	}

	os.Exit(m.Run())
}
```

После этого можно переходить к запуску тестов:

```shell
# выполнение тестов сервиса (с информацией о покрытии в %)
make test

# получить детальную информацию о покрытии участков кода тестами (cover.out, cover.html)
make test-cover
```

## Структура

Ниже представлена структура сервиса (по папкам) с кратким описанием.

```
├── cmd                 основные приложения проекта
│   └── server          приложение API сервера
├── configs             конфигурационные файлы для различных сред развёртывания
├── docs                сгенерированная Swagger-документация 
├── internal            внутренний код приложения
│   ├── config          работа с конфигурационными данными
│   ├── handler         маршрутизация HTTP-запросов
│   ├── model           модели/сущности приложения
│   ├── server          HTTP-сервер, используемый для обработки запросов
│   ├── service         слой бизнес-логики для работы с платежами (транзакциями)
│   │   └── mocks       моки бизнес-логики работы с платежами (транзакциями)
│   └── store           слой хранения данных
│       ├── sqlstore    хранилище данных
│       └── teststore   тестовое хранилище данных (проверка логики хранения данных)
├── migrations          миграции базы данных
├── scripts             скрипты для операций над сервисом
└── testdata            тестовые данные для БД
```

Компоновка пакетов в данном проекте осуществлялась в соответствии с популярным макетом организации Go-проектов
– [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ru.md).

## Зависимости

* Маршрутизация: [gin](https://github.com/gin-gonic/gin)
* Доступ к базе данных: [sqlx](https://github.com/jmoiron/sqlx)
* Драйвер PostgreSQL: [pgx](https://github.com/jackc/pgx)
* Миграции базы данных: [golang-migrate](https://github.com/golang-migrate/migrate)
* Валидация данных: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* Логгирование: [logrus](https://github.com/sirupsen/logrus)
* Генерация Swagger-документации: [swag](https://github.com/swaggo/swag)
* Генерация моков: [golang/mock](https://github.com/golang/mock)
