package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

type Swagger struct {
	Info                  SwaggerInfo            `json:"info"`
	OpenApi               string                 `json:"openapi"`
	Servers               []SwaggerServer        `json:"servers,omitempty"`
	Tags                  []SwaggerTag           `json:"tags,omitempty"`
	Paths                 map[string]SwaggerPath `json:"paths,omitempty"`
	Components            SwaggerComponent       `json:"components,omitempty"`
	ShouldGenerateSwagger bool                   `json:"-"`
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
		os.Exit(0)
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
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	Version     string         `json:"version,omitempty"`
	Contact     SwaggerContact `json:"contact,omitempty"`
}

type SwaggerContact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type SwaggerServer struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

type SwaggerTag struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type HttpMethods string

const (
	HTTP_GET_METHOD    HttpMethods = "get"
	HTTP_POST_METHOD   HttpMethods = "post"
	HTTP_DELETE_METHOD HttpMethods = "delete"
	HTTP_PUT_METHOD    HttpMethods = "put"
	HTTP_PATCH_METHOD  HttpMethods = "patch"
)

type SwaggerPath = map[HttpMethods]SwaggerPathItem

type SwaggerPathItem struct {
	Summary     string                             `json:"summary"`
	Description string                             `json:"description,omitempty"`
	OperationId string                             `json:"operationId,omitempty"`
	Parameters  []SwaggerParameter                 `json:"parameters,omitempty"`
	Responses   map[HttpStatusCode]SwaggerResponse `json:"responses"`
	RequestBody SwaggerRequestBody                 `json:"requestBody,omitempty"`
	Tags        []string                           `json:"tags"`
}

func (spi *SwaggerPathItem) AddParameter(parameter SwaggerParameter) *SwaggerPathItem {
	spi.Parameters = append(spi.Parameters, parameter)
	return spi
}

func (spi *SwaggerPathItem) AddPathParameter(parameter string) *SwaggerPathItem {
	swaggerPathParameter := SwaggerParameter{
		Name:        parameter,
		In:          SWAGGER_PARAMETER_TYPE_PATH,
		Description: fmt.Sprintf("path parameter %s ", parameter),
		Required:    true,
		Schema: SwaggerSchema{
			Type: SWAGGER_DATA_TYPE_STRING,
		},
	}
	spi.Parameters = append(spi.Parameters, swaggerPathParameter)
	return spi
}

//TODO: continue
// func (spi *SwaggerPathItem) AddQueryParameter(parameter interface{}) *SwaggerPathItem {

// 	swaggerQueryParameter := SwaggerParameter{
// 		Name:        parameter,
// 		In:          SWAGGER_PARAMETER_TYPE_PATH,
// 		Description: fmt.Sprintf("path parameter %s ", parameter),
// 		Required:    true,
// 		Schema: SwaggerSchema{
// 			Type: SWAGGER_DATA_TYPE_STRING,
// 		},
// 	}
// 	spi.Parameters = append(spi.Parameters, swaggerQueryParameter)
// 	return spi
// }

type SwaggerResponse struct {
	Description string                            `json:"description,omitempty"`
	Content     map[string]SwaggerResponseContent `json:"content,omitempty"`
}

type SwaggerRequestBody struct {
	Description string                            `json:"description,omitempty"`
	Required    bool                              `json:"required,omitempty"`
	Content     map[string]SwaggerResponseContent `json:"content"`
}
type SwaggerParameterType string

const (
	SWAGGER_PARAMETER_TYPE_PATH   SwaggerParameterType = "path"
	SWAGGER_PARAMETER_TYPE_QUERY  SwaggerParameterType = "query"
	SWAGGER_PARAMETER_TYPE_HEADER SwaggerParameterType = "header"
)

type SwaggerParameter struct {
	Name        string               `json:"name"`
	In          SwaggerParameterType `json:"in"`
	Description string               `json:"description,omitempty"`
	Required    bool                 `json:"required"`
	Schema      SwaggerSchema        `json:"schema"`
}

type SwaggerSchema struct {
	Type       SwaggerDataType        `json:"type"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Items      map[string]interface{} `json:"items,omitempty"`
	Required   []string               `json:"required,omitempty"`
	Ref        string                 `json:"$ref,omitempty"`
}

type SwaggerResponseContent struct {
}

type SwaggerComponent struct {
	Schemas map[string]SwaggerSchema `json:"schemas,omitempty"`
}

func (c *SwaggerComponent) AddSchema(schema SwaggerSchema) {
	c.Schemas[schema.Ref] = schema
}

type SwaggerDataType string

const (
	SWAGGER_DATA_TYPE_OBJECT  SwaggerDataType = "object"
	SWAGGER_DATA_TYPE_ARRAY   SwaggerDataType = "array"
	SWAGGER_DATA_TYPE_STRING  SwaggerDataType = "string"
	SWAGGER_DATA_TYPE_NUMBER  SwaggerDataType = "number"
	SWAGGER_DATA_TYPE_BOOLEAN SwaggerDataType = "boolean"
)

type SwaggerObject struct {
	Type       SwaggerDataType        `json:"type"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Required   []string               `json:"required,omitempty"`
}

func NewSwaggerObject() *SwaggerObject {
	return &SwaggerObject{Type: SWAGGER_DATA_TYPE_OBJECT}
}

func (so *SwaggerObject) AddProperty(name string, objType interface{}, isRequired bool) *SwaggerObject {
	so.Properties[name] = objType
	if isRequired {
		so.Required = append(so.Required, name)
	}
	return so
}

type SwaggerArray struct {
	Type  SwaggerDataType `json:"type"`
	Items []interface{}   `json:"items"`
}

func NewSwaggerArray() *SwaggerArray {
	return &SwaggerArray{Type: SWAGGER_DATA_TYPE_ARRAY}
}

func (sa *SwaggerArray) AddItem(item interface{}) *SwaggerArray {
	sa.Items = append(sa.Items, item)
	return sa
}

func GetSwaggerObjectType(sType reflect.Kind) interface{} {
	switch sType {
	case reflect.String:
		return SwaggerObject{Type: SWAGGER_DATA_TYPE_STRING}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return SwaggerObject{Type: SWAGGER_DATA_TYPE_NUMBER}
	case reflect.Bool:
		return SwaggerObject{Type: SWAGGER_DATA_TYPE_BOOLEAN}
	case reflect.Array, reflect.Slice:
		return SwaggerArray{Type: SWAGGER_DATA_TYPE_ARRAY}
	case reflect.Struct, reflect.Map, reflect.Interface:
		return SwaggerObject{Type: SWAGGER_DATA_TYPE_OBJECT}
	default:
		return nil
	}
}
func GetSwaggerType(sType reflect.Kind) SwaggerDataType {
	switch sType {
	case reflect.String:
		return SWAGGER_DATA_TYPE_STRING
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return SWAGGER_DATA_TYPE_NUMBER
	case reflect.Bool:
		return SWAGGER_DATA_TYPE_BOOLEAN
	case reflect.Array, reflect.Slice:
		return SWAGGER_DATA_TYPE_ARRAY
	case reflect.Struct, reflect.Map, reflect.Interface:
		return SWAGGER_DATA_TYPE_OBJECT
	default:
		return SWAGGER_DATA_TYPE_OBJECT
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

type Schemas = map[string]SwaggerSchema

func ConvertToSchema(obj interface{}, schemas Schemas) Schemas {
	objType := reflect.TypeOf(obj)
	schema := SwaggerSchema{
		Type: GetSwaggerType(objType.Kind()),
	}

	// if objType.Kind() == reflect.Struct {
	// 	for i := 0; i < objType.NumField(); i++ {
	// 		fmt.Printf("\tField %s: %s\n", objType.Field(i).Name, objType.Field(i).Type.Kind().String())
	// 	}
	// }

	// // typeInObject := GetSwaggerObjectType(objType.Kind())
	// typeInBytes, err := json.MarshalIndent(schema, "", "  ")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	schemas[objType.Name()] = schema

	return schemas
}
