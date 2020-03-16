package service

import (
	"errors"
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/dao/mysql"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/util"
)

type UserService interface {
	Register(user *models.User) error
	Login(user models.User) (*models.User, bool, error)
}

type userService struct {
	dao.UserDao
}

var ProviderUser = wire.NewSet(NewUserService, mysql.NewUserDao)

func NewUserService(dao2 dao.UserDao) (UserService, error) {
	return &userService{dao2}, nil

}
func (service userService) Register(user *models.User) error {
	// 加密密码
	user.Password = util.EncodeMD5(user.Password)
	return service.DaoAdd(user)
}

func (service userService) Login(user models.User) (*models.User, bool, error) {
	err, u := service.DaoGetUserByUsername(user.Username)
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
