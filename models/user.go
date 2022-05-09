package models

import "time"

type Users struct {
	ID             int    `json:"ID"`
	Usrid          string `json:"usrID"`
	Usrname        string `json:"username"`
	Usremail       string `json:"userEmail"`
	Usrpassword    string `json:"userPassword"`
	Usrfirstname   string `json:"userFirstName"`
	Usrlastname    string `json:"userLastName"`
	Usrphonenum    string `json:"usrPhoneNumber"`
	Usrdateofbirth string `json:"usrDateOfBirth"`
	Usrgender      string `json:"usrGender"`
	Usraddress     string `json:"usrAddress"`
	Usrpostal      string `json:"usrPostal"`
	Usrcity        string `json:"usrCity"`
	Created_date   int    `json:"createdDate"`
}

type ShortProfile struct {
	Usrid        string `json:"usrId"`
	Usrname      string `json:"usrName"`
	Usrfirstname string `json:"usrFirstName"`
	Usraddress   string `json:"usrAddress"`
}

type UsersRes struct {
	ID             int       `json:"ID"`
	Usrid          string    `json:"usrID"`
	Usrname        string    `json:"username"`
	Usremail       string    `json:"userEmail"`
	Usrpassword    string    `json:"userPassword"`
	Usrfirstname   string    `json:"userFirstName"`
	Usrlastname    string    `json:"userLastName"`
	Usrphonenum    string    `json:"usrPhoneNumber"`
	Usrdateofbirth string    `json:"usrDateOfBirth"`
	Usrgender      string    `json:"usrGender"`
	Usraddress     string    `json:"usrAddress"`
	Usrpostal      string    `json:"usrPostal"`
	Usrcity        string    `json:"usrCity"`
	Usrtype        *string   `json:"usrType"`
	Usrcreated     time.Time `json:"usrCreated"`
}
