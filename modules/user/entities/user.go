package entities

import (
	"fmt"
	"strings"

	"github.com/amrebada/go-modules/modules/user/config"
	core "github.com/amrebada/neon-core"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PublicData struct {
	core.EntityWithID
	FirstName string `json:"firstName" gorm:"not null" validate:"required,max=100"`
	LastName  string `json:"lastName" validate:"max=100"`
	Email     string `json:"email" gorm:"unique;index" validate:"email"`
	Phone     string `json:"phone" gorm:"unique;index" validate:"phone"`
	Image     string `json:"image" gorm:"default:'https://www.gravatar.com/avatar/" validate:"url"`
}

type User struct {
	PublicData
	LoggedBy config.AuthenticationProvider `json:"loggedBy"`
	Password string                        `json:"password" validate:"min=3,max=100"`
}

func NewUserRepository(db *gorm.DB) *User {
	return &User{
		PublicData: PublicData{EntityWithID: core.EntityWithID{
			DB: db,
		}},
	}
}

func (u *User) Validate() []error {
	var errors []error
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err)
		}
	}
	return errors
}

func (u *User) GetTableName() string {
	return "module_user_users"
}

func (u *User) GetByID() error {

	res := u.DB.First(&u, u.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *User) GetByEmail() error {
	res := u.DB.Where("email = ?", strings.ToLower(u.Email)).First(&u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *User) Create() error {
	u.Prepare()
	res := u.DB.Create(&u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *User) Update() error {

	u.Prepare()
	res := u.DB.Save(&u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *User) Prepare() {
	u.Email = strings.ToLower(u.Email)
	if u.Image == "" {
		u.Image = fmt.Sprintf("https://ui-avatars.com/api/?name=%s", u.FirstName+"+"+u.LastName)
	}
}

func (u *User) listBy(page int, limit int, statement *gorm.DB) ([]User, error) {
	var users []User
	offset := (page - 1) * limit
	res := statement.Limit(limit).Offset(offset).Order("created_at DESC").Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func (u *User) CountBy() (int64, error) {
	var count int64
	res := u.DB.Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

func (u *User) Delete() error {
	res := u.DB.Delete(&u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *User) SearchByName(name string) ([]User, error) {
	var users []User
	users, err := u.listBy(1, 50,
		u.DB.Where("name ILIKE ?", "%"+name+"%").Statement.Or("last_name ILIKE ?", "%"+name+"%"))
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) GetPublicData() map[string]any {
	return map[string]any{
		"id":        u.ID,
		"firstName": u.FirstName,
		"lastName":  u.LastName,
		"email":     u.Email,
		"phone":     u.Phone,
		"image":     u.Image,
	}
}
