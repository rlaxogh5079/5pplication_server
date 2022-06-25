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
        "contact": {
            "email": "rlaxogh507906@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/article/delete": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article"
                ],
                "summary": "delete article in some atclno and email",
                "operationId": "delete-article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Atclno",
                        "name": "Atclno",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message\" : \"some-message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/article/get": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article"
                ],
                "summary": "get all articles in same long and lat",
                "operationId": "get-articles",
                "parameters": [
                    {
                        "type": "string",
                        "description": "atclNo",
                        "name": "atclNo",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "long",
                        "name": "long",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "lat",
                        "name": "lat",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/database.Article"
                        }
                    }
                }
            }
        },
        "/article/insert": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article"
                ],
                "summary": "Insert article in some database.Article",
                "operationId": "insert-article",
                "parameters": [
                    {
                        "description": "article",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.Article"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message\" : \"some-message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/image": {
            "get": {
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "image"
                ],
                "summary": "get Image in some imageNo",
                "operationId": "load-image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "imageNo",
                        "name": "imageNo",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message\" : \"some-message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/image/upload": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "image"
                ],
                "summary": "insert Image in some imageNo",
                "operationId": "upload-image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "imageNo",
                        "name": "imageNo",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message\" : \"some-message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "login in same email and hashed_password",
                "operationId": "login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "Password",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message\" : \"some-message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "signup in some database.User",
                "operationId": "signup",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message\" : \"some-message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "delete user in some email and hashed_password",
                "operationId": "delete_user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "Password",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message\" : \"some-message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update/nickname": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "update user password in some email and nickname",
                "operationId": "update_user_nickname",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Nickname",
                        "name": "Nickname",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message\" : \"some-message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update/password": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "update user password in some email and hashed_password",
                "operationId": "update_user_password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "Password",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message\" : \"some-message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "database.Article": {
            "type": "object",
            "properties": {
                "atclNo": {
                    "type": "string"
                },
                "body": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "images": {
                    "type": "string"
                },
                "latitude": {
                    "type": "string"
                },
                "likecnt": {
                    "type": "integer"
                },
                "longitude": {
                    "type": "string"
                },
                "share": {
                    "type": "boolean"
                },
                "tag": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "database.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "storeArticle": {
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
