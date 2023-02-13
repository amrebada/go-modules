package company

import "github.com/amrebada/go-modules/core"

type CompanyModule = core.Module

func New() *CompanyModule {
	return &core.Module{
		Name: "company",
		Controllers: []*core.Controller{
			NewCompanyController(),
		},
		Entities: []interface{}{
			CompanyEntity{},
			CompanyUserEntity{},
		},
	}
}
