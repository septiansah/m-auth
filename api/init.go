package api

import (
	"fmt"
	"log"
	"m-authentication/api/auth"
	authUseCase "m-authentication/api/auth/usecase"
	"m-authentication/api/user"
	userRepo "m-authentication/api/user/repository"
	userUseCase "m-authentication/api/user/usecase"
	"m-authentication/config"
	"m-authentication/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	db, err := config.ConnectSQL()
	if err != nil {
		log.Fatal("error db connection ", err)
	}
	r.GET("", func(c *gin.Context) {
		return
	})

	r.Use(utils.CORSMiddleware())
	v1 := r.Group("/api")

	iUserSql := userRepo.NewUserSQL(db)
	iUserUseCase := userUseCase.NewUserUseCase(iUserSql)
	UserController := user.User{UseCase: iUserUseCase}
	UserController.User(v1)

	iAuthUseCase :=  authUseCase.NewAuthUseCase(iUserSql)
	AuthController := auth.Auth{UseCase: iAuthUseCase}
	AuthController.Auth(v1)

	r.Run(fmt.Sprintf(os.Getenv("PORT")))
}
