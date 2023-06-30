package config

import (
	"fmt"
	"log"
	"service/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDatabase(initTable bool) error {
	dns := fmt.Sprintf(`
		host=%s 
		user=%s 
		password=%s 
		dbname=%s 
		port=%s 
		sslmode=disable`,
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
	)

	var err error
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}

	if initTable {
		log.Println("Start migrate")
		errInit := db.AutoMigrate(
			&model.Credential{},
			&model.District{},
			&model.Profile{},
			&model.Province{},
		)

		if errInit != nil {
			return errInit
		}

		log.Println("Finish migrate")
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
