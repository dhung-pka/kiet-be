package config

import (
	"os"

	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	appPort string

	// DB
	db     *gorm.DB
	dbName string
	dbPass string
	dbUser string
	dbPort string
	dbHost string

	// Token
	tokenAuth *jwtauth.JWTAuth
)

type ENV string

const (
	APP_PORT ENV = "APP_PORT"

	DB_NAME ENV = "DB_NAME"
	DB_PASS ENV = "DB_PASS"
	DB_USER ENV = "DB_USER"
	DB_PORT ENV = "DB_PORT"
	DB_HOST ENV = "DB_HOST"
)

func getEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	appPort = os.Getenv(string(APP_PORT))

	dbName = os.Getenv(string(DB_NAME))
	dbPass = os.Getenv(string(DB_PASS))
	dbUser = os.Getenv(string(DB_USER))
	dbPort = os.Getenv(string(DB_PORT))
	dbHost = os.Getenv(string(DB_HOST))

	return nil
}
