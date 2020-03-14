package service

import (
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/dao/mysql"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type PostService interface {
	AddPost(p *models.Post) (err error)
	DeletePost(userId string, ids ...int) (err error)
	GetOwnPost(page *page.Page, userId string) (postList *page.Result, err error)
	UpdatePost(userId string, post *models.Post) (err error)
	// 获取某和主题的信息
	GetPost(id string) (post *models.Post, err error)
}
type PostServiceImpl struct{}

func (i PostServiceImpl) GetPost(id string) (post *models.Post, err error) {
	return postDao.GetPost(id)
}

func (i PostServiceImpl) UpdatePost(userId string, post *models.Post) (err error) {

	return postDao.UpdatePost(userId, post)
}

func (i PostServiceImpl) GetOwnPost(page *page.Page, userId string) (postList *page.Result, err error) {
	return postDao.GetOwnPost(page, userId)
}

var postDao dao.PostDao

func init() {
	postDao = new(mysql.PostDaoImpl)
}
func (PostServiceImpl) DeletePost(userId string, ids ...int) (err error) {
	return postDao.DeleteByIds(userId, ids...)
}

func (PostServiceImpl) AddPost(p *models.Post) (err error) {

	err = postDao.Add(p)
	return
}
