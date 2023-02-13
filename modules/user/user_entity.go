package user

import (
	"strings"

	"github.com/amrebada/go-modules/core"
	"gorm.io/gorm"
)

type UserEntity struct {
	core.EntityWithID
	Name     string `json:"name" gorm:"not null"`
	LastName string `json:"lastName"`
	Email    string `json:"email" gorm:"unique;index;not null"`
	Bio      string `json:"bio" gorm:"type:text"`
	Image    string `json:"image"`
	LoggedBy string `json:"loggedBy"`
}

func (u *UserEntity) GetTableName() string {
	return "user_entities"
}

func (u *UserEntity) GetByID() error {
	db := core.NewDatabase().DB
	res := db.First(&u, u.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserEntity) GetByEmail() error {
	db := core.NewDatabase().DB
	res := db.Where("email = ?", strings.ToLower(u.Email)).First(&u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserEntity) Create() error {
	db := core.NewDatabase().DB
	u.Prepare()
	res := db.Create(&u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserEntity) Update() error {
	db := core.NewDatabase().DB
	u.Prepare()
	res := db.Save(&u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserEntity) Prepare() {
	u.Email = strings.ToLower(u.Email)
}

func (u *UserEntity) ListBy(page int, limit int, db *gorm.DB) ([]UserEntity, error) {
	var users []UserEntity
	offset := (page - 1) * limit
	res := db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func (u *UserEntity) CountBy(db *gorm.DB) (int64, error) {
	var count int64
	res := db.Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

func (u *UserEntity) Delete() error {
	db := core.NewDatabase().DB
	res := db.Delete(&u)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserEntity) SearchByName(name string) ([]UserEntity, error) {
	var users []UserEntity
	users, err := u.ListBy(1, 50,
		core.NewDatabase().DB.Where("name ILIKE ?", "%"+name+"%").Statement.Or("last_name ILIKE ?", "%"+name+"%"))
	if err != nil {
		return nil, err
	}
	return users, nil
}
