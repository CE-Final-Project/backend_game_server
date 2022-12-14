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
        "contact": {
            "name": "Poomipat Chuealue",
            "url": "https://github.com/Poomipat-Ch",
            "email": "poomipat002@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/friend": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "FriendRequest with PlayerID and FriendPlayerID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "FriendRequest",
                "parameters": [
                    {
                        "description": "FriendRequest body request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.FriendInvite"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login with Username and Password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login body request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginAccountResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Create new user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Register Account",
                "parameters": [
                    {
                        "description": "Register body request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterAccount"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterAccountResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Account": {
            "type": "object",
            "required": [
                "email",
                "id",
                "player_id",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "player_id": {
                    "type": "string",
                    "maxLength": 15
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        },
        "dto.FriendInvite": {
            "type": "object",
            "required": [
                "friend_player_id",
                "player_id"
            ],
            "properties": {
                "friend_player_id": {
                    "type": "string",
                    "maxLength": 15
                },
                "player_id": {
                    "type": "string",
                    "maxLength": 15
                }
            }
        },
        "dto.LoginAccount": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        },
        "dto.LoginAccountResponse": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "account": {
                    "$ref": "#/definitions/dto.Account"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterAccount": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 320
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        },
        "dto.RegisterAccountResponse": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "account": {
                    "$ref": "#/definitions/dto.Account"
                },
                "token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "X-AUTH-TOKEN",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "API Gateway Game Server (Client)",
	Description:      "API Client Gateway microservices .",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
