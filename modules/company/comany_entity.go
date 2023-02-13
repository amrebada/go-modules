package company

import (
	"errors"
	"strings"

	"github.com/amrebada/go-modules/core"
	"github.com/amrebada/go-modules/modules/user"
	"gorm.io/gorm"
)

type CompanyEntity struct {
	core.EntityWithID
	Name        string `json:"name" gorm:"not null;unique"`
	Image       string `json:"image" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}

func (c *CompanyEntity) GetTableName() string {
	return "company_entities"
}

func (c *CompanyEntity) GetByID() error {
	db := core.NewDatabase().DB
	res := db.First(&c, c.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CompanyEntity) GetByName() error {
	db := core.NewDatabase().DB
	res := db.Where("name = ?", strings.ToLower(c.Name)).First(&c)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CompanyEntity) Create(userId string) error {
	db := core.NewDatabase().DB
	c.Prepare()
	res := db.Create(&c)
	if res.Error != nil {
		return res.Error
	}
	user := &user.UserEntity{
		EntityWithID: core.EntityWithID{
			ID: userId,
		},
	}
	err := user.GetByID()
	if err != nil || user.Email == "" {
		c.Delete()
		return errors.New("user not found")
	}
	companyUser := &CompanyUserEntity{
		UserID:    userId,
		CompanyID: c.ID,
		Role:      ROLE_COMPANY_ADMIN,
	}
	err = companyUser.Create()
	if err != nil {
		c.Delete()
		return err
	}

	return nil
}

func (c *CompanyEntity) Update() error {
	db := core.NewDatabase().DB
	c.Prepare()
	res := db.Save(&c)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CompanyEntity) Prepare() {
	// c.Name = strings.ToUpper(string(c.Name[0])) + c.Name[1:]
}

func (c *CompanyEntity) ListBy(page int, limit int, db *gorm.DB) ([]CompanyEntity, error) {
	var users []CompanyEntity
	offset := (page - 1) * limit
	res := db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func (c *CompanyEntity) CountBy(db *gorm.DB) (int64, error) {
	var count int64
	res := db.Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

func (c *CompanyEntity) Delete() error {
	db := core.NewDatabase().DB
	res := db.Delete(&c)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CompanyEntity) SearchByName(name string, page int, limit int) ([]CompanyEntity, int64, error) {
	var companies []CompanyEntity
	companies, err := c.ListBy(page, limit,
		core.NewDatabase().DB.Where("name ILIKE ?", "%"+name+"%"))
	if err != nil {
		return nil, 0, err
	}
	count, err := c.CountBy(core.NewDatabase().DB.Where("name ILIKE ?", "%"+name+"%"))
	if err != nil {
		return nil, 0, err
	}
	return companies, count, nil
}
