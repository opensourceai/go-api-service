package mysql

import (
	"github.com/opensourceai/go-api-service/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type BoardDaoImpl struct {
}

func (b BoardDaoImpl) GetPostList(id int, p *page.Page) (result *page.Result, err error) {
	var postList []models.Post
	result, err = page.PageHelper(db.Where("board_id = ? and deleted_on = 0", id), &postList, p)
	return
}

func (b BoardDaoImpl) GetBoard(idInt int) (board *models.Board, err error) {
	board = &models.Board{}
	err = db.Where("id = ? and deleted_on = 0", idInt).First(board).Error
	return
}

func (b BoardDaoImpl) GetBoardList() (boards []models.Board, err error) {
	err = db.Where(" deleted_on = 0").Find(&boards).Error
	return
}
