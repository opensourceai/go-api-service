package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/models"
)

type AuthDaoImpl struct{}

func (AuthDaoImpl) CheckAuth(username, password string) (bool, error) {
	var auth models.Auth
	err := db.Select("id").Where(models.Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}
