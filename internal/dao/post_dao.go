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
	"github.com/opensourceai/go-api-service/pkg/page"
)

type PostDao interface {
	// 新增帖子
	DaoAdd(post *models.Post) (err error)
	// 删除帖子
	DaoDeleteByIds(id string, ids ...int) (err error)
	// 通过用户ID查找用户的帖子
	DaoFindAllByUserId(page page.Page, userId int) (err error, postList []models.Post)
	// 返回该帖子的信息
	DaoFindById(id int) (err error, post *models.Post)
	// 返回帖子的信息
	DaoFindByIds(ids ...int) (post []models.Post, err error)
	// 查询主题帖是否在某版块
	DaoFindByBoardIDAndIds(boardID int, ids ...int) (post []models.Post, err error)
	// 更新帖子
	DaoUpdate(post *models.Post) (err error)
	DaoGetOwnPost(p *page.Page, userId string) (result *page.Result, err error)
	DaoUpdatePost(userId string, post *models.Post) (err error)
	DaoGetPost(id string) (post *models.Post, err error)
	DaoGetPostAllComments(id int) (result *page.Result, err error)
	GetPostComments(id int, p *page.Page) (*page.Result, error)
	// 移动主题帖到某个版块
	DaoMovePosts(boardID int, ids ...int) (err error)
}
