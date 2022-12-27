// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/post/collect": {
            "post": {
                "description": "收藏推文",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "推文"
                ],
                "summary": "收藏",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "推文编号",
                        "name": "uuid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PostCollectReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/comment-list": {
            "get": {
                "description": "查询推文的评论列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "推文"
                ],
                "summary": "查询评论",
                "parameters": [
                    {
                        "type": "string",
                        "description": "评论列表输入参数",
                        "name": "list",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/list": {
            "get": {
                "description": "查询主题信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "推文"
                ],
                "summary": "查询",
                "parameters": [
                    {
                        "type": "string",
                        "description": "列表输入参数",
                        "name": "list",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/publish": {
            "post": {
                "description": "发布一条推文",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "推文"
                ],
                "summary": "发布",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "推文信息",
                        "name": "publish_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PostPublishReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/show": {
            "get": {
                "description": "查询主题信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "主题"
                ],
                "summary": "查询",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "推文uuid",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/star": {
            "post": {
                "description": "给推文点赞",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "推文"
                ],
                "summary": "点赞",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "推文编号",
                        "name": "uuid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PostStarReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/topic/list": {
            "post": {
                "description": "查询主题信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "主题"
                ],
                "summary": "查询",
                "parameters": [
                    {
                        "description": "列表输入参数",
                        "name": "list",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.TopicListReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/topic/{topic_id}": {
            "get": {
                "description": "查询主题信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "主题"
                ],
                "summary": "查询",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "主题编号",
                        "name": "topic_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "新增用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserInfoCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user-member/show": {
            "get": {
                "description": "填写基本信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "玩家信息"
                ],
                "summary": "查询",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user-member/sign-in": {
            "post": {
                "description": "填写登录信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "玩家信息"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "登录信息",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserMemberSignInReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user-member/sign-up": {
            "post": {
                "description": "填写基本信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "玩家信息"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "description": "玩家信息",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserMemberSignUpReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user-member/suggest": {
            "get": {
                "description": "填写基本信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "玩家信息"
                ],
                "summary": "根据用户名联想的玩家",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/show/{id}": {
            "get": {
                "description": "展示指定用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "查询用户",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UserInfoShowRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.PostCollectReq": {
            "type": "object",
            "required": [
                "operate",
                "uuid"
            ],
            "properties": {
                "operate": {
                    "description": "add 是收藏 sub是取消收藏",
                    "type": "string"
                },
                "uuid": {
                    "description": "推文编号",
                    "type": "string"
                }
            }
        },
        "request.PostCommentCreateReq": {
            "type": "object",
            "required": [
                "content",
                "pid",
                "uuid"
            ],
            "properties": {
                "content": {
                    "description": "评论内容",
                    "type": "string"
                },
                "pid": {
                    "description": "上层用户ID，如果是0，表示评论的是帖子",
                    "type": "integer"
                },
                "uuid": {
                    "description": "帖子唯一标识",
                    "type": "string"
                }
            }
        },
        "request.PostPublishReq": {
            "type": "object",
            "required": [
                "content"
            ],
            "properties": {
                "content": {
                    "description": "内容",
                    "type": "string"
                },
                "if_top": {
                    "description": "是否置顶 1-是 2-否",
                    "type": "string"
                },
                "published_city": {
                    "description": "发布城市",
                    "type": "string"
                },
                "published_ip": {
                    "description": "发布Ip",
                    "type": "string"
                },
                "user_member_id": {
                    "description": "创建人",
                    "type": "integer"
                }
            }
        },
        "request.PostStarReq": {
            "type": "object",
            "required": [
                "operate",
                "uuid"
            ],
            "properties": {
                "operate": {
                    "description": "add 是点赞 sub是取消点赞",
                    "type": "string"
                },
                "uuid": {
                    "description": "推文编号",
                    "type": "string"
                }
            }
        },
        "request.TopicListReq": {
            "type": "object",
            "required": [
                "page",
                "size"
            ],
            "properties": {
                "order": {
                    "type": "string"
                },
                "page": {
                    "description": "当前页",
                    "type": "integer"
                },
                "size": {
                    "description": "每页条目",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "request.UserInfoCreateReq": {
            "type": "object",
            "required": [
                "passport",
                "password"
            ],
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "passport": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.UserMemberSignInReq": {
            "type": "object",
            "required": [
                "passport",
                "password"
            ],
            "properties": {
                "passport": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.UserMemberSignUpReq": {
            "type": "object",
            "required": [
                "email",
                "passport",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "passport": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "response.UserInfoShowRes": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "passport": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
