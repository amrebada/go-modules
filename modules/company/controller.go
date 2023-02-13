package company

import (
	"github.com/amrebada/go-modules/core"
)

type CompanyController = core.Controller

func NewCompanyController() *CompanyController {
	return &core.Controller{
		Name:     "company",
		Version:  "v1",
		Path:     "/company",
		Handlers: []*core.Handler{},
	}

}
