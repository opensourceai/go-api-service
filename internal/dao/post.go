package dao

import (
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type PostDao interface {
	// 新增帖子
	DaoAdd(post *models.Post) (err error)
	// 删除帖子
	DaoDeleteByIds(id string, ids ...int) (err error)
	// 通过用户ID查找用户的帖子
	DaoFindAllByUserId(page page.Page, userId int) (err error, postList []models.Post)
	// 返回该帖子的信息
	DaoFindById(id int) (err error, post *models.Post)
	// 更新帖子
	DaoUpdate(post *models.Post) (err error)
	DaoGetOwnPost(p *page.Page, userId string) (result *page.Result, err error)
	DaoUpdatePost(userId string, post *models.Post) (err error)
	DaoGetPost(id string) (post *models.Post, err error)
}
