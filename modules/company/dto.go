package company

import "github.com/amrebada/go-modules/modules/user"

// Create DTO
type CreateCompanyDto struct {
	Name        string `json:"name" binding:"required"`
	Image       string `json:"image" binding:"max=255"`
	Description string `json:"description" binding:"max=500"`
	UserId      string `json:"userId" binding:"required"`
}

type CreateCompanyResponseDto = CompanyEntity

// GET Company DTO
type GetCompanyDto struct {
	ID string `json:"id" binding:"required"`
}

type GetCompanyResponseDto struct {
	Company CompanyEntity     `json:"company"`
	Users   []user.UserEntity `json:"users"`
}

// Search DTO
type SearchCompanyByNameDto struct {
	Name  string `json:"name" binding:"max=255"`
	Page  int    `json:"page" binding:"default=1"`
	Limit int    `json:"limit" binding:"default=10"`
}

type SearchCompanyByNameResponseDto struct {
	Companies []CompanyEntity `json:"companies"`
	Total     int             `json:"total"`
	Page      int             `json:"page"`
	Limit     int             `json:"limit"`
}

// Update DTO
type UpdateCompanyDto struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name"`
	Image       string `json:"image" binding:"max=255"`
	Description string `json:"description" binding:"max=500"`
}

type UpdateCompanyResponseDto = CompanyEntity

// Remove user from company
type RemoveUserFromCompanyDto struct {
	CompanyId string `json:"companyId" binding:"required"`
	UserId    string `json:"userId" binding:"required"`
}

type RemoveUserFromCompanyResponseDto struct {
	IsRemoved bool `json:"isRemoved"`
}

// Change Role of user in company
type ChangeUserRoleInCompanyDto struct {
	CompanyId string `json:"companyId" binding:"required"`
	UserId    string `json:"userId" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

type ChangeUserRoleInCompanyResponseDto struct {
	IsChanged bool `json:"isChanged"`
}
