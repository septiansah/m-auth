package usecase

import (
	"fmt"
	"m-authentication/api/user/repository"
	"m-authentication/models"
	"m-authentication/utils"

	"github.com/google/uuid"
)

type userUseCase struct {
	userSql repository.IUserSql
}

func NewUserUseCase(userSql repository.IUserSql) IUserUsecase {
	return &userUseCase{
		userSql: userSql,
	}
}

func (usr userUseCase) CreateUser(user models.Users) (models.Users, error) {

	hassPass, err := utils.HashAndSalt([]byte(user.Usrpassword))
	if err != nil {
		return user, err
	}
	user.Usrid = uuid.New().String()
	user.Usrpassword = hassPass
	user, err = usr.userSql.CreateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (usr userUseCase) GetShortProfile(userId string) (models.ShortProfile, error) {

	user, err := usr.userSql.GetUserByUsrId(userId)
	if err != nil {
		return models.ShortProfile{}, err
	}

	profile := &models.ShortProfile{
		Usrid:        user.Usrid,
		Usrname:      user.Usrname,
		Usrfirstname: user.Usrfirstname,
		Usraddress:   fmt.Sprintf("%v, %v", user.Usraddress, user.Usrcity),
	}

	return *profile, nil
}
