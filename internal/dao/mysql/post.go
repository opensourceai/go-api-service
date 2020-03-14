package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type PostDaoImpl struct {
}

func (p PostDaoImpl) UpdatePost(userId string, post *models.Post) (err error) {
	// 该帖子是否属于该用户
	if err = db.Where("id = ? and user_id =?", post.ID, userId).Find(&models.Post{}).Error; err != nil {
		return
	}
	err = db.
		Model(&models.Post{}).
		Where("id = ? and user_id = ?", post.ID, userId).
		Updates(&models.Post{Title: post.Title, Content: post.Content, Summary: post.Summary}).Error
	return
}

func (p PostDaoImpl) GetPost(id string) (post *models.Post, err error) {
	post = &models.Post{}
	err = db.Where("id = ? and deleted_on = 0", id).First(post).Error
	return
}

func (PostDaoImpl) GetOwnPost(p *page.Page, userId string) (result *page.Result, err error) {
	result, err = page.PageHelper(db.Where("user_id = ? and deleted_on = ?", userId, 0), &[]models.Post{}, p)
	//result, err = page.ExeMysqlPage(db, &[]models.Post{}, p, "user_id = ?", userId)
	return
}

func (PostDaoImpl) Add(post *models.Post) (err error) {
	post.ID = 0
	err = db.Create(post).Error
	return
}

func (PostDaoImpl) FindById(id int) (err error, post *models.Post) {
	err = db.Where("id = ?", id).First(post).Error
	return
}

func (PostDaoImpl) DeleteByIds(userId string, ids ...int) (err error) {
	// 查询帖子id是否存在
	for _, id := range ids {
		if err = db.Where("id = ? and user_id = ? and deleted_on = ?", id, userId, 0).First(&models.Post{}).Error; err != nil {

			return
		}
	}

	// 软删除
	// 事务机制，出错便回滚
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if e := db.Where("id = ?", id).Delete(&models.Post{}).Error; e != nil {
				return e
			}
		}
		return nil
	})
	return
}

// TODO 分页
func (PostDaoImpl) FindAllByUserId(page page.Page, userId int) (err error, postList []models.Post) {
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
