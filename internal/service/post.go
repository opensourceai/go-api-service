package service

import (
	"github.com/google/wire"
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
type postService struct {
	dao.PostDao
}

var ProviderPost = wire.NewSet(NewPostService, mysql.NewPostDao)

func NewPostService(dao2 dao.PostDao) (PostService, error) {
	return &postService{dao2}, nil
}

func (service postService) GetPost(id string) (post *models.Post, err error) {
	return service.DaoGetPost(id)
}

func (service postService) UpdatePost(userId string, post *models.Post) (err error) {

	return service.DaoUpdatePost(userId, post)
}

func (service postService) GetOwnPost(page *page.Page, userId string) (postList *page.Result, err error) {
	return service.DaoGetOwnPost(page, userId)
}

func (service postService) DeletePost(userId string, ids ...int) (err error) {
	return service.DaoDeleteByIds(userId, ids...)
}

func (service postService) AddPost(p *models.Post) (err error) {

	err = service.DaoAdd(p)
	return
}
