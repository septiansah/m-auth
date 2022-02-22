package repository

import (
	"m-authentication/config"

	"github.com/jmoiron/sqlx"
)

type AuthSQL struct {
	db *sqlx.DB
}

func NewAuthSQL(db *config.DB) IAuthRepoInterface {
	return &AuthSQL{
		db: db.SQL,
	}
}

func (auth AuthSQL) StoreToken(user_id, token string) error {
	 
	if _, err := auth.db.NamedQuery(`INSERT INTO user_tokens(
		user_id, token)
		VALUES (:user_id, :token)`, map[string]interface{}{"user_id": user_id, "token" : token}); err != nil {
		return nil
	}

	return nil
}