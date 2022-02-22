package usecase

import (
	"errors"
	"fmt"
	"m-authentication/api/user/repository"
	"m-authentication/handlers"
	"m-authentication/models"
	"m-authentication/utils"
	"net/http"
)

type authUseCase struct {
	userSql repository.IUserSql
}

func NewAuthUseCase(userSql repository.IUserSql) IAuthUsecase {
	return &authUseCase{
		userSql: userSql,
	}
}

func (auth authUseCase) Login(user models.UserLogin) (models.Token, error) {

	var Token models.Token

	checkUsr, err := auth.userSql.CheckUser(user.Usrname)
	if err != nil {
		fmt.Println(err)
		return Token, err
	}
	if checkUsr {
		usr, err := auth.userSql.GetUserByUsrname(user.Usrname)
		if err != nil {
			return Token, err
		}
		verifyPass, _ := handlers.CompareHashedAndSalted(usr.Usrpassword, []byte(user.Usrpassword))

		if verifyPass {

			AccesToken, err := utils.CreateToken(usr)
			if err != nil {
				return Token, err
			}

			Token.Usrid = usr.Usrid
			Token.AccesToken = AccesToken
			Token.DetailUser = append(Token.DetailUser, usr)
		} else {
			return Token, errors.New("Password is incorrect")
		}
	}

	return Token, nil
}

func (r authUseCase) VerifyToken(req *http.Request) (models.VerifyToken, error) {
	// var tokenVerf models.VerifyToken
	fmt.Println(req)
	return models.VerifyToken{}, nil
}
