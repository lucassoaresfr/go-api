package handler

import (
	"github.com/lucassoaresfr/go-api.git/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandle() {
	logger = config.GetLogger("handler")
	db = config.GetSqLite()
}
