// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "node register",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterRequest"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "dto.RegisterRequest": {
            "type": "object",
            "properties": {
                "confirm_password": {
                    "type": "string",
                    "example": "123456@mmMM.$"
                },
                "email": {
                    "type": "string",
                    "example": "mehrdad@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "123456@mmMM.$"
                },
                "phone_number": {
                    "type": "string",
                    "example": "09120246217"
                },
                "referral": {
                    "type": "string",
                    "example": "L100@1245"
                },
                "username": {
                    "type": "string",
                    "example": "mehrdad"
                },
                "wallet_id": {
                    "type": "string",
                    "example": "TZJjSqfXwTP5YgjdhLdRV1zqpDwDsgdZZE"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1320",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Golang",
	Description:      "Server API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
