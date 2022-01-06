package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Swagger struct {
	Info                  SwaggerInfo            `json:"info"`
	OpenApi               string                 `json:"openapi"`
	Servers               []SwaggerServer        `json:"servers"`
	Tags                  []SwaggerTag           `json:"tags"`
	Paths                 map[string]SwaggerPath `json:"paths"`
	Components            SwaggerComponent       `json:"components"`
	ShouldGenerateSwagger bool
}

func (s *Swagger) SetInfo(info SwaggerInfo) *Swagger {
	s.Info = info
	return s
}

func (s *Swagger) AddServer(servers SwaggerServer) *Swagger {
	s.Servers = append(s.Servers, servers)
	return s
}

func (s *Swagger) AddPath(path string, pathObj SwaggerPath) *Swagger {
	s.Paths[path] = pathObj
	return s
}
func (s *Swagger) AddTag(tag SwaggerTag) *Swagger {
	s.Tags = append(s.Tags, tag)
	return s
}
func (s *Swagger) AddComponent(component SwaggerComponent) *Swagger {
	s.Components = component
	return s
}
func (s *Swagger) SetShouldGenerateSwagger(shouldGenerateSwagger bool) *Swagger {
	s.ShouldGenerateSwagger = shouldGenerateSwagger
	return s
}

func (s *Swagger) ToJson() []byte {
	var w io.Writer = &bytes.Buffer{}
	json.NewEncoder(w).Encode(s)
	return w.(*bytes.Buffer).Bytes()
}

func (s *Swagger) GenerateSwagger() {
	if s.ShouldGenerateSwagger {
		err := os.WriteFile("swagger.json", []byte(s.ToJson()), 0644)
		if err != nil {
			fmt.Printf("Error writing swagger file: %s\n", err)
		}
		dir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %s\n", err)
		}
		fmt.Printf("Swagger file generated at %s/swagger.json\n", dir)
	}
}

var swaggerInstance *Swagger

func NewSwagger() *Swagger {
	if swaggerInstance == nil {
		swaggerInstance = &Swagger{
			OpenApi: "3.0.2",
			Tags:    []SwaggerTag{},
			Components: SwaggerComponent{
				Schemas: map[string]SwaggerSchema{},
			},
			Paths:   map[string]SwaggerPath{},
			Info:    SwaggerInfo{},
			Servers: []SwaggerServer{},
		}
	}
	return swaggerInstance
}

type SwaggerInfo struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Version     string         `json:"version"`
	Contact     SwaggerContact `json:"contact"`
}

type SwaggerContact struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Email string `json:"email"`
}

type SwaggerServer struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

type SwaggerTag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SwaggerPath struct {
	Get    SwaggerPathItem `json:"get"`
	Post   SwaggerPathItem `json:"post"`
	Put    SwaggerPathItem `json:"put"`
	Delete SwaggerPathItem `json:"delete"`
	Patch  SwaggerPathItem `json:"patch"`
}

type SwaggerPathItem struct {
	Summary     string                     `json:"summary"`
	Description string                     `json:"description"`
	OperationId string                     `json:"operationId"`
	Parameters  []SwaggerParameter         `json:"parameters"`
	Responses   map[string]SwaggerResponse `json:"responses"`
	RequestBody SwaggerRequestBody         `json:"requestBody"`
	Tags        []string                   `json:"tags"`
}
type SwaggerResponse struct {
	Description string                            `json:"description"`
	Content     map[string]SwaggerResponseContent `json:"content"`
}

type SwaggerRequestBody struct {
	Description string                            `json:"description"`
	Required    bool                              `json:"required"`
	Content     map[string]SwaggerResponseContent `json:"content"`
}

type SwaggerParameter struct {
	Name        string        `json:"name"`
	In          string        `json:"in"`
	Description string        `json:"description"`
	Required    bool          `json:"required"`
	Schema      SwaggerSchema `json:"schema"`
}

type SwaggerSchema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Items      map[string]interface{} `json:"items"`
	Required   []string               `json:"required"`
	Ref        string                 `json:"$ref"`
}

type SwaggerResponseContent struct {
}

type SwaggerComponent struct {
	Schemas map[string]SwaggerSchema `json:"schemas"`
}

func (c *SwaggerComponent) AddSchema(schema SwaggerSchema) {
	c.Schemas[schema.Ref] = schema
}

type SwaggerObject struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required   []string               `json:"required"`
}

func NewSwaggerObject() *SwaggerObject {
	return &SwaggerObject{Type: "object"}
}

func (so *SwaggerObject) AddProperty(name string, objType interface{}, isRequired bool) *SwaggerObject {
	so.Properties[name] = objType
	if isRequired {
		so.Required = append(so.Required, name)
	}
	return so
}

type SwaggerArray struct {
	Type  string        `json:"type"`
	Items []interface{} `json:"items"`
}

func NewSwaggerArray() *SwaggerArray {
	return &SwaggerArray{Type: "array"}
}

func (sa *SwaggerArray) AddItem(item interface{}) *SwaggerArray {
	sa.Items = append(sa.Items, item)
	return sa
}

func GetSwaggerType(sType string) interface{} {
	switch sType {
	case "string":
		return SwaggerObject{Type: "string"}
	case "int", "int32", "int64", "uint", "uint32", "uint64", "float32", "float64":
		return SwaggerObject{Type: "number"}
	case "bool":
		return SwaggerObject{Type: "boolean"}
	case "array":
		return SwaggerArray{Type: "array"}
	case "struct":
		return SwaggerObject{Type: "object"}
	default:
		return nil
	}
}

func IsSwaggerArray(obj interface{}) bool {
	switch obj.(type) {
	case SwaggerArray:
		return true
	default:
		return false
	}
}
