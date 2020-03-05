package service

import (
	"errors"
	"github.com/opensourceai/go-api-service/dao"
	"github.com/opensourceai/go-api-service/dao/mysql"
	"github.com/opensourceai/go-api-service/models"
	"github.com/opensourceai/go-api-service/pkg/util"
)

type UserService interface {
	Register(user *models.User) error
	Login(user models.User) (*models.User, bool, error)
}

type UserServiceImpl struct{}

var userDao dao.UserDao

func init() {
	userDao = new(mysql.UserDaoImpl)
}
func (UserServiceImpl) Register(user *models.User) error {
	// 加密密码
	user.Password = util.EncodeMD5(user.Password)
	return userDao.Add(user)
}

func (i UserServiceImpl) Login(user models.User) (*models.User, bool, error) {
	err, u := userDao.GetUserByUsername(user.Username)
	if err != nil {
		return nil, false, errors.New("登录失败")
	}
	// 匹配密码
	md5Password := util.EncodeMD5(user.Password)
	if u.Password == md5Password {
		return &u, true, nil
	}
	return nil, false, errors.New("登录失败")
}
