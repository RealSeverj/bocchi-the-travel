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
        "/bocchi/access_token/get": {
            "get": {
                "description": "get available access-token by refresh-token",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get_access-token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/apply": {
            "get": {
                "description": "apply to join the party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "apply_party",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "活动id",
                        "name": "party_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/apply/list": {
            "get": {
                "description": "show party's applicants",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "apply_list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "活动id",
                        "name": "party_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/apply/permit": {
            "get": {
                "description": "permit user to join the party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "permit_join",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "活动id",
                        "name": "party_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "member_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/create": {
            "post": {
                "description": "create party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create_party",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标题",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "介绍",
                        "name": "content",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "类型",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "活动省份",
                        "name": "province",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "活动城市",
                        "name": "city",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "开始时间(例:2006-01-02)",
                        "name": "start_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "结束时间(例:2006-01-02)",
                        "name": "end_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/members": {
            "get": {
                "description": "get members who have join the party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get_party_members",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "活动id",
                        "name": "party_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/search": {
            "post": {
                "description": "search party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "search_party",
                "parameters": [
                    {
                        "type": "string",
                        "description": "搜索内容",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "活动类型",
                        "name": "party_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "活动省份",
                        "name": "province",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "活动城市",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "几天后开始",
                        "name": "start_time_duration",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "搜索方式(0:默认;1:按开始时间;2:按活动已有人数)",
                        "name": "search_type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/avatar/upload": {
            "put": {
                "description": "revise user's avatar",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "PutAvatar",
                "parameters": [
                    {
                        "type": "file",
                        "description": "头像",
                        "name": "avatar_file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/info": {
            "get": {
                "description": "get user's info",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/login/": {
            "post": {
                "description": "login to get your auth token",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "otp",
                        "name": "otp",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/register/": {
            "post": {
                "description": "userRegister",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/signature": {
            "post": {
                "description": "revise signature",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "signature",
                "parameters": [
                    {
                        "type": "string",
                        "description": "signature",
                        "name": "signature",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/switch2fa": {
            "post": {
                "description": "switch on/off 2fa mode",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "switch_2fa",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "关闭:0;开启:1",
                        "name": "action_type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "totp",
                        "name": "totp",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
