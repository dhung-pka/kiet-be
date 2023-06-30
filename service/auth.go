package service

import (
	"errors"
	"service/config"
	"service/model"
	"service/repo"

	"gorm.io/gorm"
)

type authService struct {
	db       *gorm.DB
	authRepo repo.AuthRepo
}

type AuthService interface {
	Login(info model.LoginRequest) (model.Credential, error)
}

func (a *authService) Login(info model.LoginRequest) (model.Credential, error) {
	var accountUser model.Credential

	if err := a.db.Model(&model.Credential{}).
		Preload("Profile").
		Where("username = ?", info.Username).
		First(&accountUser).Error; err != nil {
		return accountUser, err
	}

	checkPassworld := a.authRepo.CheckPasswordHash(info.Passworld, accountUser.Password)

	if !checkPassworld {
		return accountUser, errors.New("Error passworld")
	}

	return accountUser, nil
}

func NewAuthService() AuthService {
	db := config.GetDB()
	authRepo := repo.NewAuthRepo()
	return &authService{
		db:       db,
		authRepo: authRepo,
	}
}
