package dao

import (
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type BoardDao interface {
	DaoGetBoardList() (broads []models.Board, err error)
	DaoGetPostList(id int, page *page.Page) (result *page.Result, err error)
	DaoGetBoard(idInt int) (board *models.Board, err error)
}
