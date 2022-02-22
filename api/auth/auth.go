package auth

import (
	"fmt"
	"m-authentication/api/auth/usecase"
	"m-authentication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	UseCase usecase.IAuthUsecase
}

func (r Auth) Auth(route *gin.RouterGroup) {
	route.POST("/login", r.Login)
	route.POST("/verifyToken", r.Login)
}

func (r Auth) Login(c *gin.Context) {
	var auth models.UserLogin
	if err := c.BindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	payload, err := r.UseCase.Login(auth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payload)
}

func (r Auth) VerifyToken(c *gin.Context) {
	// var verifyToken models.VerifyToken
	// token, err := r.UseCase.VerifyToken(c.Request)
	// if err != nil {
	// 	log.Panicln(err)
	// }
	fmt.Println(c.Request)

	c.JSON(http.StatusOK, "tes")
}
