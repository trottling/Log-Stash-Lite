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
        "/add_log": {
            "post": {
                "description": "Adds a single log entry to storage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logs"
                ],
                "summary": "Add a log entry",
                "parameters": [
                    {
                        "description": "Log entry",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddLogRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.AddLogResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/add_logs": {
            "post": {
                "description": "Adds multiple log entries to storage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logs"
                ],
                "summary": "Add multiple log entries",
                "parameters": [
                    {
                        "description": "Log entries",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddLogsRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.AddLogsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get_logs": {
            "get": {
                "description": "Returns logs with filters and limit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logs"
                ],
                "summary": "Get logs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Log level",
                        "name": "level",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetLogsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get_logs_count": {
            "get": {
                "description": "Returns only count of logs by filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logs"
                ],
                "summary": "Get logs count",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Log level",
                        "name": "level",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetLogsCountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Returns health status of the service and system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.HealthResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/dto.HealthResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AddLogRequest": {
            "type": "object",
            "required": [
                "log",
                "parse_type"
            ],
            "properties": {
                "log": {
                    "type": "object",
                    "additionalProperties": true
                },
                "parse_type": {
                    "type": "string",
                    "enum": [
                        "default",
                        "zap",
                        "logrus",
                        "pino"
                    ]
                }
            }
        },
        "dto.AddLogResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.AddLogsRequest": {
            "type": "object",
            "required": [
                "logs",
                "parse_type"
            ],
            "properties": {
                "logs": {
                    "type": "array",
                    "minItems": 1,
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                },
                "parse_type": {
                    "type": "string",
                    "enum": [
                        "default",
                        "zap",
                        "logrus",
                        "pino"
                    ]
                }
            }
        },
        "dto.AddLogsResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "dto.GetLogsCountResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "dto.GetLogsResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "logs": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                }
            }
        },
        "dto.HealthResponse": {
            "type": "object",
            "properties": {
                "elastic_status": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "system_status": {
                    "type": "object",
                    "properties": {
                        "cpu": {
                            "type": "object",
                            "properties": {
                                "temperature": {
                                    "type": "number"
                                },
                                "usage_percent": {
                                    "type": "number"
                                }
                            }
                        },
                        "ram": {
                            "type": "object",
                            "properties": {
                                "total_mb": {
                                    "type": "integer"
                                },
                                "used_mb": {
                                    "type": "integer"
                                }
                            }
                        },
                        "rom": {
                            "type": "object",
                            "properties": {
                                "total_mb": {
                                    "type": "integer"
                                },
                                "used_mb": {
                                    "type": "integer"
                                }
                            }
                        }
                    }
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
	//LeftDelim:        "{{",
	//RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
