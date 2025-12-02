package config

import (
	"os"

	"github.com/alephjunio/go-opportunities-api/schemas"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitializationSqlite() (*gorm.DB, error) {

	logger := GetLogger("sqlite")
	dbPath := "./db/oppotunities.db"
	//check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("database file not found, creating....")
		// create the database file and directory
		err := os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite Opening error - %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite AutoMigrate error - %v", err)
		return nil, err
	}

	return db, nil

}
