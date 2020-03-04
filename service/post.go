package service

import (
	"github.com/opensourceai/go-api-service/dao"
	"github.com/opensourceai/go-api-service/dao/mysql"
	"github.com/opensourceai/go-api-service/models"
)

type PostService interface {
	AddPost(p *models.Post) (err error)
	DeletePost(ids ...uint) (err error)
}
type PostServiceImpl struct{}

var postDao dao.PostDao

func init() {
	postDao = new(mysql.PostDaoImpl)
}
func (PostServiceImpl) DeletePost(ids ...uint) (err error) {
	return postDao.DeleteByIds(ids...)
}

func (PostServiceImpl) AddPost(p *models.Post) (err error) {

	err = postDao.Add(p)
	return
}
