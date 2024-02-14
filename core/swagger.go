package core

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

const SWAGGER_TAG = "swagger"

type Swagger struct {
	Info       SwaggerInfo            `json:"info"`
	OpenApi    string                 `json:"openapi"`
	Servers    []SwaggerServer        `json:"servers,omitempty"`
	Tags       []SwaggerTag           `json:"tags,omitempty"`
	Paths      map[string]SwaggerPath `json:"paths,omitempty"`
	Components SwaggerComponent       `json:"components,omitempty"`
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

func (s *Swagger) ToJson() []byte {
	jsonObj, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Errorf("Error marshalling swagger: %s\n", err)
	}
	return jsonObj
}

/**
 * Write swagger file to the root directory
 */
func (s *Swagger) Write() {
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
	RequestBody *SwaggerRequestBody                `json:"requestBody,omitempty"`
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

func (spi *SwaggerPathItem) AddQueryParameters(obj interface{}, swagger *Swagger) *SwaggerPathItem {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}
	swaggerQueryParameters := convertStructToQueryParameters(objType, swagger)
	spi.Parameters = append(spi.Parameters, swaggerQueryParameters...)
	return spi
}

type SwaggerResponse struct {
	Description string                            `json:"description,omitempty"`
	Content     map[string]SwaggerResponseContent `json:"content,omitempty"`
}

type SwaggerRequestBody struct {
	Description string                            `json:"description,omitempty"`
	Required    bool                              `json:"required,omitempty"`
	Content     map[string]SwaggerResponseContent `json:"content,omitempty"`
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
	Type                 SwaggerDataType          `json:"type,omitempty"`
	Format               string                   `json:"format,omitempty"`
	Properties           map[string]SwaggerSchema `json:"properties,omitempty"`
	Items                *SwaggerSchema           `json:"items,omitempty"`
	Required             []string                 `json:"required,omitempty"`
	Ref                  string                   `json:"$ref,omitempty"`
	AdditionalProperties *SwaggerSchema           `json:"additionalProperties,omitempty"`
	Example              string                   `json:"example,omitempty"`
	Description          string                   `json:"description,omitempty"`
}

type SwaggerResponseContent struct {
	Schema SwaggerSchema `json:"schema"`
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
	SWAGGER_DATA_TYPE_INTEGER SwaggerDataType = "integer"
	SWAGGER_DATA_TYPE_BOOLEAN SwaggerDataType = "boolean"
)

func ConvertDtoToSchema(obj interface{}, _swaggerInstance *Swagger, description string) SwaggerSchema {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}
	name := objType.Name()

	if objType.Kind() == reflect.Slice && name == "" {
		name = fmt.Sprintf("ListOf%s", objType.Elem().Name())
	}

	schema := SwaggerSchema{
		Ref: fmt.Sprintf("#/components/schemas/%s", name),
	}

	if _, ok := _swaggerInstance.Components.Schemas[name]; !ok {
		_swaggerInstance.Components.Schemas[name] = convertStructToDefinition(objType, _swaggerInstance, description)
	}

	return schema
}

type SwaggerGoTag struct {
	Example     string
	Format      string
	DataType    SwaggerDataType
	Skip        bool
	Required    bool
	Description string
}

func fieldTypeToSwaggerType(t reflect.Type, _swaggerInstance *Swagger, goTag SwaggerGoTag) SwaggerSchema {
	typeInfo := SwaggerSchema{
		Description: goTag.Description,
	}
	if goTag.Skip {
		return typeInfo
	}
	if goTag.DataType == "" {
		switch t.Kind() {
		case reflect.String:
			typeInfo.Type = SWAGGER_DATA_TYPE_STRING
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			typeInfo.Type = SWAGGER_DATA_TYPE_INTEGER
			typeInfo.Format = "int32"
		case reflect.Int64:
			typeInfo.Type = SWAGGER_DATA_TYPE_INTEGER
			typeInfo.Format = "int64"
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
			typeInfo.Type = SWAGGER_DATA_TYPE_INTEGER
			typeInfo.Format = "int32"
		case reflect.Uint64:
			typeInfo.Type = SWAGGER_DATA_TYPE_INTEGER
			typeInfo.Format = "int64"
		case reflect.Bool:
			typeInfo.Type = SWAGGER_DATA_TYPE_BOOLEAN
		case reflect.Float32:
			typeInfo.Type = SWAGGER_DATA_TYPE_NUMBER
			typeInfo.Format = "float"
		case reflect.Float64:
			typeInfo.Type = SWAGGER_DATA_TYPE_NUMBER
			typeInfo.Format = "double"
		case reflect.Slice, reflect.Array:
			typeInfo.Type = SWAGGER_DATA_TYPE_ARRAY
			elemType := fieldTypeToSwaggerType(t.Elem(), _swaggerInstance, goTag)
			typeInfo.Items = &elemType
		case reflect.Map:
			typeInfo.Type = SWAGGER_DATA_TYPE_OBJECT
			additionalProperties := fieldTypeToSwaggerType(t.Elem(), _swaggerInstance, goTag)
			typeInfo.AdditionalProperties = &additionalProperties
		case reflect.Ptr:
			return fieldTypeToSwaggerType(t.Elem(), _swaggerInstance, goTag)
		case reflect.Struct:
			structName := t.Name()

			if _, exists := _swaggerInstance.Components.Schemas[structName]; !exists {
				definition := convertStructToDefinition(t, _swaggerInstance, goTag.Description)
				_swaggerInstance.Components.Schemas[structName] = definition
			}
			typeInfo.Ref = "#/components/schemas/" + structName
		default:
			typeInfo.Type = SWAGGER_DATA_TYPE_STRING
		}
	} else {
		typeInfo.Type = goTag.DataType
	}
	if goTag.Format != "" {
		typeInfo.Format = goTag.Format
	}

	if goTag.Example != "" {
		typeInfo.Example = goTag.Example
	}
	return typeInfo
}

func convertStructToQueryParameters(obj reflect.Type, swaggerInstance *Swagger) []SwaggerParameter {
	parameters := []SwaggerParameter{}
	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)
		jsonTag := field.Tag.Get("json")
		swaggerTag := field.Tag.Get(SWAGGER_TAG)
		if jsonTag == "-" {
			continue
		}

		goTag := buildSwaggerTag(swaggerTag)
		if goTag.Skip {
			continue
		}

		if goTag.Required {
			parameters = append(parameters, SwaggerParameter{
				Name:        jsonTag,
				In:          SWAGGER_PARAMETER_TYPE_QUERY,
				Description: goTag.Description,
				Required:    true,
				Schema:      fieldTypeToSwaggerType(field.Type, swaggerInstance, goTag),
			})
		} else {
			parameters = append(parameters, SwaggerParameter{
				Name:        jsonTag,
				In:          SWAGGER_PARAMETER_TYPE_QUERY,
				Description: goTag.Description,
				Schema:      fieldTypeToSwaggerType(field.Type, swaggerInstance, goTag),
			})
		}
	}
	return parameters
}

func convertStructToDefinition(t reflect.Type, _swaggerInstance *Swagger, description string) SwaggerSchema {

	schema := SwaggerSchema{
		Type:        SWAGGER_DATA_TYPE_OBJECT,
		Properties:  map[string]SwaggerSchema{},
		Required:    []string{},
		Description: description,
	}

	isSlice := false

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Slice {
		t = t.Elem()
		isSlice = true
		schema = SwaggerSchema{
			Type: SWAGGER_DATA_TYPE_ARRAY,
			Items: &SwaggerSchema{
				Type:       SWAGGER_DATA_TYPE_OBJECT,
				Properties: map[string]SwaggerSchema{},
				Required:   []string{},
			},
			Description: description,
		}
	}

	fields := getFieldsOfType(t)

	for i := 0; i < len(fields); i++ {
		field := fields[i]

		jsonTag := field.Tag.Get("json")
		swaggerTag := field.Tag.Get(SWAGGER_TAG)
		if jsonTag == "-" {
			continue
		}

		goTag := buildSwaggerTag(swaggerTag)
		if goTag.Skip {
			continue
		}

		if goTag.Required {
			schema.Required = append(schema.Required, jsonTag)
		}

		fieldInfo := fieldTypeToSwaggerType(field.Type, _swaggerInstance, goTag)

		if isSlice {
			schema.Items.Properties[jsonTag] = fieldInfo
		} else {
			schema.Properties[jsonTag] = fieldInfo
		}
	}

	return schema
}

func getFieldsOfType(t reflect.Type) []reflect.StructField {
	fields := []reflect.StructField{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			fields = append(fields, getFieldsOfType(field.Type)...)
		} else {
			fields = append(fields, field)
		}
	}
	return fields
}

func buildSwaggerTag(tag string) SwaggerGoTag {
	swaggerTag := SwaggerGoTag{}
	tagParts := strings.Split(tag, ",")

	for _, tagPart := range tagParts {
		tagPart = strings.TrimSpace(tagPart)
		if tagPart == "" {
			continue
		}
		if tagPart == "-" {
			swaggerTag.Skip = true
			continue
		}
		tagSides := strings.Split(tagPart, "=")
		tagName := strings.TrimSpace(tagSides[0])
		tagValue := ""
		if len(tagSides) == 2 {
			tagValue = strings.TrimSpace(strings.Split(tagPart, "=")[1])
		}
		switch tagName {
		case "required":
			swaggerTag.Required = true
		case "example":
			swaggerTag.Example = tagValue
		case "format":
			swaggerTag.Format = tagValue
		case "description":
			swaggerTag.Description = tagValue
		case "type":
			swaggerTag.DataType = SwaggerDataType(tagValue)
		}
	}

	return swaggerTag
}
