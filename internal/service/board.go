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
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/dao/mysql"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/gredis"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type BoardService interface {
	ServiceGetBoardList() (board []models.Board, err error)
	ServiceGetPostList(id int, p *page.Page) (result *page.Result, err error)
	ServiceGetBoard(idInt int) (board *models.Board, err error)
}

type boardService struct {
	dao.BoardDao
	*gredis.RedisDao
}

var ProviderBoard = wire.NewSet(NewBoardService, mysql.NewBoardDao, gredis.NewRedis)

func NewBoardService(dao2 dao.BoardDao, redisDao *gredis.RedisDao) (BoardService, error) {
	return &boardService{dao2, redisDao}, nil
}
func (service *boardService) ServiceGetBoard(idInt int) (board *models.Board, err error) {
	board, err = service.DaoGetBoard(idInt)
	return
}

func (service *boardService) ServiceGetBoardList() (boards []models.Board, err error) {
	boards, err = service.DaoGetBoardList()
	return
}

func (service *boardService) ServiceGetPostList(id int, p *page.Page) (result *page.Result, err error) {
	result, err = service.DaoGetPostList(id, p)
	return
}
