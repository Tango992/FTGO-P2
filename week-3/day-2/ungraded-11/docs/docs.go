// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Daniel Osvaldo Rahmanto",
            "email": "email@mail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/products": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "View all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Product"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Establish a transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Transaction data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RequestTransaction"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.TransactionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Log in an existing account",
                "parameters": [
                    {
                        "description": "Log in data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register an account",
                "parameters": [
                    {
                        "description": "Registration data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.Login": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "username": {
                    "type": "string",
                    "x-order": "0"
                },
                "password": {
                    "type": "string",
                    "x-order": "1"
                }
            }
        },
        "dto.LoginResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "x-order": "0"
                },
                "token": {
                    "type": "string",
                    "x-order": "1"
                }
            }
        },
        "dto.RegisterResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "x-order": "0"
                },
                "username": {
                    "type": "string",
                    "x-order": "1"
                },
                "deposit_amount": {
                    "type": "number",
                    "x-order": "2"
                }
            }
        },
        "dto.RegisterUser": {
            "type": "object",
            "required": [
                "deposit_amount",
                "password",
                "username"
            ],
            "properties": {
                "username": {
                    "type": "string",
                    "x-order": "0"
                },
                "password": {
                    "type": "string",
                    "x-order": "1"
                },
                "deposit_amount": {
                    "type": "number",
                    "x-order": "2"
                }
            }
        },
        "dto.RequestTransaction": {
            "type": "object",
            "required": [
                "product_id",
                "quantity"
            ],
            "properties": {
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "dto.TransactionResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "x-order": "0"
                },
                "data": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/entity.Transaction"
                        }
                    ],
                    "x-order": "1"
                }
            }
        },
        "entity.Product": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "integer",
                    "x-order": "0"
                },
                "name": {
                    "type": "string",
                    "x-order": "1"
                },
                "stock": {
                    "type": "integer",
                    "x-order": "2"
                },
                "price": {
                    "type": "number",
                    "x-order": "3"
                }
            }
        },
        "entity.Transaction": {
            "type": "object",
            "properties": {
                "transaction_id": {
                    "type": "integer",
                    "x-order": "0"
                },
                "user_id": {
                    "type": "integer",
                    "x-order": "1"
                },
                "product_id": {
                    "type": "integer",
                    "x-order": "2"
                },
                "quantity": {
                    "type": "integer",
                    "x-order": "3"
                },
                "total_amount": {
                    "type": "number",
                    "x-order": "4"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1323",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Ungraded 11-12 API",
	Description:      "Made for Ungraded Challenge 11 - Hacktiv8 FTGO",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
