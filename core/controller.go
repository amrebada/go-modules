package core

import (
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
)

/**
- Base Handler
- with Builder pattern
*/

type Handler struct {
	Method      string
	Path        string
	Middleware  []gin.HandlerFunc
	HandlerFunc gin.HandlerFunc
	RequestDto  interface{}
	ResponseDto interface{}
	Description string
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetMethod(method string) *Handler {
	h.Method = method
	return h
}

func (h *Handler) SetPath(path string) *Handler {
	h.Path = path
	return h
}

func (h *Handler) AddMiddleware(m ...gin.HandlerFunc) *Handler {
	h.Middleware = append(h.Middleware, m...)
	return h
}

func (h *Handler) SetHandlerFunc(handlerFunc gin.HandlerFunc) *Handler {
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

func (c *Controller) RegisterRoutes(e *gin.Engine) {
	for _, h := range c.Handlers {
		full_path := path.Join(c.Version, c.Path, h.Path)
		fmt.Println("    H", register, " NewHandler: ", full_path)
		h.Middleware = append(h.Middleware, h.HandlerFunc)
		e.Handle(h.Method, full_path, h.Middleware...)
	}
}

func (c *Controller) GenerateSwagger(moduleName string) {
	fmt.Println("  C", generate, " Swagger for controller: ", c.Name)
	for _, h := range c.Handlers {
		fmt.Println("    H", generate, " Swagger for handler: ", h.Path)

	}
}
