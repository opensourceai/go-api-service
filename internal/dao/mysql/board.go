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
