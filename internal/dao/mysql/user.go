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

package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/pkg/errors"
)

func NewUserDao(db *gorm.DB) (dao.UserDao, error) {
	return &userDao{db: db}, nil
}

type userDao struct {
	db *gorm.DB
}

func (dao *userDao) DaoFindByIds(ids ...int) (users []models.User, err error) {
	// 查询ids用户
	err = dao.db.Where(ids).Find(users).Error
	return
}

func (dao *userDao) DaoAdd(user *models.User) error {
	// 防止主键ID被人为更新
	user.ID = 0
	dao.db.Create(user)
	return nil
}

func (dao *userDao) DaoEdit(user *models.User) error {
	userQuery := models.User{}
	userQuery.ID = user.ID

	if err := dao.db.First(&userQuery).Error; err == gorm.ErrRecordNotFound {
		return errors.New("数据不存在")
	}
	dao.db.Save(&user)
	return nil
}

func (dao *userDao) DaoGetUserByUsername(username string) (err error, user models.User) {
	if err := dao.db.Where("username = ?", username).First(&user).Error; err != nil {
		return err, user
	}
	return nil, user

}

func (dao *userDao) DaoDeleteById(ids ...int) error {
	user := models.User{}
	// 查询用户id是否存在
	for _, id := range ids {
		if err := dao.db.Where("id = ?", id).First(&user).Error; err != nil {
			return err
		}
	}

	// 软删除
	var err error = nil
	// 事务机制，出错便回滚
	err = dao.db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if e := dao.db.Where("id = ?", id).Delete(&models.User{}).Error; e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
