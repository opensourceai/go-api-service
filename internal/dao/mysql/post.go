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
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/dao/do"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

func NewPostDao(db *gorm.DB) (dao.PostDao, error) {
	return &postDao{DB: db}, nil
}

type postDao struct {
	*gorm.DB
}

func (dao postDao) GetPostComments(id int, p *page.Page) (*page.Result, error) {
	var comments []do.CommentDO

	option, err := page.PageHelperOption(dao.Where("post_id = ?", id), &models.Comment{}, p)
	if err != nil {

		return nil, errors.New("")
	}
	dao.Preload("FromUser").
		Preload("ToUser").
		Where("post_id = ?", id).
		Offset(option.Offset).
		Limit(option.Limit).
		Order(option.Order).
		Find(&comments)
	p1 := &page.Result{Page: p, Data: comments, Option: option}
	return p1, nil

}

func (dao postDao) DaoGetPostAllComments(id int) (result *page.Result, err error) {
	panic("implement me")
}
func (dao postDao) DaoUpdatePost(userId string, post *models.Post) (err error) {
	// 该帖子是否属于该用户
	if err = dao.Where("id = ? and user_id =?", post.ID, userId).Find(&models.Post{}).Error; err != nil {
		return
	}
	err = dao.
		Model(&models.Post{}).
		Where("id = ? and user_id = ?", post.ID, userId).
		Updates(&models.Post{Title: post.Title, Content: post.Content, Summary: post.Summary}).Error
	return
}

func (dao postDao) DaoGetPost(id string) (post *models.Post, err error) {
	post = &models.Post{}
	err = dao.Where("id = ? and deleted_on = 0", id).First(post).Error
	return
}

func (dao postDao) DaoGetOwnPost(p *page.Page, userId string) (result *page.Result, err error) {
	result, err = page.PageHelper(dao.Where("user_id = ? and deleted_on = ?", userId, 0), &[]models.Post{}, p)
	//result, err = page.ExeMysqlPage(db, &[]models.Post{}, p, "user_id = ?", userId)
	return
}

func (dao postDao) DaoAdd(post *models.Post) (err error) {
	post.ID = 0
	err = dao.Create(post).Error
	return
}

func (dao postDao) DaoFindById(id int) (err error, post *models.Post) {
	err = dao.Where("id = ?", id).First(post).Error
	return
}

func (dao postDao) DaoDeleteByIds(userId string, ids ...int) (err error) {
	// 查询帖子id是否存在
	for _, id := range ids {
		if err = dao.Where("id = ? and user_id = ? and deleted_on = ?", id, userId, 0).First(&models.Post{}).Error; err != nil {

			return
		}
	}

	// 软删除
	// 事务机制，出错便回滚
	err = dao.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if e := dao.Where("id = ?", id).Delete(&models.Post{}).Error; e != nil {
				return e
			}
		}
		return nil
	})
	return
}

func (dao postDao) DaoFindAllByUserId(page page.Page, userId int) (err error, postList []models.Post) {
	dao.Order(page.Sorter+" "+page.Sorter).Where("user_id = ?", userId).Find(&postList)
	err = dao.Where("user_id = ?", userId).Find(&postList).Error
	return
}

func (dao postDao) DaoUpdate(post *models.Post) (err error) {
	//
	err, _ = dao.DaoFindById(post.ID)
	if err == nil {
		// 更新
		err = dao.Save(post).Error
	}
	return
}
