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
	"github.com/opensourceai/go-api-service/pkg/e"
)

// wire依赖
var ProviderComment = wire.NewSet(NewCommentService, mysql.NewCommentDao)

// 依赖注入函数
func NewCommentService(commentDao dao.CommentDao, userDao dao.UserDao, postDao dao.PostDao) (CommentService, error) {
	return &commentService{commentDao, userDao, postDao}, nil
}

//依赖注入结构体
type commentService struct {
	commentDao dao.CommentDao
	userDao    dao.UserDao
	postDao    dao.PostDao
}

//评论service层接口
type CommentService interface {
	// 新增评论
	ServiceAdd(userId int, comment *models.Comment) (err error)
	// 评论删除某用户的评论
	ServiceDeleteByIds(userId int, ids *dto.Ids) (err error)
	// 修改评论
	ServiceUpdate(userId int, comment *dto.CommentUpdateDTO) (err error)
}

func (service *commentService) ServiceUpdate(userId int, comment *dto.CommentUpdateDTO) (err error) {
	c := &models.Comment{}
	c.ID = comment.ID
	c.CommentContent = comment.CommentContent
	//当前评论是否存在
	if comments, err := service.commentDao.DaoFindByIds(comment.ID); len(comments) == 1 && err == nil {
		// 如果当前评论的评论者ID不等于当前用户ID,则无权限
		if comments[0].FromUserID != userId {
			return errors.New("没有权限")
		}

	} else {
		return errors.New(e.GetMsg(e.ERROR))
	}

	// 修改评论
	return service.commentDao.DaoUpdate(c)
}

func (service *commentService) ServiceDeleteByIds(userId int, ids *dto.Ids) error {
	// 检查是否均存在
	if users, err2 := service.commentDao.DaoFindByIds(ids.Ids...); len(users) != len(ids.Ids) || err2 != nil {
		return err2
	}
	return service.commentDao.DaoDeleteByIds(userId, ids.Ids...)
}
func (service *commentService) ServiceAdd(userId int, comment *models.Comment) (err error) {

	if userId != comment.FromUserID {
		return errors.New("当前用户与评论用户不符")
	}
	// 查询改主题是否存在
	err, _ = service.postDao.DaoFindById(comment.PostID)
	if err != nil {
		return
	}
	var ids []int
	if comment.ToUserID == -1 {
		ids = append(ids, comment.FromUserID)
	} else {
		ids = append(ids, comment.FromUserID, comment.FromUserID)
	}
	var users []models.User
	users, err = service.userDao.DaoFindByIds(ids...)
	// 当查询出的用户数量不等于ids长度时,异常
	if len(users) != len(ids) || err != nil {
		return errors.New(e.GetMsg(e.ERROR_USER_NOT_EXIST))
	}
	// 新增
	return service.commentDao.DaoAdd(comment)
}
