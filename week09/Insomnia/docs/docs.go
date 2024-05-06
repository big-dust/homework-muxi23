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
        "/api/v1/auth/changeAvatar": {
            "post": {
                "description": "更改头像接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "更改头像接口",
                "parameters": [
                    {
                        "description": "新头像",
                        "name": "newAvatar",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "jwt验证",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "头像更改成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/changePassword": {
            "post": {
                "description": "更改密码接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "更改密码接口",
                "parameters": [
                    {
                        "description": "邮箱",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "验证码",
                        "name": "verificationCode",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "jwt验证",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "新密码",
                        "name": "newPassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "密码更改成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/getMyData": {
            "post": {
                "description": "获取数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "获取数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt验证",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取数据成功",
                        "schema": {
                            "$ref": "#/definitions/response.GetMyDataResponse"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/login": {
            "post": {
                "description": "用户登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "邮箱",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功",
                        "schema": {
                            "$ref": "#/definitions/response.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/signup": {
            "post": {
                "description": "用户注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "用户注册接口",
                "parameters": [
                    {
                        "description": "邮箱",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "验证码",
                        "name": "verificationCode",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "头像",
                        "name": "avatar",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功",
                        "schema": {
                            "$ref": "#/definitions/response.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/common/sendEmail": {
            "post": {
                "description": "发送验证码接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SendEmail"
                ],
                "summary": "发送验证码接口",
                "parameters": [
                    {
                        "description": "邮箱",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SendEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "发送邮件成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "邮箱服务器出错",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "发送邮件失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/thread/createThread": {
            "post": {
                "description": "用户创建帖子接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "用户创建帖子接口",
                "parameters": [
                    {
                        "description": "标题",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "主题",
                        "name": "topic",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "内容",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "图片链接列表",
                        "name": "images",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    {
                        "type": "string",
                        "description": "jwt验证",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "帖子创建成功",
                        "schema": {
                            "$ref": "#/definitions/response.GetThreadResponse"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/thread/destroyThread": {
            "post": {
                "description": "用户删除帖子接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "删除帖子接口",
                "parameters": [
                    {
                        "description": "帖子唯一标识",
                        "name": "tUuid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "jwt验证",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除帖子成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/thread/getMyThreads": {
            "post": {
                "description": "获取指定用户的帖子列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "获取用户的帖子列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt验证",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取帖子成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.GetThreadResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/thread/getThreads": {
            "post": {
                "description": "获取指定主题的帖子列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "获取帖子列表接口",
                "parameters": [
                    {
                        "description": "主题",
                        "name": "topic",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "jwt验证",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取帖子成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.GetThreadResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/thread/likeThread": {
            "post": {
                "description": "用户点赞或取消点赞帖子接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "点赞/取消点赞帖子接口",
                "parameters": [
                    {
                        "description": "这里对应的就是tUuid,但是方便你复制粘贴,帖子唯一标识",
                        "name": "uid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "jwt验证",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "点赞状态切换成功",
                        "schema": {
                            "$ref": "#/definitions/response.LikesResponse"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/thread/readThread": {
            "post": {
                "description": "获取指定帖子的详情信息接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "获取帖子详情接口",
                "parameters": [
                    {
                        "description": "帖子唯一标识",
                        "name": "tUuid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "jwt验证",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取帖子成功",
                        "schema": {
                            "$ref": "#/definitions/response.GetThreadResponse"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.SendEmailRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.GetMyDataResponse": {
            "type": "object",
            "properties": {
                "GetPost": {
                    "type": "integer"
                },
                "likes": {
                    "type": "integer"
                },
                "myPost": {
                    "type": "integer"
                }
            }
        },
        "response.GetThreadResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "exist": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "likes": {
                    "type": "integer"
                },
                "postNumber": {
                    "type": "integer"
                },
                "tUuid": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "topic": {
                    "type": "string"
                },
                "uuId": {
                    "type": "string"
                }
            }
        },
        "response.LikesResponse": {
            "type": "object",
            "properties": {
                "exist": {
                    "type": "string"
                }
            }
        },
        "response.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "不眠夜",
	Description:      "一个匿名熬夜论坛",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
