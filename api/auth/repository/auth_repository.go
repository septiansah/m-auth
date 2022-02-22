package repository

type IAuthRepoInterface interface {
	StoreToken(user_id, token string) error
}