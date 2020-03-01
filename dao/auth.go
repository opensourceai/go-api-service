package dao

type AuthDao interface {
	CheckAuth(username, password string) (bool, error)
}
