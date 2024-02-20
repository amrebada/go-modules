package dtos

import "github.com/amrebada/go-modules/modules/user/entities"

type RegisterDto struct {
	AccessToken string `json:"accessToken" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

type RegisterResponseDto struct {
	Token string        `json:"token"`
	User  entities.User `json:"user"`
}

type FindUserDto struct {
	Name string `json:"name" form:"name"`
}

type FindUserResponseDto struct {
	Users []entities.User `json:"users"`
	Total int64           `json:"total"`
}

type LoginUserDto struct {
	Name string `json:"name" form:"name"`
}

type LoginUserResponseDto struct {
	Users []entities.User `json:"users"`
	Total int64           `json:"total"`
}
