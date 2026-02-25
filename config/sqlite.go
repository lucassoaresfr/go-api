package config

import (
	"os"

	"github.com/lucassoaresfr/go-api.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbpath := "./db/main.db"

	_, err := os.Stat(dbpath)

	if os.IsNotExist(err) {
		logger.Info("DATABASE NOT FOUND, CREATING...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbpath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})

	if err != nil {
		logger.ErrorF("SQL OPENING ERROR: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.ErrorF("SQLITE AUTOMIGRATION ERROR: %v", err)
		return nil, err
	}
	return db, nil
}
