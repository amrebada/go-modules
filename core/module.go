package core

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Entity = interface{}

type Module struct {
	Name        string
	Controllers []*Controller
	Imports     []*Module
	Entities    []Entity
	Extensions  interface{}
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) SetName(name string) *Module {
	m.Name = name
	return m
}

func (m *Module) AddController(c *Controller) *Module {
	m.Controllers = append(m.Controllers, c)
	return m
}

func (m *Module) AddImport(im *Module) *Module {
	m.Imports = append(m.Imports, im)
	return m
}

func (m *Module) AddEntity(e Entity) *Module {
	m.Entities = append(m.Entities, e)
	return m
}

func (m *Module) SetExtensions(ext interface{}) *Module {
	m.Extensions = ext
	return m
}

func (m *Module) Migrate() error {

	fmt.Println(" M", migrate, " NewModule: ", m.Name)
	db := NewDatabase()
	err := db.DB.AutoMigrate(m.Entities...)

	if err != nil {
		return err
	}
	for _, imported := range m.Imports {
		err = imported.Migrate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Module) RegisterRoutes(e *gin.Engine) {
	fmt.Println(" M", register, " NewModule: ", m.Name)
	for _, c := range m.Controllers {
		fmt.Println("   C", register, " NewController: ", c.Path)
		c.RegisterRoutes(e)
	}
	m.GenerateSwagger()
	for _, imported := range m.Imports {
		imported.RegisterRoutes(e)
		imported.GenerateSwagger()
	}
}

func (m *Module) GenerateSwagger() {
	fmt.Println("M", generate, " Swagger for module: ", m.Name)
	for _, c := range m.Controllers {
		c.GenerateSwagger(m.Name)
	}
}
