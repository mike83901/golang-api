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
        "/GetAddressBookByUserName": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get by username address books",
                "produces": [
                    "application/json"
                ],
                "summary": "Get by username address books",
                "operationId": "get-address-book-by-username",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.AddressBook"
                            }
                        }
                    }
                }
            }
        },
        "/getAddress": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get all address books",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all address books",
                "operationId": "get-address-book-all",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.AddressBook"
                            }
                        }
                    }
                }
            }
        },
        "/getToken": {
            "get": {
                "description": "generate a JWT token for a given username",
                "produces": [
                    "application/json"
                ],
                "summary": "Generate JWT token",
                "operationId": "generate-token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddressBook": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GoApi Example API",
	Description:      "This is a sample server for GoApi.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}