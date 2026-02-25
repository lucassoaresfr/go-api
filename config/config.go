package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("ERRO INITIALIZING SQLITE: %v", err)
	}

	return nil
}

func GetSqLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

