package repository

import (
	"errors"
	"fmt"
	"m-authentication/config"
	"m-authentication/models"

	"github.com/jmoiron/sqlx"
)

type UserSQL struct {
	db *sqlx.DB
}

func NewUserSQL(db *config.DB) IUserSql {
	return &UserSQL{
		db: db.SQL,
	}
}

func (usr UserSQL) CreateUser(user models.Users) (models.Users, error) {
	fmt.Println("USER ", user)
	tx := usr.db.MustBegin()
	query, err := tx.NamedQuery(`INSERT INTO public.users(
		usrid, usrname, usremail, usrpassword, usrfirstname, usrlastname, usrphonenum, usrdateofbirth, usrgender, usraddress, usrpostal, usrcity, usrcreated)
		VALUES (:usrid, :usrname, :usremail, :usrpassword, :usrfirstname, :usrlastname, :usrphonenum, :usrdateofbirth, :usrgender, :usraddress, :usrpostal, :usrcity, :usrcreated) 
		RETURNING id, usrname, usremail, usrpassword, usrfirstname, usrlastname, usrphonenum, usrdateofbirth, usrgender, usraddress, usrpostal, usrcity, usrcreated`,
		map[string]interface{}{
			"usrid":          user.Usrid,
			"usrname":        user.Usrname,
			"usremail":       user.Usremail,
			"usrpassword":    user.Usrpassword,
			"usrfirstname":   user.Usrfirstname,
			"usrlastname":    user.Usrlastname,
			"usrphonenum":    user.Usrphonenum,
			"usrdateofbirth": user.Usrdateofbirth,
			"usrgender":      user.Usrgender,
			"usraddress":     user.Usraddress,
			"usrpostal":      user.Usrpostal,
			"usrcity":        user.Usrcity,
			"usrcreated":     user.Usrcreated,
		})

	if err != nil {
		tx.Rollback()
		return user, err
	}
	for query.Next() {
		if err := query.StructScan(&user); err != nil {
			tx.Rollback()
			return user, err
		}
	}

	tx.Commit()
	return user, nil
}

func (usr UserSQL) CheckUser(userName string) (bool, error) {
	var usrName string
	if err := usr.db.Get(&usrName, `SELECT usrname FROM users WHERE usrname = $1`, userName); err != nil {
		return false, errors.New("Username is not registered")
	}

	return true, nil
}

func (usr UserSQL) GetUserByUsrname(usrname string) (models.UsersRes, error) {
	var user models.UsersRes
	err := usr.db.Get(&user, `SELECT id, usrid, usrname, usremail, usrpassword, usrfirstname, usrlastname, usrphonenum, usrdateofbirth, usrgender, usraddress, usrpostal, usrcity, usrtype, usrcreated
	FROM public.users WHERE usrname = $1`, usrname)
	if err != nil {
		return models.UsersRes{}, err
	}

	return user, nil
}

func (usr UserSQL) GetUserByUsrId(userId string) (models.UsersRes, error) {
	var user models.UsersRes
	err := usr.db.Get(&user, `SELECT id, usrid, usrname, usremail, usrpassword, usrfirstname, usrlastname, usrphonenum, usrdateofbirth, usrgender, usraddress, usrpostal, usrcity, usrtype, usrcreated
	FROM public.users WHERE usrid = $1`, userId)
	if err != nil {
		return models.UsersRes{}, err
	}

	return user, nil
}
