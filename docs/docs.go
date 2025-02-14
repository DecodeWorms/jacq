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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/change_password": {
            "put": {
                "description": "changes user's existing password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "changes user's existing password",
                "parameters": [
                    {
                        "description": "User request data",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ChangePassword"
                        }
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        },
                                        "token": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/change_pin": {
            "put": {
                "description": "change user's transaction pin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "change user's transaction pin",
                "parameters": [
                    {
                        "description": "User request data",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TransactionPin"
                        }
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        },
                                        "token": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/email_verification": {
            "post": {
                "description": "Send Verification email",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Send Verification email",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        },
                                        "token": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/forgot_password": {
            "put": {
                "description": "sends user forgot password 6 digits code",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "send user forgot password 6 digits code",
                "parameters": [
                    {
                        "description": "User request data",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        },
                                        "token": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User request data",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        },
                                        "token": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/secure_transaction": {
            "post": {
                "description": "Secures user` + "`" + `s transaction pin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Secures user` + "`" + `s transaction pin",
                "parameters": [
                    {
                        "description": "User request data",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        },
                                        "token": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Create a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        },
                                        "token": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/update_record": {
            "put": {
                "description": "Updates user's existing record",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Updates user's existing record",
                "parameters": [
                    {
                        "description": "User request data",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        },
                                        "token": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/verify_bvn": {
            "post": {
                "description": "verify user's bvn",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "verify user's bvn",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "User request data",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        },
                                        "token": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/verify_phone_number": {
            "post": {
                "description": "verify user's number",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "verify user's number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        },
                                        "token": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/verify_token": {
            "post": {
                "description": "verify user's token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "verify user's token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.ServerResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        },
                                        "token": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ChangePassword": {
            "type": "object",
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "current_password": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                }
            }
        },
        "model.ServerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/model.User"
                },
                "error": {},
                "status": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "model.TransactionPin": {
            "type": "object",
            "properties": {
                "confirm_new_pin": {
                    "type": "integer"
                },
                "current_pin": {
                    "type": "integer"
                },
                "new_pin": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "IDType": {
                    "type": "string"
                },
                "bvn": {
                    "type": "string"
                },
                "confirm_password": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "document": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "home_address": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "status_ts": {
                    "type": "string"
                },
                "transaction_code": {
                    "type": "integer"
                },
                "ts": {
                    "type": "string"
                }
            }
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
