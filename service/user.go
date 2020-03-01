package service

import (
	"github.com/opensourceai/go-api-service/dao"
	"github.com/opensourceai/go-api-service/dao/mysql"
	"github.com/opensourceai/go-api-service/models"
)

type UserService interface {
	Register(user *models.User) error
}

type UserServiceImpl struct{}

var userDao dao.UserDao

func init() {
	userDao = new(mysql.UserDaoImpl)
}
func (UserServiceImpl) Register(user *models.User) error {
	return userDao.Add(user)
}
