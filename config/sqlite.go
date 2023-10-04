package config

import (
	"os"

	"github.com/iamrosada/go-rest-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	//Check if file database exist
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Infof("database file not, creating")
		//Create database file and directory
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}
		file.Close()

	}

	//Create db and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Sqlite oportunity error: %v", err)
		return nil, err
	}

	//Migrate the schemas
	err = db.AutoMigrate(&schemas.Oportunity{})

	if err != nil {
		logger.Errorf("Sqlite automigration error: %v", err)
		return nil, err
	}
	//Return DB
	return db, nil
}
