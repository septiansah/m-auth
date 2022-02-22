package user

import (
	"m-authentication/api/user/usecase"
	"m-authentication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UseCase usecase.IUserUsecase
}

func (r User) User(route *gin.RouterGroup) {
	route.POST("/user", r.CreateUser)
	route.GET("/user/shortProfile/:usrId", r.GetShortProfile)
}

func (r User) CreateUser(c *gin.Context) {
	var user models.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := r.UseCase.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (r User) GetShortProfile(c *gin.Context) {
	usrId := c.Param("usrId")
	user, err := r.UseCase.GetShortProfile(usrId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)

}
