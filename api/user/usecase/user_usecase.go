package usecase

import "m-authentication/models"

type IUserUsecase interface {
	CreateUser(models.Users) (models.Users, error)
	GetShortProfile(userId string) (models.ShortProfile, error)
}
