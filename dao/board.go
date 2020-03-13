package dao

import (
	"github.com/opensourceai/go-api-service/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type BoardDao interface {
	GetBoardList() (broads []models.Board, err error)
	GetPostList(id int, page *page.Page) (result *page.Result, err error)
	GetBoard(idInt int) (board *models.Board, err error)
}
