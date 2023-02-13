package company

import (
	"errors"

	"github.com/amrebada/go-modules/core"
)

func CreateCompany(companyDto *CreateCompanyDto) (*CreateCompanyResponseDto, error) {
	company := &CompanyEntity{
		Name:        companyDto.Name,
		Description: companyDto.Description,
		Image:       companyDto.Image,
	}
	err := company.Create(companyDto.UserId)
	if err != nil {
		return nil, err
	}
	return company, nil

}
func UpdateCompany(companyDto *UpdateCompanyDto) (*UpdateCompanyResponseDto, error) {
	company := &CompanyEntity{
		EntityWithID: core.EntityWithID{ID: companyDto.ID},
	}
	err := company.GetByID()
	if err != nil {
		return nil, err
	}
	if company.Name == "" {
		return nil, errors.New("Cannot find company")
	}
	if companyDto.Name != "" {
		company.Name = companyDto.Name
	}
	if companyDto.Description != "" {
		company.Description = companyDto.Description
	}
	if companyDto.Image != "" {
		company.Image = companyDto.Image
	}
	err = company.Update()

	if err != nil {
		return nil, err
	}
	return company, nil

}

func GetCompany(companyDto *GetCompanyDto) (*GetCompanyResponseDto, error) {
	company := &CompanyEntity{
		EntityWithID: core.EntityWithID{ID: companyDto.ID},
	}
	err := company.GetByID()
	if err != nil {
		return nil, err
	}

	if company.Name == "" {
		return nil, errors.New("Company not found")
	}
	companyUser := CompanyUserEntity{}
	users, _, _ := companyUser.ListUsersByCompanyId(companyDto.ID, []string{ROLE_COMPANY_ADMIN, ROLE_COMPANY_USER}, 1, 100)
	return &GetCompanyResponseDto{
		Company: *company,
		Users:   users,
	}, nil
}

func SearchCompaniesByName(companyDto *SearchCompanyByNameDto) (*SearchCompanyByNameResponseDto, error) {
	company := &CompanyEntity{}
	companies, total, err := company.SearchByName(companyDto.Name, companyDto.Page, companyDto.Limit)
	if err != nil {
		return nil, err
	}
	return &SearchCompanyByNameResponseDto{
		Companies: companies,
		Total:     int(total),
		Page:      companyDto.Page,
		Limit:     companyDto.Limit,
	}, nil
}

func RemoveUserFromCompany(companyDto *RemoveUserFromCompanyDto) (*RemoveUserFromCompanyResponseDto, error) {
	companyUser := &CompanyUserEntity{
		UserID:    companyDto.UserId,
		CompanyID: companyDto.CompanyId,
	}
	err := companyUser.GetByCompanyAndUser()
	if err != nil {
		return nil, err
	}
	if companyUser.ID == "" {
		return nil, errors.New("this user is not in this company")
	}
	err = companyUser.Delete()
	if err != nil {
		return &RemoveUserFromCompanyResponseDto{
			IsRemoved: false,
		}, err
	}
	return &RemoveUserFromCompanyResponseDto{
		IsRemoved: true,
	}, nil
}

func ChangeRoleOfUser(companyDto *ChangeUserRoleInCompanyDto) (*ChangeUserRoleInCompanyResponseDto, error) {
	companyUser := &CompanyUserEntity{
		UserID:    companyDto.UserId,
		CompanyID: companyDto.CompanyId,
	}
	err := companyUser.GetByCompanyAndUser()
	if err != nil {
		return nil, err
	}
	if companyUser.ID == "" {
		return nil, errors.New("this user is not in this company")
	}
	if companyUser.Role == companyDto.Role {
		return &ChangeUserRoleInCompanyResponseDto{
			IsChanged: true,
		}, nil
	}
	companyUser.Role = companyDto.Role
	err = companyUser.Update()
	if err != nil {
		return nil, err
	}
	return &ChangeUserRoleInCompanyResponseDto{
		IsChanged: true,
	}, nil

}
