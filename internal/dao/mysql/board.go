package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type boardDao struct {
	*gorm.DB
}

func NewBoardDao(db *gorm.DB) (dao.BoardDao, error) {
	return &boardDao{DB: db}, nil
}
func (dao boardDao) DaoGetPostList(id int, p *page.Page) (result *page.Result, err error) {
	var postList []models.Post
	result, err = page.PageHelper(dao.Where("board_id = ? and deleted_on = 0", id), &postList, p)
	return
}

func (dao boardDao) DaoGetBoard(idInt int) (board *models.Board, err error) {
	board = &models.Board{}
	err = dao.Where("id = ? and deleted_on = 0", idInt).First(board).Error
	return
}

func (dao boardDao) DaoGetBoardList() (boards []models.Board, err error) {
	err = dao.Where(" deleted_on = 0").Find(&boards).Error
	return
}
