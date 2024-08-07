package helpers

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"rexai.com/models"
)

var (
	sqliteInstance *SqliteHelpers
	sqlOnce        sync.Once
)

func GetDbInstance() *SqliteHelpers {
	sqlOnce.Do(func() {
		sqliteInstance = &SqliteHelpers{}
		sqliteInstance.New()
		sqliteInstance.GetDb().AutoMigrate(&models.Image{})
	})
	return sqliteInstance
}

type SqliteHelpers struct {
	db *gorm.DB
}

func (helper *SqliteHelpers) New() {
	instance, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database!", err)
	}
	helper.db = instance
}

func (helper *SqliteHelpers) GetDb() *gorm.DB {
	return helper.db
}
