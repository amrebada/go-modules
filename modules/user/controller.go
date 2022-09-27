package user

import (
	"fmt"

	"github.com/amrebada/go-template/core"
	"github.com/amrebada/go-template/home"
	"github.com/gin-gonic/gin"
)

type AuthController = core.Controller

func NewAuthController() *AuthController {

	return core.NewController().
		SetName("Auth").
		SetVersion("v1").
		SetPath("/auth").
		AddHandler(core.NewHandler().
			SetMethod(core.HTTP_POST_METHOD).
			SetPath("/register").
			SetHandlerFunc(RegisterUser).
			SetDescription("Register user").
			SetRequestDto(&RegisterDto{}).
			SetResponseDto(&RegisterResponseDto{})).
		AddHandler(core.NewHandler().
			SetMethod(core.HTTP_POST_METHOD).
			SetPath("/login").
			SetHandlerFunc(LoginUser).
			SetDescription("Login user").
			SetRequestDto(&LoginUserDto{}).
			SetResponseDto(&LoginUserResponseDto{})).
		AddHandler(core.NewHandler().
			SetMethod(core.HTTP_GET_METHOD).
			SetPath("/login").
			SetHandlerFunc(LoginUser).
			SetDescription("get token of user").
			SetResponseDto(&LoginUserResponseDto{})).
		AddHandler(core.NewHandler().
			SetMethod(core.HTTP_GET_METHOD).
			SetPath("/:id").
			SetHandlerFunc(LoginUser).
			SetDescription("get user").
			SetResponseDto(&UserEntity{}))
}

//Login user
func RegisterUser(ctx *gin.Context) {
	registerDto := &RegisterDto{}
	err := ctx.ShouldBindJSON(registerDto)
	if err != nil {
		ctx.JSON(400, home.ErrorResponse([]error{err}, home.CANNOT_PARSE_BODY))
		return
	}
	user, err := Register(registerDto)
	if err != nil {
		ctx.JSON(401, home.ErrorResponse([]error{err}, home.OAUTH_TOKEN_NOT_CORRECT))
		return
	}
	ctx.JSON(200, user)
}

func LoginUser(ctx *gin.Context) {
	findUserDto := &LoginUserDto{}
	err := ctx.ShouldBindQuery(findUserDto)
	if err != nil {
		ctx.JSON(400, home.ErrorResponse([]error{err}, home.CANNOT_PARSE_BODY))
		return
	}
	fmt.Println(findUserDto)
	// users, err := Search(findUserDto)
	// if err != nil {
	// 	ctx.JSON(500, home.ErrorResponse([]error{err}, home.USER_SERVER_ERROR))
	// 	return
	// }
	// ctx.JSON(200, users)
}
func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	// users, err := Search(findUserDto)
	// if err != nil {
	// 	ctx.JSON(500, home.ErrorResponse([]error{err}, home.USER_SERVER_ERROR))
	// 	return
	// }
	// ctx.JSON(200, users)
}
