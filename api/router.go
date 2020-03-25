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

package api

import (
	"net/http"

	v1 "github.com/opensourceai/go-api-service/api/v1"

	"github.com/gin-gonic/gin"
	// swagger
	_ "github.com/opensourceai/go-api-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/opensourceai/go-api-service/pkg/export"
	"github.com/opensourceai/go-api-service/pkg/qrcode"
	"github.com/opensourceai/go-api-service/pkg/upload"
)

// InitRouter initialize routing information
func NewApi() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	// swagger
	//url := ginSwagger.URL("http://0.0.0.0:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Oauth认证
	OauthApi(r)
	// 普通认证
	NewAuthRouter(r)
	// 版块
	v1.NewBoardRouter(r)
	// 主题
	v1.NewPostRouter(r)

	// 用户
	v1.NewUserRouter(r)
	// 评论
	v1.NewCommentRouter(r)

	return r
}
