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

package service

import (
	"errors"
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/api/v1/dto"
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/dao/mysql"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

// wire依赖注入
var ProviderPost = wire.NewSet(NewPostService, mysql.NewPostDao)

// 依赖注入函数
func NewPostService(dao2 dao.PostDao, board dao.BoardDao) (PostService, error) {
	return &postService{dao2, board}, nil
}

// 主题帖服务层
type PostService interface {
	// 新增主题帖
	ServiceAddPost(p *models.Post) (err error)
	// 删除用户自身的主题帖
	ServiceDeletePost(userId string, ids ...int) (err error)
	// [分页]获取用户自身的全部主题帖信息
	ServiceGetOwnPost(page *page.Page, userId string) (postList *page.Result, err error)
	// 更新用户自身的主题帖信息
	ServiceUpdatePost(userId int, post *models.Post) (err error)
	// 获取某和主题的信息
	ServiceGetPost(id string) (post *models.Post, err error)
	// 获取主题的评论[分页]
	ServiceGetPostComments(id int, p *page.Page) (result *page.Result, err error)
	// 移动主题帖到某个版块
	ServiceMovePosts(srcID, targetID int, postIDs *dto.Ids) (err error)
}
type postService struct {
	postDao  dao.PostDao
	boardDao dao.BoardDao
}

func (service postService) ServiceMovePosts(srcID, targetID int, postIDs *dto.Ids) (err error) {
	posts, err := service.postDao.DaoFindByBoardIDAndIds(srcID, postIDs.Ids...)
	if len(posts) != len(postIDs.Ids) {
		return errors.New("不存在")
	}
	return service.postDao.DaoMovePosts(targetID, postIDs.Ids...)
}

func (service postService) ServiceGetPostComments(id int, p *page.Page) (result *page.Result, err error) {
	return service.postDao.GetPostComments(id, p)
}

func (service postService) ServiceGetPost(id string) (post *models.Post, err error) {
	return service.postDao.DaoGetPost(id)
}

func (service postService) ServiceUpdatePost(userId int, post *models.Post) (err error) {
	// 获取主题帖信息
	err, p := service.postDao.DaoFindById(post.ID)
	if err != nil {
		return err
	}
	// 判断主题是否属于该用户
	if p.UserID != userId {
		// 不属于
		return errors.New("无权限")
	}
	//更新用户主题信息
	return service.postDao.DaoUpdatePost(userId, post)
}

func (service postService) ServiceGetOwnPost(page *page.Page, userId string) (postList *page.Result, err error) {
	return service.postDao.DaoGetOwnPost(page, userId)
}

func (service postService) ServiceDeletePost(userId string, ids ...int) (err error) {
	return service.postDao.DaoDeleteByIds(userId, ids...)
}

func (service postService) ServiceAddPost(p *models.Post) (err error) {
	err = service.postDao.DaoAdd(p)
	return
}
