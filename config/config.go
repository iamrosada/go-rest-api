package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	// Initialize logger
	loger := NewLogger(prefix)

	return loger
}
