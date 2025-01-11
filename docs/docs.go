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
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "login dev",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
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
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Get user list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user list",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "query",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "username",
                            "email",
                            "nik",
                            "fullname",
                            "legalname",
                            "role"
                        ],
                        "type": "string",
                        "description": "leave empty",
                        "name": "query_by",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "updated_at",
                            "username",
                            "email",
                            "nik",
                            "fullname",
                            "legalname",
                            "role"
                        ],
                        "type": "string",
                        "default": "updated_at",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "default": "desc",
                        "x-enum-varnames": [
                            "SortOrder_asc",
                            "SortOrder_desc"
                        ],
                        "name": "sort_order",
                        "in": "query"
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
                                            "$ref": "#/definitions/dto.GetUserListRespData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user (current user)",
                "parameters": [
                    {
                        "description": "User update request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserReq"
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
                                            "$ref": "#/definitions/dto.UpdateUserRespData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create user (admin only)",
                "parameters": [
                    {
                        "description": "User create request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserReq"
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
                                            "$ref": "#/definitions/dto.CreateUserRespData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/users/face-photo": {
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Upload my Face photo (current user)",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Face photo file",
                        "name": "file",
                        "in": "formData",
                        "required": true
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
                                            "$ref": "#/definitions/dto.UploadFacePhotoRespData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/users/ktp-photo": {
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Upload my KTP photo (current user)",
                "parameters": [
                    {
                        "type": "file",
                        "description": "KTP photo file",
                        "name": "file",
                        "in": "formData",
                        "required": true
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
                                            "$ref": "#/definitions/dto.UploadKtpPhotoRespData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user (current user)",
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
                                            "$ref": "#/definitions/dto.GetUserByUUIDResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/users/{uuid}": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user by UUID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
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
                                            "$ref": "#/definitions/dto.GetUserByUUIDResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user by UUID (admin only)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User update request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserReq"
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
                                            "$ref": "#/definitions/dto.UpdateUserRespData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user by UUID (admin only)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
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
                                            "$ref": "#/definitions/dto.DeleteUserRespData"
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
        "dto.CreateUserReq": {
            "type": "object",
            "required": [
                "birthdate",
                "birthplace",
                "current_salary",
                "email",
                "fullname",
                "legalname",
                "nik",
                "password",
                "role",
                "username"
            ],
            "properties": {
                "birthdate": {
                    "description": "DD-MM-YYYY",
                    "type": "string"
                },
                "birthplace": {
                    "type": "string"
                },
                "current_salary": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "legalname": {
                    "type": "string"
                },
                "nik": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "enum": [
                        "admin",
                        "user"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/enum.UserRole"
                        }
                    ]
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.CreateUserRespData": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "type": "string"
                },
                "birthplace": {
                    "type": "string"
                },
                "current_salary": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "face_photo": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "ktp_photo": {
                    "type": "string"
                },
                "legalname": {
                    "type": "string"
                },
                "nik": {
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
        "dto.DeleteUserRespData": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "type": "string"
                },
                "birthplace": {
                    "type": "string"
                },
                "current_salary": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "face_photo": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "ktp_photo": {
                    "type": "string"
                },
                "legalname": {
                    "type": "string"
                },
                "nik": {
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
        "dto.GetUserByUUIDResp": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "type": "string"
                },
                "birthplace": {
                    "type": "string"
                },
                "current_salary": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "face_photo": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "ktp_photo": {
                    "type": "string"
                },
                "legalname": {
                    "type": "string"
                },
                "nik": {
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
        "dto.GetUserListRespData": {
            "type": "object",
            "properties": {
                "current_page": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.BaseUserResponse"
                    }
                },
                "total": {
                    "type": "integer"
                },
                "total_page": {
                    "type": "integer"
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
                "birthdate",
                "birthplace",
                "current_salary",
                "email",
                "fullname",
                "legalname",
                "nik",
                "password",
                "username"
            ],
            "properties": {
                "birthdate": {
                    "description": "DD-MM-YYYY",
                    "type": "string"
                },
                "birthplace": {
                    "type": "string"
                },
                "current_salary": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "legalname": {
                    "type": "string"
                },
                "nik": {
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
        },
        "dto.UpdateUserReq": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "description": "DD-MM-YYYY",
                    "type": "string"
                },
                "birthplace": {
                    "type": "string"
                },
                "current_salary": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "legalname": {
                    "type": "string"
                },
                "nik": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/enum.UserRole"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateUserRespData": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "type": "string"
                },
                "birthplace": {
                    "type": "string"
                },
                "current_salary": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "face_photo": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "ktp_photo": {
                    "type": "string"
                },
                "legalname": {
                    "type": "string"
                },
                "nik": {
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
        "dto.UploadFacePhotoRespData": {
            "type": "object",
            "properties": {
                "face_photo": {
                    "type": "string"
                }
            }
        },
        "dto.UploadKtpPhotoRespData": {
            "type": "object",
            "properties": {
                "ktp_photo": {
                    "type": "string"
                }
            }
        },
        "enum.SortOrder": {
            "type": "string",
            "enum": [
                "asc",
                "desc"
            ],
            "x-enum-varnames": [
                "SortOrder_asc",
                "SortOrder_desc"
            ]
        },
        "enum.UserRole": {
            "type": "string",
            "enum": [
                "user",
                "admin"
            ],
            "x-enum-varnames": [
                "UserRole_User",
                "UserRole_Admin"
            ]
        },
        "model.BaseUserResponse": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "type": "string"
                },
                "birthplace": {
                    "type": "string"
                },
                "current_salary": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "face_photo": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "ktp_photo": {
                    "type": "string"
                },
                "legalname": {
                    "type": "string"
                },
                "nik": {
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
