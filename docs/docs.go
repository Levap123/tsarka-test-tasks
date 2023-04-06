// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/email/check": {
            "post": {
                "description": "Находит все emails по паттерну  «Email:__$email»",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "summary": "find emails",
                "parameters": [
                    {
                        "description": "Input string",
                        "name": "str",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/substr/find": {
            "post": {
                "description": "Находит самую длинную подстроку без повторений",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "summary": "find substring",
                "parameters": [
                    {
                        "description": "Input string",
                        "name": "str",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/rest",
	Schemes:          []string{},
	Title:            "TSARKA test tasks API",
	Description:      "Тестовое задние на позицию Golang junior в ЦАРКА",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}