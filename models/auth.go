package models

import "time"

type Token struct {
	Usrid string `json:"usrID"`
	TokenDate time.Duration `json:"tokenDate"`
	AccesToken string `json:"accessToken"`
	DetailUser []UsersRes `json:"detailUser"`
}

type UserLogin struct {
	Usrname string `json:"usrName"`
	Usrpassword string `json:"usrPassword"`
}

type VerifyToken struct {
	Raw string `json:"raw"`
	Valid bool `json:"valid"`
}