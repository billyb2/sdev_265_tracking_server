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
        "/get_tracking": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Gets the status of tracking numbers",
                "operationId": "get-tracking-numbers",
                "parameters": [
                    {
                        "description": "Tracking numbers",
                        "name": "getTrackingInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.getTracking"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.trackingInfo"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api.trackingInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.trackingInfo"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Verifies and logs in the user, returning a token",
                "operationId": "login-user",
                "parameters": [
                    {
                        "description": "Login Info",
                        "name": "loginInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.loginInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.loginResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.loginResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.loginResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Registers a new user",
                "operationId": "register-user",
                "parameters": [
                    {
                        "description": "Registration Info",
                        "name": "registrationInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.registrationInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.registerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.registerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.registerResponse"
                        }
                    }
                }
            }
        },
        "/start_tracking": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Starts tracking the package tracking numbers given by the user",
                "operationId": "start-tracking-groups",
                "parameters": [
                    {
                        "description": "Tracking Info",
                        "name": "startTrackingInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.startTracking"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.startTrackingResp"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api.startTrackingResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.startTrackingResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.getTracking": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "tracking_numbers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.loginInfo": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.loginResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "api.registerResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "api.registrationInfo": {
            "type": "object",
            "properties": {
                "company": {
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
        "api.startTracking": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "tracking_number_groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.trackingNumberGroup"
                    }
                }
            }
        },
        "api.startTrackingResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "api.trackingInfo": {
            "type": "object",
            "properties": {
                "tracking_info": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "api.trackingNumberGroup": {
            "type": "object",
            "properties": {
                "group_name": {
                    "type": "string"
                },
                "tracking_numbers": {
                    "type": "array",
                    "items": {
                        "type": "string"
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
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Tracking Server API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
