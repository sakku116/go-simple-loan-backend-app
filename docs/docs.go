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
        "/auth/check-token": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CheckTokenReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.BaseJSONResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.CheckTokenRespData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.BaseJSONResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.LoginRespData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/login/dev": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "login dev",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDevReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.BaseJSONResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/dto.LoginDevResp"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "data": {
                                                            "$ref": "#/definitions/dto.LoginRespData"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/refresh-token": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshTokenReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.BaseJSONResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.RefreshTokenRespData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "register new user",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.BaseJSONResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.RegisterUserRespData"
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
        "dto.BaseJSONResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "detail": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.CheckTokenReq": {
            "type": "object",
            "required": [
                "access_token"
            ],
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "dto.CheckTokenRespData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "dto.LoginDevReq": {
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
                    "description": "username or email, but swagger oauth2password need username field",
                    "type": "string"
                }
            }
        },
        "dto.LoginDevResp": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "code": {
                    "type": "integer"
                },
                "data": {},
                "detail": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.LoginReq": {
            "type": "object",
            "required": [
                "password",
                "username_or_email"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username_or_email": {
                    "type": "string"
                }
            }
        },
        "dto.LoginRespData": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.RefreshTokenReq": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.RefreshTokenRespData": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterUserReq": {
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
        "dto.RegisterUserRespData": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2Password": {
            "description": "JWT Authorization header using the Bearer scheme (add 'Bearer ' prefix).",
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "/auth/login/dev"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Loan Backend API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
