package service

import (
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/dao/mysql"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/gredis"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type BoardService interface {
	ServiceGetBoardList() (board []models.Board, err error)
	ServiceGetPostList(id int, p *page.Page) (result *page.Result, err error)
	ServiceGetBoard(idInt int) (board *models.Board, err error)
}

type boardService struct {
	dao.BoardDao
	*gredis.RedisDao
}

var ProviderBoard = wire.NewSet(NewBoardService, mysql.NewBoardDao, gredis.NewRedis)

func NewBoardService(dao2 dao.BoardDao, redisDao *gredis.RedisDao) (BoardService, error) {
	return &boardService{dao2, redisDao}, nil
}
func (service *boardService) ServiceGetBoard(idInt int) (board *models.Board, err error) {
	board, err = service.DaoGetBoard(idInt)
	return
}

func (service *boardService) ServiceGetBoardList() (boards []models.Board, err error) {
	boards, err = service.DaoGetBoardList()
	return
}

func (service *boardService) ServiceGetPostList(id int, p *page.Page) (result *page.Result, err error) {
	result, err = service.DaoGetPostList(id, p)
	return
}
