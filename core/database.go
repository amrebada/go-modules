package core

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	dsn string
	DB  *gorm.DB
}

var database *Database

func NewDatabase() *Database {
	if database == nil {
		config := ConfigInstance()
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.Database.HOST, config.Database.PORT, config.Database.USER, config.Database.DBNAME, config.Database.PASSWORD)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic("failed to connect database")
		}
		fmt.Println(connect, "Connected to database")
		db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
		database = &Database{
			dsn: dsn,
			DB:  db,
		}
	}
	return database
}

type EntityWithID struct {
	ID        string    `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}
