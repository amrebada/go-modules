package user

type RegisterDto struct {
	AccessToken string `json:"accessToken" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

type RegisterResponseDto struct {
	Token string     `json:"token"`
	User  UserEntity `json:"user"`
}

type FindUserDto struct {
	Name string `json:"name" form:"name"`
}

type FindUserResponseDto struct {
	Users []UserEntity `json:"users"`
	Total int64        `json:"total"`
}

type LoginUserDto struct {
	Name string `json:"name" form:"name"`
}

type LoginUserResponseDto struct {
	Users []UserEntity `json:"users"`
	Total int64        `json:"total"`
}
