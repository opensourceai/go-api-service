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
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/api/v1/dto"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/internal/service"
	"github.com/opensourceai/go-api-service/middleware/jwt"
	"github.com/opensourceai/go-api-service/pkg/app"
)

type CommentApi struct {
}

var ProviderComment = wire.NewSet(NewCommentService, service.ProviderComment)
var commentService service.CommentService

func NewCommentService(comment service.CommentService) (*CommentApi, error) {
	commentService = comment
	return &CommentApi{}, nil
}
func NewCommentRouter(router *gin.Engine) {
	comment := router.Group("/v1/comment")
	// 认证
	comment.Use(jwt.JWT())
	{
		comment.POST("", addComment)
		comment.DELETE("", deleteComment)
		comment.PUT("", updateComment)
	}
}

// @Summary 修改评论
// @Tags Comment
// @Produce  json
// @Param comment body dto.CommentUpdate true "comment"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Security ApiKeyAuth
// @Router /v1/comment [put]
func updateComment(context *gin.Context) {
	appG := app.Gin{C: context}
	// 获取当前用户信息
	userInfo := app.GetUserInfo(context)
	comment := &dto.CommentUpdate{}
	app.BindAndValid(context, comment)
	err := commentService.ServiceUpdate(userInfo.UserId, comment)
	if err != nil {
		fmt.Println(err)
		appG.Fail(nil)
		return
	}

	appG.Success(nil)

}

// @Summary 新增评论
// @Tags Comment
// @Produce  json
// @Param comment body models.Comment true "comment"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Security ApiKeyAuth
// @Router /v1/comment [post]
func addComment(context *gin.Context) {
	appG := app.Gin{C: context}
	comment := &models.Comment{}
	app.BindAndValid(context, comment)
	info := app.GetUserInfo(context)
	if err := commentService.ServiceAdd(info.UserId, comment); err != nil {
		appG.Fail(nil)
		return
	}
	appG.Success(nil)
}

// @Summary 删除评论
// @Tags Comment
// @Produce  json
// @Param ids body dto.Ids true "ids"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Security ApiKeyAuth
// @Router /v1/comment [delete]
func deleteComment(context *gin.Context) {
	appG := app.Gin{C: context}
	ids := &dto.Ids{}
	app.BindAndValid(context, ids)
	userInfo := app.GetUserInfo(context)
	if err := commentService.ServiceDeleteByIds(userInfo.UserId, ids); err != nil {
		appG.Fail(nil)
		return
	}
	appG.Success(nil)
}
