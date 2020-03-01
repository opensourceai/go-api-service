package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/models"
	"github.com/pkg/errors"
)

type UserDaoImpl struct{}

func (u UserDaoImpl) Add(user *models.User) error {
	if db.NewRecord(&user) {
		return errors.New("主键已存在")
	}
	db.Create(user)
	return nil
}

func (u UserDaoImpl) Edit(user *models.User) error {
	userQuery := models.User{}
	userQuery.ID = user.ID

	if err := db.First(&userQuery).Error; err == gorm.ErrRecordNotFound {
		return errors.New("数据不存在")
	}
	db.Save(&user)
	return nil
}

func (u UserDaoImpl) GetUserByUsername(username string) (err error, user models.User) {
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return err, user
	}
	return nil, user

}

func (u UserDaoImpl) DeleteById(ids ...int) error {
	user := models.User{}
	// 查询用户id是否存在
	for _, id := range ids {
		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			return err
		}
	}

	// 软删除
	var err error = nil
	// 事务机制，出错便回滚
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if e := db.Where("id = ?", id).Delete(&models.User{}).Error; e != nil {
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
