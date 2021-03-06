// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/exercises": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exercise"
                ],
                "summary": "Create a new exercise tests",
                "parameters": [
                    {
                        "description": "Add an exercise",
                        "name": "exercise",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_igor-ferreira-almeida_workout-api_src_api_app_presentation_exercisedto_request.ExerciseRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github.com_igor-ferreira-almeida_workout-api_src_api_app_presentation_exercisedto_response.ExerciseResponse"
                        }
                    }
                }
            }
        },
        "/exercises/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exercise"
                ],
                "summary": "Find exercise by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exercise ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github.com_igor-ferreira-almeida_workout-api_src_api_app_presentation_exercisedto_response.ExerciseResponse"
                        }
                    }
                }
            }
        },
        "/training": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "training"
                ],
                "summary": "Create a new training",
                "parameters": [
                    {
                        "description": "Add an training",
                        "name": "training",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.TrainingRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.TrainingResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github.com_igor-ferreira-almeida_workout-api_src_api_app_presentation_exercisedto_request.ExerciseRequest": {
            "type": "object",
            "properties": {
                "muscle_group": {
                    "type": "string",
                    "example": "Chest"
                },
                "name": {
                    "type": "string",
                    "example": "Bench Press"
                },
                "reps": {
                    "type": "integer",
                    "example": 12
                },
                "rest": {
                    "type": "integer",
                    "example": 60
                },
                "weight": {
                    "type": "integer",
                    "example": 20
                }
            }
        },
        "github.com_igor-ferreira-almeida_workout-api_src_api_app_presentation_exercisedto_response.ExerciseResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "muscular_group": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "reps": {
                    "type": "integer"
                },
                "rest": {
                    "type": "integer"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "github.com_igor-ferreira-almeida_workout-api_src_api_app_presentation_trainingdto_request.ExerciseRequest": {
            "type": "object",
            "properties": {
                "executions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.ExecutionRequest"
                    }
                },
                "muscle_group": {
                    "type": "string",
                    "example": "Chest"
                },
                "name": {
                    "type": "string",
                    "example": "Bench Press"
                }
            }
        },
        "github.com_igor-ferreira-almeida_workout-api_src_api_app_presentation_trainingdto_response.ExerciseResponse": {
            "type": "object",
            "properties": {
                "executions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.ExecutionResponse"
                    }
                },
                "muscle_group": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "volume": {
                    "type": "integer"
                }
            }
        },
        "request.ExecutionRequest": {
            "type": "object",
            "properties": {
                "sets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.SetRequest"
                    }
                },
                "technique": {
                    "type": "string",
                    "example": "Normal,Drop Set,Bi Set,Rest Pause"
                }
            }
        },
        "request.SetRequest": {
            "type": "object",
            "properties": {
                "reps": {
                    "type": "integer",
                    "example": 12
                },
                "rest": {
                    "type": "integer",
                    "example": 30
                },
                "weight": {
                    "type": "integer",
                    "example": 20
                }
            }
        },
        "request.TrainingRequest": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string",
                    "example": "2021-12-07T00:00:00-03:00"
                },
                "exercises": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github.com_igor-ferreira-almeida_workout-api_src_api_app_presentation_trainingdto_request.ExerciseRequest"
                    }
                },
                "muscle_groups": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Chest",
                        "Back",
                        "Deltoids",
                        "Biceps",
                        "Triceps",
                        "Forearms",
                        "Quadriceps",
                        "Hamstrings",
                        "Calves"
                    ]
                }
            }
        },
        "response.ExecutionResponse": {
            "type": "object",
            "properties": {
                "sets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.SetResponse"
                    }
                },
                "technique": {
                    "type": "string"
                },
                "volume": {
                    "type": "integer"
                }
            }
        },
        "response.SetResponse": {
            "type": "object",
            "properties": {
                "reps": {
                    "type": "integer"
                },
                "rest": {
                    "type": "integer"
                },
                "volume": {
                    "type": "integer"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "response.TrainingResponse": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "exercises": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github.com_igor-ferreira-almeida_workout-api_src_api_app_presentation_trainingdto_response.ExerciseResponse"
                    }
                },
                "muscle_groups": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "volume_per_muscle_group": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
