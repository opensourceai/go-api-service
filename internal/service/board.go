package service

import (
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/dao/mysql"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

var boardDao dao.BoardDao

func init() {
	boardDao = new(mysql.BoardDaoImpl)
}

type BoardService interface {
	GetBoardList() (board []models.Board, err error)
	GetPostList(id int, p *page.Page) (result *page.Result, err error)
	GetBoard(idInt int) (board *models.Board, err error)
}

type BoardServiceImpl struct {
}

func (b BoardServiceImpl) GetBoard(idInt int) (board *models.Board, err error) {
	board, err = boardDao.GetBoard(idInt)
	return
}

func (b BoardServiceImpl) GetBoardList() (boards []models.Board, err error) {
	boards, err = boardDao.GetBoardList()
	return
}

func (b BoardServiceImpl) GetPostList(id int, p *page.Page) (result *page.Result, err error) {
	result, err = boardDao.GetPostList(id, p)
	return
}
