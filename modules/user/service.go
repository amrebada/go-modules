package user

import (
	"time"

	"github.com/amrebada/go-modules/core"
	"github.com/golang-jwt/jwt"
)

func Register(registerDto *RegisterDto) (*RegisterResponseDto, error) {
	//according to type get user info
	//check if user exist
	user := &UserEntity{
		Name:     "Amr",
		LastName: "Ebada",
		Email:    "amr.app.engine@gmail.com",
		Bio:      "I am a developer",
		Image:    "http://dkkdopkeq",
		LoggedBy: registerDto.Type,
	}
	user.GetByEmail()
	if user.ID != "" {
		token, err := generateToken(user)
		if err != nil {
			return nil, err
		}
		return &RegisterResponseDto{
			Token: token,
			User:  *user,
		}, nil
	}
	//if not create user
	err := user.Create()
	if err != nil {
		return nil, err
	}
	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}
	//if exist return it
	return &RegisterResponseDto{
		Token: token,
		User:  *user,
	}, nil
}

func Search(findUserDto *FindUserDto) (*FindUserResponseDto, error) {
	user := &UserEntity{}
	users, err := user.SearchByName(findUserDto.Name)
	if err != nil {
		return &FindUserResponseDto{
			Users: []UserEntity{},
			Total: 0,
		}, nil
	}

	return &FindUserResponseDto{
		Users: users,
		Total: int64(len(users)),
	}, nil
}

func generateToken(user *UserEntity) (string, error) {
	secret := core.ConfigInstance().JWT_SECRET
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["image"] = user.Image
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	return token.SignedString([]byte(secret))
}
