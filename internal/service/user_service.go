/*
 *    Copyright 2020 opensourceai
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package service

import (
	"errors"
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/api/v1/dto"
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/dao/mysql"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/util"
)

//NewUserService方法依赖mysql.NewUserDao，需要传入实现了gorm方法的结构体
var ProviderUser = wire.NewSet(NewUserService, mysql.NewUserDao)

//参数需要一个注入了UserDao接口的结构体,返回一个实现了UserDao接口的业务结构体
func NewUserService(dao2 dao.UserDao) (UserService, error) {
	return &userService{dao2}, nil
}

type UserService interface {
	// 用户注册
	ServiceRegister(user *models.User) error
	// 用户登录
	ServiceLogin(user models.User) (*models.User, bool, error)
	// 修改用户密码
	// Deprecated: 接口重复
	ServiceUpdatePwd(username string, s string) error
	// 修改用户信息
	// Deprecated: 接口重复
	ServiceUpdateMsg(username string, user *models.User) error
	// 修改用户信息
	ServiceUpdateUser(onlineUser *app.Auth, user *dto.UserDTO) (err error)
}

type userService struct {
	userDao dao.UserDao
}

func (service userService) ServiceUpdateUser(onlineUser *app.Auth, user *dto.UserDTO) (err error) {
	if onlineUser.UserId != user.ID {
		return errors.New("用户不存在")
	}
	if users, err := service.userDao.DaoFindByIds(user.ID); err != nil || len(users) != 1 {
		return errors.New("不存在")
	}
	// TODO:检验用户头像地址是否存在
	u := &models.User{
		Model:       models.Model{ID: user.ID},
		Name:        user.Name,
		Password:    util.EncodeMD5(user.Password),
		Description: user.Description,
		Sex:         user.Sex,
		AvatarSrc:   user.AvatarSrc,
		Email:       user.Email,
		WebSite:     user.WebSite,
		Company:     user.Company,
		Position:    user.Password,
	}
	return service.userDao.DaoUpdate(u)
}

func (service userService) ServiceRegister(user *models.User) error {
	// 加密密码
	user.Password = util.EncodeMD5(user.Password)
	return service.userDao.DaoAdd(user)
}

func (service userService) ServiceLogin(user models.User) (*models.User, bool, error) {
	err, u := service.userDao.DaoGetUserByUsername(user.Username)
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

func (service userService) ServiceUpdatePwd(username string, s string) error {
	//通过用户名从数据库获取用户对象
	_, u := service.userDao.DaoGetUserByUsername(username)
	//修改密码
	u.Password = s
	//调用修改用户信息方法将对象重新写入数据库，有错误就返回错误
	return service.userDao.DaoEdit(&u)
}

func (service userService) ServiceUpdateMsg(username string, user *models.User) error {
	//通过用户名从数据库获取用户对象
	_, u := service.userDao.DaoGetUserByUsername(username)
	//修改内容
	//// 描述
	u.Description = user.Description
	//Description string `json:"description" grom:"column:description" valid:"MaxSize(200)"`
	//// 性别
	u.Sex = user.Sex
	//Sex int `json:"sex" grom:"column:sex;not null" valid:"Min(1)"`
	//// 头像地址
	u.AvatarSrc = user.AvatarSrc
	//AvatarSrc string `json:"avatar_src" grom:"column:avatar_src;not null"`
	//// 电子邮件
	u.Email = user.Email
	//Email string `json:"email" grom:"column:email" valid:"Required;Email;MaxSize(100)"`
	//// 网站
	u.WebSite = user.WebSite
	//WebSite string `json:"web_site" grom:"column:web_site" valid:"MaxSize(50)"`
	//// 公司
	u.Company = user.Company
	//Company string `json:"company" grom:"column:company" valid:"MaxSize(50)"`
	//// 职位
	u.Position = user.Position
	//Position string `json:"position" grom:"column:position" valid:"MaxSize(50)"`
	//调用修改用户信息方法将对象重新写入数据库，有错误就返回错误
	return service.userDao.DaoEdit(&u)
}
