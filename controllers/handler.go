package controllers

import (
	"github.com/DamiaoCanndido/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() (*config.Logger, *gorm.DB) {
	logger = config.GetLogger("handler")
	db = config.GetPostgreSQL()
	return logger, db
}
