package user

import (
	"fmt"

	"github.com/amrebada/go-modules/core"
	"github.com/amrebada/go-modules/home"
	"github.com/gofiber/fiber/v2"
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
			SetHandlerFunc(GetUser).
			SetDescription("get user").
			SetResponseDto(&UserEntity{}))
}

// Login user
func RegisterUser(ctx *fiber.Ctx) error {
	registerDto := &RegisterDto{}
	err := ctx.BodyParser(registerDto)
	if err != nil {
		ctx.Status(400).JSON(home.ErrorResponse([]error{err}, home.CANNOT_PARSE_BODY))
		return err
	}
	user, err := Register(registerDto)
	if err != nil {
		ctx.Status(401).JSON(home.ErrorResponse([]error{err}, home.OAUTH_TOKEN_NOT_CORRECT))
		return err
	}
	ctx.Status(200).JSON(user)
	return nil
}

func LoginUser(ctx *fiber.Ctx) error {
	findUserDto := &LoginUserDto{}
	err := ctx.BodyParser(findUserDto)
	if err != nil {
		ctx.Status(400).JSON(home.ErrorResponse([]error{err}, home.CANNOT_PARSE_BODY))
		return err
	}
	fmt.Println(findUserDto)
	return nil
}

type QueryStruct struct {
	Q      string   `json:"q"`
	Filter []string `json:"filter"`
}
type ParamStruct struct {
	ID string `json:"id"`
}

func GetUser(ctx *fiber.Ctx) error {
	params := ParamStruct{}
	err := ctx.ParamsParser(&params)
	query := QueryStruct{}
	err = ctx.QueryParser(&query)
	if err != nil {
		return err
	}
	ctx.Status(200).JSON(map[string]interface{}{
		"path":  params,
		"query": query,
	})
	return nil
}
