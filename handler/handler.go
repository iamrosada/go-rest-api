package handler

import (
	"github.com/iamrosada/go-rest-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHendler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
