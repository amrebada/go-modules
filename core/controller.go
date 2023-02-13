package core

import (
	"fmt"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
)

/**
- Base Handler
- with Builder pattern
*/

type Handler struct {
	Method      HttpMethods
	Path        string
	Middleware  []FiberHandler
	HandlerFunc FiberHandler
	RequestDto  interface{}
	ResponseDto interface{}
	Description string
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetMethod(method HttpMethods) *Handler {
	h.Method = method
	return h
}

func (h *Handler) SetPath(path string) *Handler {
	h.Path = path
	return h
}

func (h *Handler) AddMiddleware(m ...FiberHandler) *Handler {
	h.Middleware = append(h.Middleware, m...)
	return h
}

func (h *Handler) SetHandlerFunc(handlerFunc FiberHandler) *Handler {
	h.HandlerFunc = handlerFunc
	return h
}

func (h *Handler) SetDescription(description string) *Handler {
	h.Description = description
	return h
}

func (h *Handler) SetRequestDto(requestDto interface{}) *Handler {
	h.RequestDto = requestDto
	return h
}

func (h *Handler) SetResponseDto(responseDto interface{}) *Handler {
	h.ResponseDto = responseDto
	return h
}

func (h *Handler) GenerateSwagger(controllerName string, version string, controllerPath string, moduleName string) {
	fmt.Println("    H", generate, " Swagger for handler", GetStatusString(h.Method), ":", h.Path)
	/**
	Summary     string                     `json:"summary"`
	Description string                     `json:"description"`
	OperationId string                     `json:"operationId"`
	Parameters  []SwaggerParameter         `json:"parameters"`
	Responses   map[string]SwaggerResponse `json:"responses"`
	RequestBody SwaggerRequestBody         `json:"requestBody"`
	Tags        []string                   `json:"tags"`
	*/
	operationId := GenerateOperationId(h.Path, controllerName, h.Method, version)
	pathParameters := ExtractPathParameters(h.Path)
	pathItem := SwaggerPathItem{
		Summary:     h.Description,
		Description: h.Description,
		OperationId: operationId,
		Tags:        []string{moduleName},
		RequestBody: SwaggerRequestBody{
			Content: map[string]SwaggerResponseContent{},
		},
		Responses: map[HttpStatusCode]SwaggerResponse{
			HttpStatusOK: {},
		},
	}
	for _, pathParameter := range pathParameters {
		pathItem.AddPathParameter(pathParameter)
	}
	if swaggerInstance.Paths[GenerateFullPath(h.Path, version, controllerPath)] == nil {
		swaggerInstance.Paths[GenerateFullPath(h.Path, version, controllerPath)] = SwaggerPath{}
	}

	swaggerInstance.Paths[GenerateFullPath(h.Path, version, controllerPath)][h.Method] = pathItem
}

/**
- Base controller
- with Builder pattern
*/

type Controller struct {
	Name     string
	Version  string
	Path     string
	Handlers []*Handler
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) SetName(name string) *Controller {
	c.Name = name
	return c
}

func (c *Controller) SetVersion(version string) *Controller {
	c.Version = version
	return c
}

func (c *Controller) SetPath(path string) *Controller {
	c.Path = path
	return c
}

func (c *Controller) AddHandler(h *Handler) *Controller {
	c.Handlers = append(c.Handlers, h)
	return c
}

func (c *Controller) RegisterRoutes(e *fiber.App) {
	for _, h := range c.Handlers {
		full_path := path.Join(c.Version, c.Path, h.Path)
		fmt.Println("    H", register, " NewHandler: ", GetStatusString(h.Method), full_path)
		h.Middleware = append(h.Middleware, h.HandlerFunc)
		e.Add(strings.ToUpper(string(h.Method)), full_path, h.Middleware...)
	}
}

func (c *Controller) GenerateSwagger(moduleName string) {
	fmt.Println("  C", generate, " Swagger for controller: ", c.Name)
	for _, h := range c.Handlers {
		h.GenerateSwagger(c.Name, c.Version, c.Path, moduleName)
	}
}
