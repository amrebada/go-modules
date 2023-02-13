package company

import (
	"github.com/amrebada/go-modules/core"
	"github.com/amrebada/go-modules/modules/user"
	"gorm.io/gorm"
)

type CompanyUserEntity struct {
	core.EntityWithID
	UserID    string `gorm:"type:uuid;not null"`
	CompanyID string `gorm:"type:uuid;not null"`
	Role      string
}

func (c *CompanyUserEntity) GetTableName() string {
	return "company_user_entities"
}

func (c *CompanyUserEntity) GetByID() error {
	db := core.NewDatabase().DB
	res := db.First(&c, c.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CompanyUserEntity) GetByCompanyAndUser() error {
	db := core.NewDatabase().DB
	res := db.Where("user_id = ?", c.UserID).Where("company_id = ?", c.CompanyID).First(&c)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CompanyUserEntity) Create() error {
	db := core.NewDatabase().DB
	c.Prepare()
	res := db.Create(&c)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CompanyUserEntity) Update() error {
	db := core.NewDatabase().DB
	c.Prepare()
	res := db.Save(&c)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CompanyUserEntity) Prepare() {
	// c.Name = strings.ToUpper(string(c.Name[0])) + c.Name[1:]
}

func (c *CompanyUserEntity) ListBy(page int, limit int, db *gorm.DB) ([]CompanyUserEntity, error) {
	var companyUsers []CompanyUserEntity
	offset := (page - 1) * limit
	res := db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&companyUsers)
	if res.Error != nil {
		return nil, res.Error
	}
	return companyUsers, nil
}

func (c *CompanyUserEntity) CountBy(db *gorm.DB) (int64, error) {
	var count int64
	res := db.Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

func (c *CompanyUserEntity) Delete() error {
	db := core.NewDatabase().DB
	res := db.Delete(&c)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CompanyUserEntity) ListCompaniesByUserId(id string, page int, limit int) ([]CompanyEntity, int64, error) {
	var companies []CompanyEntity
	companyUsers, err := c.ListBy(page, limit,
		core.NewDatabase().DB.Where("user_id = ?", id))

	if err != nil {
		return companies, 0, err
	}
	total, err := c.CountBy(core.NewDatabase().DB.Where("user_id = ?", id))
	if err != nil {
		return companies, 0, err
	}
	ids := []string{}

	for _, companyUser := range companyUsers {
		ids = append(ids, companyUser.CompanyID)
	}
	res := core.NewDatabase().DB.Where("id IN (?)", ids).Find(&companies)
	if res.Error != nil {
		return companies, 0, res.Error
	}
	return companies, total, nil
}

func (c *CompanyUserEntity) ListUsersByCompanyId(id string, roles []string, page int, limit int) ([]user.UserEntity, int64, error) {
	var users []user.UserEntity
	companyUsers, err := c.ListBy(page, limit,
		core.NewDatabase().DB.Where("company_id = ?", id).Where("role IN (?)", roles))

	if err != nil {
		return users, 0, err
	}
	total, err := c.CountBy(core.NewDatabase().DB.Where("company_id = ?", id).Where("role IN (?)", roles))
	if err != nil {
		return users, 0, err
	}
	ids := []string{}
	for _, companyUser := range companyUsers {
		ids = append(ids, companyUser.UserID)
	}

	res := core.NewDatabase().DB.Where("id IN (?)", ids).Find(&users)
	if res.Error != nil {
		return users, 0, res.Error
	}
	return users, total, nil
}
