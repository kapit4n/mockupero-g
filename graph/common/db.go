package common

import (
	models "stats-mockupero/graph/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() (*gorm.DB, error) {
	var err error

	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Todo{})
	db.AutoMigrate(&models.Project{})

	return db, nil
}
