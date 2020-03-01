package service

import (
	"github.com/opensourceai/go-api-service/dao"
	"github.com/opensourceai/go-api-service/dao/mysql"
)

type Auth struct {
	Username string
	Password string
}

var authDao dao.AuthDao

func init() {
	// 注入
	authDao = new(mysql.AuthDaoImpl)
}

type AuthService interface {
	Check() (bool, error)
}

func (a *Auth) Check() (bool, error) {
	return authDao.CheckAuth(a.Username, a.Password)
}
