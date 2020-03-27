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
