/*
 *    Copyright 2020 opensourceai
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/models"
)

func NewCommentDao(db *gorm.DB) (dao.CommentDao, error) {
	return &commentDao{db}, nil
}

type commentDao struct {
	db *gorm.DB
}

func (dao commentDao) DaoUpdate(comment *models.Comment) (err error) {
	err = dao.db.Model(&models.Comment{}).Where("id = ?", comment.ID).Update("comment_content", comment.CommentContent).Error
	return
}

func (dao commentDao) DaoDeleteByIds(userId int, ids ...int) (err error) {
	// 启用事务机制
	err = dao.db.Transaction(func(tx *gorm.DB) error {
		for id := range ids {
			return dao.db.Where("id = ? and from_uid = ? deleted_on = 0", id, userId).Delete(&models.Comment{}).Error
		}
		return nil
	})
	return
}

func (dao commentDao) DaoAdd(comment *models.Comment) (err error) {
	comment.ID = 0
	// 新增
	dao.db.Create(comment)
	return
}

func (dao commentDao) DaoFindByIds(ids ...int) (comments []models.Comment, err error) {
	// 查询评论
	comments = []models.Comment{}
	err = dao.db.Where("id in (?) and deleted_on = 0", ids).Find(&comments).Error
	return
}
