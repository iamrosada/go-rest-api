package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}
func GetLogger(prefix string) *Logger {
	// Initialize logger
	logger = NewLogger(prefix)

	return logger
}
