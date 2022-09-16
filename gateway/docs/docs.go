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
        "/account": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update account: username or email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Update account",
                "parameters": [
                    {
                        "description": "Update body request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateAccountDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateAccountResponseDto"
                        }
                    }
                }
            }
        },
        "/account/password": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update password account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Change password account",
                "parameters": [
                    {
                        "description": "ChangePassword body request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ChangePasswordDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateAccountResponseDto"
                        }
                    }
                }
            }
        },
        "/account/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all information Account by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Account information by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AccountResponseDto"
                        }
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
                            "$ref": "#/definitions/dto.LoginAccountDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginAccountResponseDto"
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
                            "$ref": "#/definitions/dto.RegisterAccountDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterAccountResponseDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AccountResponseDto": {
            "type": "object",
            "required": [
                "account_id",
                "email",
                "is_ban",
                "player_id",
                "username"
            ],
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "is_ban": {
                    "type": "boolean"
                },
                "player_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.ChangePasswordDto": {
            "type": "object",
            "required": [
                "account_id",
                "new_password",
                "old_password"
            ],
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                }
            }
        },
        "dto.LoginAccountDto": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LoginAccountResponseDto": {
            "type": "object",
            "required": [
                "account_id",
                "player_id",
                "token"
            ],
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "player_id": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterAccountDto": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterAccountResponseDto": {
            "type": "object",
            "required": [
                "account_id",
                "player_id",
                "token"
            ],
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "player_id": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateAccountDto": {
            "type": "object",
            "required": [
                "account_id"
            ],
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateAccountResponseDto": {
            "type": "object",
            "required": [
                "account_id"
            ],
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "updated_at": {
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
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
