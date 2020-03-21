package dao

import (
	"github.com/opensourceai/go-api-service/internal/models"
)

type CommentDao interface {
	// 新增评论
	DaoAdd(comment *models.Comment) (err error)
	// 通过IDs删除评论
	DaoDeleteByIds(userId int, ids ...int) (err error)
	// 修改评论的内容
	DaoUpdate(comment *models.Comment) (err error)
	// 查找评论
	DaoFindByIds(ids ...int) (comments []models.Comment, err error)
}
