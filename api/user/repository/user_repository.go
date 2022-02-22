package repository

import "m-authentication/models"

type IUserSql interface {
	CreateUser(models.Users) (models.Users, error)
	CheckUser(userName string) (bool, error)
	GetUserByUsrname(userName string) (models.UsersRes, error)
	GetUserByUsrId(userId string) (models.UsersRes, error)
}
