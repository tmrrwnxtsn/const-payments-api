// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "The MIT License (MIT)",
            "url": "https://mit-license.org/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/transactions/": {
            "get": {
                "description": "Необходимо передать либо ID, либо email пользователя, чтобы получить его платежи (транзакции).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Получить список всех платежей (транзакций) пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email пользователя",
                        "name": "user_email",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "ID пользователя",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/handler.getAllUserTransactionsResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректные данные пользователя",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка на стороне сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Платёжная система создаёт платёж (транзакцию) и уведомляет, прошёл ли он в системе.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Создать платёж (транзакцию)",
                "parameters": [
                    {
                        "description": "Информация о транзакции",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.createTransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/handler.createTransactionResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректные данные транзакции",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка на стороне сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/transactions/{id}/": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Отменяет платёж (транзакцию) по его ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID платежа (транзакции)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/handler.statusResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректный ID платежа (транзакции)",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка на стороне сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/transactions/{id}/status/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Возвращает статус платежа (транзакции) по его ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID платежа (транзакции)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/handler.getTransactionStatusResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректный ID платежа (транзакции)",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка на стороне сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Статусы \"УСПЕХ\" и \"НЕУСПЕХ\" являются терминальными - если платеж находится в них, его статус невозможно поменять.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Изменяет статус платежа (транзакции) по его ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID платежа (транзакции)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Новый статус транзакции",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.changeTransactionStatusRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/handler.statusResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректный ID платежа (транзакции) или терминальный статус платежа (транзакции)",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка на стороне сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.changeTransactionStatusRequest": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "УСПЕХ"
                }
            }
        },
        "handler.createTransactionRequest": {
            "type": "object",
            "required": [
                "amount",
                "currency_code",
                "user_email",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "type": "string",
                    "example": "123.456"
                },
                "currency_code": {
                    "type": "string",
                    "example": "RUB"
                },
                "user_email": {
                    "type": "string",
                    "example": "tmrrwnxtsn@gmail.com"
                },
                "user_id": {
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "handler.createTransactionResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "status": {
                    "type": "string",
                    "example": "УСПЕХ"
                }
            }
        },
        "handler.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "invalid transaction id"
                }
            }
        },
        "handler.getAllUserTransactionsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Transaction"
                    }
                }
            }
        },
        "handler.getTransactionStatusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "НОВЫЙ"
                }
            }
        },
        "handler.statusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "model.Transaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 123.456
                },
                "creation_time": {
                    "type": "string",
                    "example": "2022-06-07T15:25:16.046823Z"
                },
                "currency_code": {
                    "type": "string",
                    "example": "RUB"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "modified_time": {
                    "type": "string",
                    "example": "2022-06-07T15:25:16.046823Z"
                },
                "status": {
                    "type": "string",
                    "example": "НОВЫЙ"
                },
                "user_email": {
                    "type": "string",
                    "example": "tmrrwnxtsn@gmail.com"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Constanta Payments API",
	Description:      "Эмулятор платёжного сервиса, позволяющего работать с платежами (транзакциями).",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
