package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type PostDaoImpl struct {
}

func (PostDaoImpl) Add(post *models.Post) (err error) {
	post.ID = 0
	err = db.Create(post).Error
	return
}

func (PostDaoImpl) FindById(id uint) (err error, post *models.Post) {
	err = db.Where("id = ?", id).First(post).Error
	return
}

func (PostDaoImpl) DeleteByIds(ids ...uint) (err error) {
	post := models.Post{}
	// 查询帖子id是否存在
	for _, id := range ids {
		if err = db.Where("id = ? and deleted_on = ?", id, 0).First(&post).Error; err != nil {

			return
		}
	}

	// 软删除
	// 事务机制，出错便回滚
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if e := db.Where("id = ?", id).Delete(&post).Error; e != nil {
				return e
			}
		}
		return nil
	})
	return
}

// TODO 分页
func (PostDaoImpl) FindAllByUserId(page page.Page, userId uint) (err error, postList []models.Post) {
	db.Order(page.Sorter+" "+page.Sorter).Where("user_id = ?", userId).Find(&postList)
	err = db.Where("user_id = ?", userId).Find(&postList).Error
	return
}

func (p PostDaoImpl) Update(post *models.Post) (err error) {
	//
	err, _ = p.FindById(post.ID)
	if err == nil {
		// 更新
		err = db.Save(post).Error
	}
	return
}
