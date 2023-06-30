package repo

import (
	"service/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

type AuthRepo interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

func (a *authRepo) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (a *authRepo) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewAuthRepo() AuthRepo {
	db := config.GetDB()
	return &authRepo{
		db: db,
	}
}
