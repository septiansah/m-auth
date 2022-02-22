package utils

import (
	"log"
	"m-authentication/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func CreateToken(user models.UsersRes)(string, error){
	jwtClaims := jwt.MapClaims{
		"access_uuid" : uuid.New().String(),
		"usrID" : user.Usrid,
		"usrName" : user.Usrname,
		"usrEmail" : user.Usremail,
		"userFirstName" : user.Usrfirstname,
		"userLastName" : user.Usrlastname,
		"exp" : time.Now().Add(time.Minute * 60 * 12).Unix(),
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	accessToken, err := Token.SignedString([]byte("secret"))
	if err != nil {
		log.Panicln(err)
	}

	return accessToken, nil
}