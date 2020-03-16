package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/pkg/errors"
)

type userDao struct {
	*gorm.DB
}

func NewUserDao(db *gorm.DB) (dao.UserDao, error) {
	return &userDao{DB: db}, nil
}
func (dao *userDao) DaoAdd(user *models.User) error {
	// 防止主键ID被人为更新
	user.ID = 0
	dao.Create(user)
	return nil
}

func (dao *userDao) DaoEdit(user *models.User) error {
	userQuery := models.User{}
	userQuery.ID = user.ID

	if err := dao.First(&userQuery).Error; err == gorm.ErrRecordNotFound {
		return errors.New("数据不存在")
	}
	dao.Save(&user)
	return nil
}

func (dao *userDao) DaoGetUserByUsername(username string) (err error, user models.User) {
	if err := dao.Where("username = ?", username).First(&user).Error; err != nil {
		return err, user
	}
	return nil, user

}

func (dao *userDao) DaoDeleteById(ids ...int) error {
	user := models.User{}
	// 查询用户id是否存在
	for _, id := range ids {
		if err := dao.Where("id = ?", id).First(&user).Error; err != nil {
			return err
		}
	}

	// 软删除
	var err error = nil
	// 事务机制，出错便回滚
	err = dao.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if e := dao.Where("id = ?", id).Delete(&models.User{}).Error; e != nil {
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
