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

package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/internal/service"
	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/e"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type BoardApi struct {
}

var ProviderBoard = wire.NewSet(NewBoardApi, service.ProviderBoard)

// 待注入的service
var boardService service.BoardService

// 注入
func NewBoardApi(service2 service.BoardService) (*BoardApi, error) {
	boardService = service2
	return &BoardApi{}, nil
}

func NewBoardRouter(router *gin.Engine) {
	broad := router.Group("/v1/broad")
	{
		broad.GET("", getBroadList)
		broad.GET("/:id", getBroad)
		broad.GET("/:id/posts", getPostListInBroad)

	}
}

// @Summary 获取全部版块列表信息
// @Tags Board
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/broad [get]
func getBroadList(context *gin.Context) {
	appG := app.Gin{C: context}
	if boards, err := boardService.ServiceGetBoardList(); err == nil {
		appG.Success(boards)
	} else {
		appG.Fail(nil)
	}
}

// @Summary 获取某版块信息
// @Tags Board
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/broad/{id} [get]
func getBroad(context *gin.Context) {

	appG := app.Gin{C: context}
	id := context.Param("id")
	if id == "" {
		appG.Fail(nil)
		return
	}
	if idInt, err := strconv.Atoi(id); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	} else {
		// 获取数据
		if board, err := boardService.ServiceGetBoard(idInt); err == nil {
			appG.Success(board)
		} else {
			appG.Fail(nil)
		}
	}

}

// @Summary 获取某版块的帖子
// @Tags Board
// @Produce  json
// @Param id path string true "id"
// @Param page query page.Page true "page"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/broad/{id}/posts [get]
func getPostListInBroad(context *gin.Context) {

	appG := app.Gin{C: context}

	// 板块ID
	id := context.Param("id")
	if id == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	// 转整型
	if idInt, err := strconv.Atoi(id); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	} else {
		// 绑定分页参数
		bindPage := page.BindPage(context)
		// 获取数据
		if list, err := boardService.ServiceGetPostList(idInt, bindPage); err != nil {
			appG.Response(http.StatusBadRequest, e.ERROR_POST_NOT_EXIST, nil)
		} else {
			appG.Success(list)
		}
	}

}
