package usecase

import (
	"m-authentication/models"
	"net/http"
)

type IAuthUsecase interface {
	Login(models.UserLogin) (models.Token, error)
	VerifyToken(req *http.Request) (models.VerifyToken, error)
}